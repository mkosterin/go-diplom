package web

import (
	"bytes"
	"diplom/internal/dataStructs"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(host string, status *dataStructs.ResultT) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// curl http://localhost:8080
		w.WriteHeader(200)
		w.Write(structToBytes(*status))
	})

	http.ListenAndServe(host, r)
}

func structToBytes(s any) []byte {
	//Any struct to []bytes
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()
}
