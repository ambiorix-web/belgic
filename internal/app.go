package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

var mu sync.Mutex
var idx int = 0

// home handles the home page of the application and the reverse
// proxy redirection.
func (lb loadBalancer) balance(procs []proc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		maxLen := len(lb.Backends)
		// Round Robin
		mu.Lock()
		currentBackend := lb.Backends[idx%maxLen]
		targetURL, err := url.Parse(currentBackend.Path)
		if err != nil {
			log.Fatal(err.Error())
		}
		idx++
		mu.Unlock()
		reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
		reverseProxy.ServeHTTP(w, r)
	}
}

// handlers binds handlers to the server.
func (lb loadBalancer) handlers(procs []proc) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", lb.balance(procs))
	return mux
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
	host  string
	proxy *httputil.ReverseProxy
}

// createProxies creates an individual reverse proxy for
// each application.
func (lb loadBalancer) createProxies() []proc {
	var proxies []proc
	for _, cmd := range lb.Backends {
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
			proxy: httputil.NewSingleHostReverseProxy(uri),
		}

		proxies = append(proxies, p)
	}

	return proxies
}

// StartApp starts the application and creates the proxies.
func (lb loadBalancer) StartApp() error {
	proxies := lb.createProxies()
	srv := &http.Server{
		Addr:     ":" + lb.Config.Port,
		ErrorLog: lb.ErrorLog,
		Handler:  lb.handlers(proxies),
	}

	return srv.ListenAndServe()
}
