package app

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/devOpifex/belgic/internal/config"
)

// Application the core belgic application.
type Application struct {
	Conf config.Config
	Cmds config.RCommands
}

// home handles the home page of the application and the reverse
// proxy redirection.
func (app Application) home(procs []proc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" {
			app.render(w, r, "index.html", "ui/index.html")
			return
		}

		path := strings.Split(r.URL.Path, "/")[1]

		for _, p := range procs {
			if p.name != path {
				continue
			}

			p.serve(w, r)
			return
		}

		app.render(w, r, "404.html", "ui/404.html")
	}
}

// handlers binds handlers to the server.
func (app Application) handlers(procs []proc) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home(procs))
	return mux
}

// StartApp starts the application and creates the proxies.
func StartApp(conf config.Config, cmds config.RCommands) error {
	app := Application{
		Conf: conf,
		Cmds: cmds,
	}
	proxies := createProxies(cmds)
	srv := &http.Server{
		Addr:     ":" + conf.Port,
		ErrorLog: conf.ErrorLog,
		Handler:  app.handlers(proxies),
	}

	return srv.ListenAndServe()
}

// serve serves a reverse proxy
func (p proc) serve(w http.ResponseWriter, r *http.Request) {
	r.Host = p.host
	r.URL.Path = cleanPath(r.URL.Path)
	w.Header().Set("X-Powered-By", "Belgic")
	p.proxy.ServeHTTP(w, r)
}

// cleanPath removes the first element from a path for the
// reverse proxy to forward to the correct path.
func cleanPath(path string) string {
	var cleaned string
	paths := strings.Split(path, "/")[2:]

	for _, p := range paths {
		cleaned += "/" + p
	}

	return cleaned
}

// proc rerpresents a proxy.
type proc struct {
	name  string
	host  string
	proxy *httputil.ReverseProxy
}

// createProxies creates an individual reverse proxy for
// each application.
func createProxies(cmds config.RCommands) []proc {
	var proxies []proc
	for _, cmd := range cmds {
		if cmd.Err != nil {
			continue
		}

		target := "http://localhost:" + fmt.Sprint(cmd.Port)
		uri, err := url.Parse(target)

		if err != nil {
			continue
		}

		p := proc{
			host:  uri.Host,
			name:  string(cmd.Name),
			proxy: httputil.NewSingleHostReverseProxy(uri),
		}

		proxies = append(proxies, p)
	}

	return proxies
}
