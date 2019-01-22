package main

import (
	"fmt"
	"goa-cellar/app"

	"github.com/goadesign/goa"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement
	// IDに0が指定されたらNotFound返すようにしてみる
	if ctx.BottleID == 0 {
		return ctx.NotFound()
	}

	// レスポンスデータ
	bottle := app.GoaExampleBottle{
		ID:   ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}

	return ctx.OK(&bottle)
	// BottleController_Show: end_implement
}
