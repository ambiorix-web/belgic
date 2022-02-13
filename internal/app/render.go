package app

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/devOpifex/belgic/internal/config"
)

//go:embed ui/index.html
var templates embed.FS

func (app Application) render(w http.ResponseWriter, r *http.Request, index string, apps config.RCommands) {
	tmpls := []string{"ui/index.html"}
	ts, err := template.ParseFS(templates, tmpls...)
	if index != "" {
		ts, err = template.ParseFiles(index)
	}

	if err != nil {
		app.Conf.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, apps)
	if err != nil {
		app.Conf.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
