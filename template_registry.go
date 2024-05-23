package main

import (
	"errors"
	"io"
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func NewTemplateRegistry(root string) *TemplateRegistry {
	tr := &TemplateRegistry{}
	tr.ParseTemplates(root)

	return tr
}

func (t *TemplateRegistry) ParseTemplates(root string) error {
	t.templates = make(map[string]*template.Template)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		t.templates[d.Name()] = template.Must(template.ParseFiles(path, "views/base.html"))
		return nil
	})

	return err
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
