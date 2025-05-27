package utils

import (
	"fmt"
	"os"
)

// Seu HTML de exemplo (preencha aqui o conteúdo real)
const index_html string = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Telegram WebApp</title>
    <script src="https://telegram.org/js/telegram-web-app.js"></script>
    <style>
        :root {
            --tg-theme-bg-color: #ffffff;
            --tg-theme-text-color: #222222;
            --tg-theme-hint-color: #999999;
            --tg-theme-link-color: #2678b6;
            --tg-theme-button-color: #50a8eb;
            --tg-theme-button-text-color: #ffffff;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: var(--tg-theme-bg-color);
            color: var(--tg-theme-text-color);
        }

        .container {
            max-width: 500px;
            margin: 0 auto;
            padding: 20px;
        }

        h1 {
            font-size: 1.5rem;
            text-align: center;
            margin-bottom: 20px;
        }

        .user-info {
            background-color: rgba(0, 0, 0, 0.05);
            border-radius: 10px;
            padding: 15px;
            margin-bottom: 20px;
            text-align: center;
        }

        .user-info img {
            width: 60px;
            height: 60px;
            border-radius: 50%;
            margin-bottom: 10px;
        }

        .form-group {
            margin-bottom: 15px;
        }

        label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
        }

        input, textarea {
            width: 100%;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 8px;
            font-size: 16px;
            background-color: var(--tg-theme-bg-color);
            color: var(--tg-theme-text-color);
            box-sizing: border-box;
        }

        textarea {
            min-height: 100px;
            resize: vertical;
        }

        button {
            background-color: var(--tg-theme-button-color);
            color: var(--tg-theme-button-text-color);
            border: none;
            border-radius: 8px;
            padding: 12px 20px;
            font-size: 16px;
            cursor: pointer;
            width: 100%;
            margin-top: 10px;
        }

        .hidden {
            display: none;
        }

        .status {
            text-align: center;
            padding: 10px;
            margin-top: 20px;
            border-radius: 8px;
        }

        .success {
            background-color: rgba(76, 175, 80, 0.2);
            color: #2e7d32;
        }

        .error {
            background-color: rgba(244, 67, 54, 0.2);
            color: #c62828;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Meu WebApp Telegram</h1>
        
        <div class="user-info">
            <div id="user-photo"></div>
            <p>Olá, <span id="username">Usuário</span>!</p>
        </div>
        
        <form id="message-form">
            <div class="form-group">
                <label for="name">Nome:</label>
                <input type="text" id="name" name="name" required>
            </div>
            
            <div class="form-group">
                <label for="message">Mensagem:</label>
                <textarea id="message" name="message" required></textarea>
            </div>
            
            <button type="button" id="submit-btn">Enviar Mensagem</button>
        </form>
        
        <div id="status" class="status hidden"></div>
    </div>
    
    <script>
        // Inicializar o WebApp do Telegram
        const tg = window.Telegram.WebApp;
        
        // Expandir para ocupar toda a altura disponível
        tg.expand();
        
        // Aplicar tema do Telegram
        document.documentElement.style.setProperty('--tg-theme-bg-color', tg.themeParams.bg_color);
        document.documentElement.style.setProperty('--tg-theme-text-color', tg.themeParams.text_color);
        document.documentElement.style.setProperty('--tg-theme-hint-color', tg.themeParams.hint_color);
        document.documentElement.style.setProperty('--tg-theme-link-color', tg.themeParams.link_color);
        document.documentElement.style.setProperty('--tg-theme-button-color', tg.themeParams.button_color);
        document.documentElement.style.setProperty('--tg-theme-button-text-color', tg.themeParams.button_text_color);
        
        // Configurar o botão principal do Telegram (se disponível)
        if (tg.MainButton) {
            tg.MainButton.setText('Enviar Mensagem');
            tg.MainButton.onClick(function() {
                sendFormData();
            });
            tg.MainButton.show();
        } else {
            console.log("MainButton não está disponível na API do Telegram");
        }
        
        // Determinar a URL base da API (dinâmico)
        const getApiBaseUrl = () => {
            // Se estamos em um ambiente de desenvolvimento local
            if (window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1') {
                return window.location.origin;
            }
            
            // Se estamos em um ambiente de produção, use o mesmo domínio do WebApp
            return window.location.origin;
        };
        
        // Mostrar mensagem de status
        function showStatus(message, isError = false) {
            const statusElement = document.getElementById('status');
            statusElement.textContent = message;
            statusElement.classList.remove('hidden', 'success', 'error');
            statusElement.classList.add(isError ? 'error' : 'success');
            
            // Esconder após 3 segundos
            setTimeout(() => {
                statusElement.classList.add('hidden');
            }, 3000);
        }
        
        // Função para enviar os dados do formulário
        function sendFormData() {
            const name = document.getElementById('name').value;
            const message = document.getElementById('message').value;
            
            // Validar campos
            if (!name || !message) {
                showStatus('Por favor, preencha todos os campos', true);
                return;
            }
            
            // Desabilitar o botão durante o envio
            const submitBtn = document.getElementById('submit-btn');
            submitBtn.disabled = true;
            submitBtn.textContent = 'Enviando...';
            
            // Desabilitar o botão do Telegram também
            if (tg.MainButton) {
                tg.MainButton.disable();
                tg.MainButton.setText('Enviando...');
            }
            
            // Obter a URL base da API
            const apiBaseUrl = getApiBaseUrl();
            
            // Enviar dados para o servidor
            fetch(apiBaseUrl+'/api/submit', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Telegram-Web-App-Init-Data': tg.initData
                },
                body: JSON.stringify({
                    name: name,
                    message: message
                })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Erro HTTP: '+response.status);
                }
                return response.json();
            })
            .then(data => {
                // Reabilitar o botão
                submitBtn.disabled = false;
                submitBtn.textContent = 'Enviar Mensagem';
                
                // Reabilitar o botão do Telegram
                if (tg.MainButton) {
                    tg.MainButton.enable();
                    tg.MainButton.setText('Enviar Mensagem');
                }
                
                if (data.success) {
                    showStatus('Mensagem enviada com sucesso!');
                    
                    // Limpar o formulário
                    document.getElementById('message').value = '';
                    
                    // Opcional: fechar o WebApp após alguns segundos
                    //setTimeout(() => {
                    //    tg.close();
                    //}, 2000);
                } else {
                    showStatus('Erro ao enviar mensagem: ' + (data.error || 'Erro desconhecido'), true);
                }
            })
            .catch(error => {
                console.error('Erro:', error);
                submitBtn.disabled = false;
                submitBtn.textContent = 'Enviar Mensagem';
                
                // Reabilitar o botão do Telegram
                if (tg.MainButton) {
                    tg.MainButton.enable();
                    tg.MainButton.setText('Enviar Mensagem');
                }
                
                showStatus('Ocorreu um erro ao enviar a mensagem: '+error.message, true);
            });
        }
        
        // Obter dados do usuário
        document.addEventListener('DOMContentLoaded', function() {
            // Exibir nome do usuário
            const user = tg.initDataUnsafe.user;
            if (user) {
                document.getElementById('username').textContent = user.first_name + 
                    (user.last_name ? ' ' + user.last_name : '');
                
                // Exibir foto do usuário se disponível
                if (user.photo_url) {
                    const img = document.createElement('img');
                    img.src = user.photo_url;
                    img.alt = user.first_name;
                    document.getElementById('user-photo').appendChild(img);
                }
                
                // Preencher o campo nome com o nome do usuário
                document.getElementById('name').value = user.first_name;
            }
            
            // Configurar o botão de envio
            const submitBtn = document.getElementById('submit-btn');
            submitBtn.addEventListener('click', sendFormData);
            
            // Configurar o formulário para não fazer submit padrão
            const form = document.getElementById('message-form');
            form.addEventListener('submit', function(e) {
                e.preventDefault();
                sendFormData();
            });
        });
    </script>
</body>
</html>

`

const pastaTemplates string = "./templates"
const arquivoPath string = pastaTemplates + "/index.html"

// Função para tratar erros


func CreateTemplate() {
	// Cria o diretório (inclui subdiretórios se necessário)
	err := os.MkdirAll(pastaTemplates, os.ModePerm)
	NilError("Erro ao criar diretório", err)

	// Cria o arquivo (ou sobrescreve)
	arquivo, err := os.Create(arquivoPath)
	NilError("Erro ao criar o index.html", err)
	defer arquivo.Close()

	// Escreve no arquivo
	_, err = arquivo.WriteString(index_html) // aqui era "arquivo.writes" — só "arquivo.WriteString"
	NilError("Erro ao escrever no arquivo", err)

	fmt.Println("Código exemplo criado com sucesso...")
}
