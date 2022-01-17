package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ribaraka/mongo-go-srv/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ribaraka/mongo-go-srv/pkg/models"
)

func NewPostHandler(collection *mongo.Collection) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		v := validator.New()
		err = v.Struct(user)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				fmt.Println(e)
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbClient := db.DBClient{
			Ctx: r.Context(),
			Col: collection,
		}

		result, err := dbClient.SignUp(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("result is:", result)
		w.WriteHeader(http.StatusOK)
	}
}