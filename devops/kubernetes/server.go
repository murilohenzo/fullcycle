package main

import (
	"fmt"
	"net/http"
)

type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Settings Up Our App")
	http.HandleFunc("/", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")
	var app App
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Hello Full Cycle </h1>"))
}