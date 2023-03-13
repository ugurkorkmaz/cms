package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:generate pnpm build
//go:embed all:dist
var Dist embed.FS

func SPA() http.FileSystem {
	distSub, err := fs.Sub(Dist, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(distSub)
}
