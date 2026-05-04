package handlers

import (
	"hoxt/internal/helpers"
	"html/template"
	"log"
	"net/http"
)

func AddPaste(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("AddPaste.html").Funcs(helpers.FuncMap).ParseFiles("./templates/AddPaste.html", "./templates/attr.html")
	if err != nil {
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
