package routes

import (
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/controller"

	"github.com/gorilla/mux"
)

//LoadRoutes :for loading routing
func LoadRoutes() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controller.UserLogin)
	api.HandleFunc("/signup", controller.UserSignUp)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	log.Print("Serving on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}