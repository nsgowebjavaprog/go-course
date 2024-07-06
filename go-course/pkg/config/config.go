package config

import "html/template"

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
