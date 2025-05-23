package docuval

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/villsource/docuval-identity/pkg/adapters/fiber"
	"github.com/villsource/docuval/configs"
)

type App struct {
	fiberApp  *fiber.App
	appConfig *configs.Config
}

func New() *App {
	return &App{
		fiberApp: fiber.New(),
        appConfig: configs.GetAppConfig(),
	}
}

func (a *App) Start() {
	fmt.Println("Docuval server is running...")

	a.configIdentity()

	a.fiberApp.Listen(fmt.Sprintf(":%s", a.appConfig.Port))
}

func (s *App) configIdentity() {
	s.fiberApp.Use(docuvalIdentityFiberAdapter.NewFiberMiddleware())
}
