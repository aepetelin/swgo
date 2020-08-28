package main

import (
	"fmt"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "swgo/docs"
)

func Route() error {
	router := chi.NewRouter()
	router.Use(corsMiddleware.Handler)

	router.Get("/info", infoHandler)
	router.With(auth).Get("/private", privateHandler)

	// SWAGGER DOCS
	sw := chi.NewRouter()
	sw.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"),
	))
	go http.ListenAndServe(":1323", sw)

	http.Handle("/", router)
	server := &http.Server{Addr: fmt.Sprintf(":%d", 8801)}
	return server.ListenAndServe()
}

// ApplicationsList godoc
// @Summary Show public info
// @Description
// @Tags applications
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /info [get]
func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("info"))
}

// ApplicationsList godoc
// @Summary Show private info
// @Description
// @Tags applications
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /private [get]
// @Security ApiKeyAuth
func privateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("private"))
}

func auth(next http.Handler) http.Handler {
	var isValidToken = func(token string) bool {
		return len(token) > 0
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("auth check")

		token := r.Header.Get("X-JWT")
		if !isValidToken(token) {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		log.Printf("X-JWT: %s", token)

		//a := r.Header.Get("Authorization")
		//if a == "" {
		//	http.Error(w, "not authorized", http.StatusUnauthorized)
		//	return
		//}
		//log.Printf("Authorization: %s", a)
		next.ServeHTTP(w, r)
	})
}
