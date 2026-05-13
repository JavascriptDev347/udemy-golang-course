package main

import (
	"html/template"
	"net/http"
	"sync"
)

type TemplatesRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool
	templateDir string
}

func NewTemplatesRenderer(templateDir string, isDev bool) *TemplatesRenderer {
	return &TemplatesRenderer{
		cache: make(map[string]*template.Template),
		dev:   isDev,
	}
}

func (t *TemplatesRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplatesRenderer) getTemplate(name string) (*template.Template, error) {
	if !t.dev {
		t.mutex.RLock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.Unlock()
			return tmpl, nil
		}
		t.mutex.Unlock()
	}

	t.mutex.RLock()
}
