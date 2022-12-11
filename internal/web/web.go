package web

import (
	"bytes"
	"diplom/internal/dataStructs"
	"diplom/internal/repository"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Router(host string, config dataStructs.Config, countries map[string]string) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		// curl http://localhost:8585/api
		w.WriteHeader(200)
		w.Write(structToBytes(repository.RefreshStatusPage(config, countries)))
	})
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "internal/static"))
	FileServer(r, "/", filesDir)

	http.ListenAndServe(host, r)

}

func structToBytes(s any) []byte {
	//Any struct to []bytes
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
