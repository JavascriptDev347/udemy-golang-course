package main

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

type TemplatesRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool
	templateDir string
}

type templateData struct {
	Form            Form
	IsAuthenticated bool
	Flash           string
}

func NewTemplatesRenderer(templateDir string, isDev bool) *TemplatesRenderer {
	return &TemplatesRenderer{
		templateDir: templateDir,
		cache:       make(map[string]*template.Template),
		dev:         isDev,
	}
}

func (t *TemplatesRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplatesRenderer) getTemplate(name string) (*template.Template, error) {
	if !t.dev {
		t.mutex.Lock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.Unlock()
			return tmpl, nil
		}
		t.mutex.Unlock()
	}

	tmpl, err := t.parseTemplate(name)
	if err != nil {
		return nil, err
	}

	if !t.dev {
		t.mutex.Lock()
		t.cache[name] = tmpl
		t.mutex.Unlock()
	}

	return tmpl, nil
}

func (t *TemplatesRenderer) parseTemplate(name string) (*template.Template, error) {
	templatePath := path.Join(t.templateDir, name)
	files := []string{templatePath}

	layoutPath := path.Join(t.templateDir, "layouts/*.html")
	layouts, err := filepath.Glob(layoutPath)
	if err == nil {
		files = append(files, layouts...)
	}

	partialPath := path.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialPath)
	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
