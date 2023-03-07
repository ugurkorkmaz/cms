package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var UI embed.FS

func Dist() http.FileSystem {
	distSub, err := fs.Sub(UI, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(distSub)
}
