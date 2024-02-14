package controllers

import (
	"github.com/MohabMohamed/mohab.dev/src/templates"
	"github.com/MohabMohamed/mohab.dev/src/util"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return util.Render(c, templates.Health())
}
