package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/rs/cors"
)

func main()  {
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	c:=cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:9010",handler))
}