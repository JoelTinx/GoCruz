package main

import (
  "log"
  "net/http"
  "os"

  "github.com/joeltinx/GoCruz/handlers"
)

func main() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/",          handlers.IndexHandler)
  http.HandleFunc("/convert",   handlers.ConvertHandler)
  http.HandleFunc("/gallery",   handlers.GalleryHandler)
  http.HandleFunc("/about",     handlers.AboutHandler)
  http.HandleFunc("/display",   handlers.DisplayHandler)
  http.HandleFunc("/export",   handlers.DisplayHandler) // NF

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("Listening in port: %s", port)
  http.ListenAndServe(":" + port, nil)
}
