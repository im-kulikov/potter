package config

import (
	"github.com/im-kulikov/potter/logger"
	"github.com/labstack/echo"
)

const infAttachTpl = "Attach [%s]%s - %q"

type API []Endpoint

func (a API) Attach(engine *echo.Echo, log logger.Logger) {
	for _, endpoint := range a {
		log.Infof(
			infAttachTpl,
			endpoint.Method,
			endpoint.URL,
			endpoint.Fixture,
		)
		endpoint.attach(engine)
	}
}

type Endpoint struct {
	Method  string `yaml:"method"`
	URL     string `yaml:"url"`
	Fixture string `yaml:"fixture"`
}

func (e Endpoint) handler(ctx echo.Context) error {
	return ctx.File(e.Fixture)
}

func (e Endpoint) attach(engine *echo.Echo) {
	engine.Add(e.Method, e.URL, e.handler)
}
