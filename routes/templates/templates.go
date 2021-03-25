package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderTemplateForUser(w http.ResponseWriter, tmpl string, data1 interface{}, data2 interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]interface{}{
		"Profile":      data1,
		"Access-token": data2,
	}
	err = t.Execute(w, m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
