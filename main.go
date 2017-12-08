package main

import (
	"flag"
	"os"

	"github.com/im-kulikov/potter/config"
	"github.com/im-kulikov/potter/logger"
	"github.com/im-kulikov/potter/logger/zap"
	"github.com/im-kulikov/yaml"
	"github.com/labstack/echo"
)

var (
	log   logger.Logger
	bind  = flag.String("bind", ":8081", "server bind address")
	conf  = flag.String("c", "config.yml", "config for server")
	debug = flag.Bool("debug", false, "use debug")
	level = zap.Production
)

func main() {
	var (
		err      error
		file     *os.File
		settings config.Config
	)

	flag.Parse()

	engine := echo.New()
	engine.HideBanner = true

	if *debug {
		level = zap.Development
	}

	log = zap.New(level)
	engine.Logger = log

	// Open config:
	if file, err = os.Open(*conf); err != nil {
		log.Panic(err)
	}

	// Parse config:
	if err = yaml.NewDecoder(file).Decode(&settings); err != nil {
		log.Panic(err)
	}

	// Apply proxy server
	if err = settings.Proxy.Apply(); err != nil {
		log.Panic(err)
	}

	// Attach endpoints:
	settings.API.Attach(engine, log)

	// Start web-server:
	if err = engine.Start(*bind); err != nil {
		log.Panic(err)
	}
}
