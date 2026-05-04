package handlers

import (
	"fmt"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func PastesList(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "No topic id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid topic id", http.StatusBadRequest)
		return
	}

	var pastes []modules.Paste

	var count int64
	var offset int64

	db.DB.Model(&modules.Paste{}).Count(&count)

	if count <= offset {
		offset = count
	} else {
		offset = offset * int64(id)
	}

	db.DB.Offset(int(offset)).Limit(5).Find(&pastes)
	for i := range pastes {
		pastes[i].Title = html.EscapeString(pastes[i].Title)
	}
	fmt.Println(pastes)

	tpl, err := template.New("PastesList.html").Funcs(helpers.FuncMap).ParseFiles("./templates/PastesList.html", "./templates/attr.html", "./templates/search.html")
	if err != nil {
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"pastes": pastes,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
