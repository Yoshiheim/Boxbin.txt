package router

import (
	"hoxt/internal/handlers"
	"hoxt/internal/helpers"
	"net/http"
)

func InitRoute() {
	// Route all Handlers from /HOXT/internal/*
	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/paste/", handlers.Local)
	http.HandleFunc("/list/", handlers.PastesList)
	http.HandleFunc("/add-paste", handlers.AddPaste)
	http.HandleFunc("/search", handlers.SearchPaste)
	http.Handle("/create", helpers.LimitMiddleware(http.HandlerFunc(handlers.CreatePaste)))
}
