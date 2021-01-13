package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	mux := http.NewServeMux()

	url, err := url.Parse("http://loki:3100")
	if err != nil {
		panic(err)
	}
	rp := httputil.NewSingleHostReverseProxy(url)

	director := rp.Director
	rp.Director = func(r *http.Request) {
		fmt.Printf("%s \n", r.RequestURI)
		director(r)
	}

	mux.Handle("/loki/api/v1/query_range", rp)
	mux.Handle("/loki/api/v1/query", rp)
	mux.Handle("/loki/api/v1/label", rp)
	mux.Handle("/loki/api/v1/labels", rp)
	mux.Handle("/loki/api/v1/label/{name}/values", rp)
	mux.Handle("/loki/api/v1/series", rp)
	mux.Handle("/api/prom/query", rp)
	mux.Handle("/api/prom/label", rp)
	mux.Handle("/api/prom/label/{name}/values", rp)
	mux.Handle("/api/prom/series", rp)

	// defer tail endpoints to the default handler
	mux.Handle("/loki/api/v1/tail", rp)
	mux.Handle("/api/prom/tail", rp)

	fmt.Println("Welcome to loki-acls")

	panic(http.ListenAndServe("0.0.0.0:3100", mux))
}
