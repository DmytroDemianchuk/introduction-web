package main

import (
	"fmt"
	"net/http"

	"github.com/Easy-Job-Developer/catalog_plus/handlers/data"
	"github.com/Easy-Job-Developer/catalog_plus/handlers/horoshop"
	"github.com/Easy-Job-Developer/catalog_plus/handlers/register"
	"github.com/Easy-Job-Developer/catalog_plus/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://catalogplus.com.ua", "http://24.199.127.152"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	// excel
	r.Post("/upload", data.UploadFileHandler)
	r.Post("/merge", data.MergeFilesHandler)
	r.Post("/get-sheets", data.GetSheets)

	// auth
	r.Post("/horoshop-auth", horoshop.HoroshopAuthHandler)

	// products
	r.Get("/get-products", horoshop.GetProductsDataHandler)
	r.Post("/post-products", horoshop.PostProductsDataHandler)
	r.Get("/get-categories", horoshop.GetCategories)

	// register

	r.Post("/create-user", register.CreateName)
	r.Post("/sign-up", register.SignUpHandler)
	r.Post("sign-in", service.SignIn)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", r)
}
