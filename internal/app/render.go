package app

import (
	"embed"
	"errors"
	"net/http"
	"os"
	"path"
	"text/template"
)

//go:embed ui
var templates embed.FS

func (app Application) GetTemplate(file, tmpl string) (*template.Template, error) {
	p := path.Join(app.Conf.Applications, file)

	_, err := os.Stat(p)

	if errors.Is(err, os.ErrNotExist) {
		return template.ParseFS(templates, tmpl)
	}

	return template.ParseFiles(p)
}

func (app Application) render(w http.ResponseWriter, r *http.Request, file, tmpl string) {
	ts, err := app.GetTemplate(file, tmpl)

	if err != nil {
		app.Conf.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, app.Cmds)
	if err != nil {
		app.Conf.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
