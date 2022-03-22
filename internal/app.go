package internal

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/devOpifex/belgic/internal/config"
)

var mu sync.Mutex
var idx int = 0
var attempts int

// home handles the home page of the application and the reverse
// proxy redirection (load balancer).
func (lb *loadBalancer) balance(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	i := idx % len(lb.Backends)
	selectedBackend := &lb.Backends[i]
	targetURL, err := url.Parse(selectedBackend.Path)
	if err != nil {
		lb.ErrorLog.Printf(err.Error())
	}
	idx++
	mu.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		if attempts > lb.Config.Attempts {
			return
		}

		lb.ErrorLog.Printf(
			"%v is dead, attempting %v to create new one",
			attempts,
			selectedBackend.Port,
		)

		attempts++
		mu.Lock()
		var back config.Backend
		err := back.RunApp()
		// too much recursion
		// should skip?
		if err != nil {
			lb.balance(w, r)
		}

		lb.Backends = append(lb.Backends[:i], lb.Backends[i+1:]...)
		lb.Backends = append(lb.Backends, back)
		mu.Unlock()
		lb.balance(w, r)
	}
	attempts = 0
	lb.InfoLog.Printf(
		"%v on %s directed to %v",
		r.Method,
		r.URL.Path,
		selectedBackend.Port,
	)
	reverseProxy.ServeHTTP(w, r)
}

// handlers binds handlers to the server.
func (lb loadBalancer) handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", lb.balance)
	return mux
}

// serve serves a reverse proxy
func (p proc) serve(w http.ResponseWriter, r *http.Request) {
	r.Host = p.host
	w.Header().Set("X-Powered-By", "Belgic")
	p.proxy.ServeHTTP(w, r)
}

// proc rerpresents a proxy.
type proc struct {
	host  string
	proxy *httputil.ReverseProxy
}

// StartApp starts the application and creates the proxies.
func (lb loadBalancer) StartApp() error {
	srv := &http.Server{
		Addr:     ":" + lb.Config.Port,
		ErrorLog: lb.ErrorLog,
		Handler:  lb.handlers(),
	}

	return srv.ListenAndServe()
}
