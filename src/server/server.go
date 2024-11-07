package server

import (
	"log"
	"net/http"
	"os"

	"routes"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
)

func InitServer() {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Method", "Pattern"})

	for i, route := range routes.Routes {
		t.AppendRow([]any{i + 1, route.Method, route.Pattern})
		http.HandleFunc(route.GetMethodPatternStr(), route.Controller.HandleRequest)
	}

	t.Render()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
