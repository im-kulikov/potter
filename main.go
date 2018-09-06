package main

import (
	"context"
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	"github.com/chapsuk/mserv"
	"github.com/im-kulikov/helium"
	"github.com/im-kulikov/helium/grace"
	"github.com/im-kulikov/helium/logger"
	"github.com/im-kulikov/helium/module"
	"github.com/im-kulikov/helium/settings"
	"github.com/im-kulikov/helium/web"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	app struct {
		servers mserv.Server
		log     *zap.Logger
	}

	handler struct {
		Fixture string
		Method  string
		URL     string
		Echo    bool
	}

	handlers []handler
)

var (
	conf = flag.String("c", "config.yml", "config for server")

	BuildTime    = "now"
	BuildVersion = "dev"

	mod = module.Module{
		{Constructor: newApp},
		{Constructor: newAPI},
	}.Append(
		grace.Module,
		settings.Module,
		logger.Module,
		web.ServersModule,
	)
)

const HeaderContentType = "Content-Type"

func echoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set content-type

		w.Header().Set(HeaderContentType,
			r.Header.Get(HeaderContentType))

		// set status code:
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, r.Body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}

func fixtureHandler(fixture string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(fixture)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()

		fi, _ := f.Stat()
		http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
	}
}

func newAPI(log *zap.Logger, v *viper.Viper) (http.Handler, error) {
	var (
		items handlers
		mux   = pat.New()
	)

	if err := v.UnmarshalKey("fixtures", &items); err != nil {
		return nil, err
	}

	for _, item := range items {
		l := log.With(
			zap.String("method", item.Method),
			zap.String("url", item.URL),
			zap.String("fixture", item.Fixture),
			zap.Bool("echo", item.Echo))

		l.Debug("try to connect handler")

		switch {
		case item.URL != "" && item.Echo && item.Method != "GET":
			mux.Post(item.URL, echoHandler())
		case item.URL != "" && item.Fixture != "":
			mux.Add(item.Method, item.URL, fixtureHandler(item.Fixture))
		default:
			l.Warn("ignore handler")
		}
	}

	return mux, nil
}

func newApp(log *zap.Logger, serv mserv.Server) helium.App {
	return &app{
		servers: serv,
		log:     log,
	}
}

func (a *app) Run(ctx context.Context) error {
	a.log.Info("run servers...")
	a.servers.Start()
	a.log.Info("successful run application...")
	<-ctx.Done()
	a.log.Info("stop servers...")
	a.servers.Stop()
	a.log.Info("successful stop application...")
	return nil
}

func main() {
	h, err := helium.New(&settings.App{
		File:         *conf,
		Name:         "potter",
		BuildTime:    BuildTime,
		BuildVersion: BuildVersion,
	}, mod)
	helium.Catch(err)
	helium.Catch(h.Run())
}
