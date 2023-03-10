package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var Dist embed.FS

func SPA() http.FileSystem {
	distSub, err := fs.Sub(Dist, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(distSub)
}
