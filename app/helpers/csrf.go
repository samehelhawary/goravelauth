package helpers

import (
	"github.com/gofiber/fiber/v2"
	"html/template"
)

func CSRFField(ctx *fiber.Ctx) template.HTML {
	token := ctx.Locals("csrf_token")
	if token == nil {
		return ""
	}

	field := `<input type="hidden" name="_csrf" value="` + token.(string) + `" />`
	return template.HTML(field)
}

func CSRFToken(ctx *fiber.Ctx) string {
	token := ctx.Locals("csrf_token")
	if token == nil {
		return ""
	}
	return token.(string)
}
