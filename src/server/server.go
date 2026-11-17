package server

import (
	"net/http"
	"os"
	"path/filepath"
)

// type noDirList struct {
// 	fs http.FileSystem
// }

func ServeFile() http.Handler {
	fileServer := http.FileServer(http.Dir("./public"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filepat := filepath.Join("./public", r.URL.Path)
		info, err := os.Stat(filepat)
		if err != nil || info.IsDir() {
			http.ServeFile(w, r, "./public/index.html")
		}
		fileServer.ServeHTTP(w, r)
	})
}
