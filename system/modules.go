package system

import "fmt"

func NilError(msg string , err error) {
	if err != nil {
		fmt.Println(msg,err)
	}
}