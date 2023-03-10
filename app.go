//go:generate go run github.com/99designs/gqlgen@latest generate
//go:generate npm run build --prefix template/
//go:generate go build .
package main

import (
	"app/resolver"
	"app/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(resolver.NewSchema())

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/", serveSPA)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func serveSPA(w http.ResponseWriter, r *http.Request) {
	buildPath := "dist"
	f, err := template.Dist.Open(filepath.Join(buildPath, r.URL.Path))
	if os.IsNotExist(err) {
		index, err := template.Dist.ReadFile(filepath.Join(buildPath, "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	http.FileServer(template.SPA()).ServeHTTP(w, r)
}
