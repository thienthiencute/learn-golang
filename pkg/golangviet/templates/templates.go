package templates

import (
	"encoding/json"
	"gocas/conf"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func AssetURL(p string) string {
	domain := conf.Domain
	if domain == "" {
		domain = "/"
	}
	return domain + p
}

func SafeURL(path string) templ.SafeURL {
	domain := conf.Domain
	if domain == "" {
		domain = "/"
	}
	return templ.SafeURL(domain + path)
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func JsonString(list interface{}) string {
	bytes, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
