package botkit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Structs
type TelegramUser struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
	PhotoURL     string `json:"photo_url,omitempty"`
}

type FormData struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Utilidades
func validateSignature(initData, botToken string) bool {
	if os.Getenv("DEV_MODE") == "true" {
		return true
	}

	values, err := url.ParseQuery(initData)
	if err != nil {
		return false
	}

	receivedHash := values.Get("hash")
	if receivedHash == "" {
		return false
	}
	values.Del("hash")

	var dataCheckArray []string
	for k, vs := range values {
		for _, v := range vs {
			dataCheckArray = append(dataCheckArray, k+"="+v)
		}
	}
	sort.Strings(dataCheckArray)
	dataCheckString := strings.Join(dataCheckArray, "\n")

	secretKey := sha256.Sum256([]byte("WebAppData"))
	mac := hmac.New(sha256.New, secretKey[:])
	mac.Write([]byte(botToken))
	hmacKey := mac.Sum(nil)

	mac = hmac.New(sha256.New, hmacKey)
	mac.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(mac.Sum(nil))

	return calculatedHash == receivedHash
}

func extractUserData(initData string) (*TelegramUser, error) {
	values, err := url.ParseQuery(initData)
	if err != nil {
		return nil, err
	}

	userStr := values.Get("user")
	if userStr == "" {
		return nil, fmt.Errorf("dados do usuário não encontrados")
	}

	var user TelegramUser
	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func checkBotToken(token string) bool {
	if token == "" {
		return false
	}
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", token))
	if err != nil {
		log.Printf("Erro ao verificar token: %v", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func sendTelegramMessage(botToken, chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("erro: %s", string(body))
	}

	return nil
}

// Middleware para validar o Telegram Init Data
func withTelegramAuth(handler func(http.ResponseWriter, *http.Request, *TelegramUser)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		initData := r.Header.Get("Telegram-Web-App-Init-Data")
		if initData == "" {
			http.Error(w, "Telegram-Web-App-Init-Data não fornecido", http.StatusUnauthorized)
			return
		}

		botToken := os.Getenv("BOT_TOKEN")
		if !validateSignature(initData, botToken) {
			http.Error(w, "Assinatura inválida", http.StatusUnauthorized)
			return
		}

		user, err := extractUserData(initData)
		if err != nil {
			http.Error(w, "Dados do usuário inválidos", http.StatusBadRequest)
			return
		}

		handler(w, r, user)
	})
}

// Handlers
func submitDataHandler(w http.ResponseWriter, r *http.Request, user *TelegramUser) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var form FormData
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		http.Error(w, "BOT_TOKEN não configurado", http.StatusInternalServerError)
		return
	}

	chatID := fmt.Sprintf("%d", user.ID)
	message := fmt.Sprintf("📝 Nova mensagem de %s:\n\nNome: %s\nMensagem: %s",
		user.FirstName, form.Name, form.Message)

	log.Printf("Enviando mensagem para %s (ID: %s)", user.FirstName, chatID)

	err = sendTelegramMessage(botToken, chatID, message)
	if err != nil {
		log.Printf("Erro ao enviar mensagem: %v", err)
		http.Error(w, fmt.Sprintf("Erro ao enviar mensagem: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Mensagem enviada com sucesso!",
		"user":    user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request, user *TelegramUser) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func serveWebApp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, map[string]interface{}{
		"title": "Telegram WebApp",
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/webapp", http.StatusMovedPermanently)
}

func rootPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/api/submit", http.StatusMovedPermanently)
}

func MiniAppEX() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN não definido")
	}

	log.Println("Verificando token do bot...")
	if !checkBotToken(botToken) {
		log.Println("AVISO: O token parece inválido ou o bot não está respondendo.")
	} else {
		log.Println("Token do bot verificado com sucesso.")
	}

	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8080"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Rotas
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/webapp", serveWebApp)
	http.HandleFunc("/post", rootPostHandler)

	http.Handle("/api/submit", withTelegramAuth(submitDataHandler))
	http.Handle("/api/user-info", withTelegramAuth(getUserInfoHandler))

	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Printf("Servidor iniciado em %s na porta %s", serverURL, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
