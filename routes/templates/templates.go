package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type data interface{}

func RenderTemplate(w http.ResponseWriter, tmpl string, data1 interface{}, data2 interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data1)
	err = t.Execute(w, data2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
