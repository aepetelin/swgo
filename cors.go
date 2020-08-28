package main

import (
	"github.com/go-chi/cors"
)

var corsMiddleware = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://local.website.dev:5000", "http://localhost:8801", "http://localhost:1323"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-XSRF-TOKEN", "X-JWT"},
	ExposedHeaders:   []string{"Authorization"},
	AllowCredentials: true,
	MaxAge:           300,
})
