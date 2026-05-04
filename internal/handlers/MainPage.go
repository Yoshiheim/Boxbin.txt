package handlers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"text/template"
	"time"
)

type TopicWithCount struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	PostCount   int
}

// The Index aka '/' path in website.
// path: 'http://<HOST>:<PORT>/'
func MainPage(w http.ResponseWriter, r *http.Request) {
	/*
		var tops []modules.Topic

		act := db.DB.Find(&tops)
		if act.Error != nil {
			log.Println(act.Error.Error())
			http.Error(w, "Error With DB.", http.StatusInternalServerError)
			return
		}
	*/

	var pastes []modules.Paste

	var count int64
	var offset int64

	db.DB.Model(&modules.Paste{}).Count(&count)

	if count <= offset {
		offset = count
	}

	db.DB.Order("is_titled DESC").Order("created_at DESC").Offset(int(offset)).Limit(40).Find(&pastes)
	for i := range pastes {
		pastes[i].Title = html.EscapeString(pastes[i].Title)
	}

	/*var tops []TopicWithCount

	if err := db.DB.
		Model(&modules.Topic{}).
		Select(`
		topics.*,
		(SELECT COUNT(*) FROM pastes WHERE pastes.topic_id = topics.id) as post_count
	`).Scan(&tops); err.Error != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	for i := range tops {
		tops[i].Name = html.EscapeString(tops[i].Name)
		tops[i].Description = html.EscapeString(tops[i].Description)
	}
	*/

	helpers.UpdateLogo()

	cfg := data.GetDConfig(w)

	tpl, err := template.New("index.html").Funcs(helpers.FuncMap).ParseFiles("./templates/index.html", "./templates/attr.html", "./templates/search.html")
	if err != nil {
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"data":   cfg,
		"logo":   string(data.Logo),
		"pastes": pastes,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
