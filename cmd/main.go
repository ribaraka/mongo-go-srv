package main

import (
	"flag"
	"fmt"
	"github.com/ribaraka/mongo-go-srv/pkg/db"
	"github.com/ribaraka/mongo-go-srv/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ribaraka/mongo-go-srv/config"
)

func main() {
	var confile = flag.String("confile", "./cmd/config.yaml", "to specify config file please use flag -confile")
	flag.Parse()
	conf, err := config.LoadConfig(*confile)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("mongoURI", conf.Mongo)

	mongo := db.GetMongoDbConnection(conf)
	database := mongo.Database(conf.Mongo.Database)
	collection := database.Collection(conf.Mongo.Collection)

	postHandler := handlers.NewPostHandler(collection)
	//getHandler := handlers.ConfirmEmail()
	//checkEmailHandler := handlers.CheckBusyEmail()
	//loginHandler := handlers.SignIn()
	//profileHandler := handlers.GetProfile()


	r := mux.NewRouter()
	//r.HandleFunc("/verify", getHandler)
	//r.HandleFunc("/profile", profileHandler)
	//r.HandleFunc("/possession", checkEmailHandler)
	//r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/form", postHandler).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(conf.StaticAssets))).Methods(http.MethodGet)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}