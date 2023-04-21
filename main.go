package main

import (
	"fmt"
	"net/http"
	"syscall/js"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	js.Global().Set("login", js.FuncOf(login))
	done := make(chan struct{})
	js.Global().Get("document").Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("log", "Page loaded")
		close(done)
		return nil
	}))
	<-done
}

func login(this js.Value, args []js.Value) interface{} {
	args[0].Call("preventDefault")
	form := args[0].Get("target")
	username := form.Get("username").Get("value").String()
	password := form.Get("password").Get("value").String()
	if username == "admin" && password == "password" {
		fmt.Println("Authentication successful")
	} else {
		fmt.Println("Authentication failed")
	}

	
	form.Get("username").Set("value", "")
	form.Get("password").Set("value", "")

	return nil
}
