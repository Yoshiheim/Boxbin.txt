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
	"strings"
)

type NewPaste struct {
	ID    uint
	Title string
}

func SearchPaste(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")

	keyword = helpers.DestroySpaces(keyword)

	keyword = html.EscapeString(keyword)

	if helpers.CheckSizeString(keyword, 35) {
		fmt.Fprintf(w, "Bro this keyword is sooo big(35 symbols limit)")
		return
	}

	var pastes []modules.Paste
	var finded_pastes []NewPaste

	db.DB.Find(&pastes)

	for _, v := range pastes {
		if strings.Contains(v.Title, keyword) {
			newtitle := strings.ReplaceAll(v.Title, keyword, strings.ToUpper(keyword))
			finded_pastes = append(finded_pastes, NewPaste{
				Title: newtitle,
				ID:    v.ID,
			})
		}
	}

	tpl, err := template.New("SearchPaste.html").Funcs(helpers.FuncMap).ParseFiles("./templates/SearchPaste.html", "./templates/search.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	//and render for client
	tpl.Execute(w, map[string]any{
		"pastes":  finded_pastes,
		"keyword": keyword,
	})
}
