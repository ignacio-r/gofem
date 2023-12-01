package main

import (
	"fmt"
	"net/http"
	"text/template"

	"fem.com/go/museum/api"
	"fem.com/go/museum/data"
	"github.com/gin-gonic/gin"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from a go program!!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))

		return
	}
	//This is Server side rendering
	html.Execute(w, data.GetAll())
}
func main() {

	go func() {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run() // listen and serve on 0.0.0.0:8080
	}()

	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)

	if err == nil {
		fmt.Println("error while opening the server")
	}
}
