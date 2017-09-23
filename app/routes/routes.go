package routes

import (
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/controller"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

//LoadRoutes :for loading routing
func LoadRoutes() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", controller.UserLogin)
	api.HandleFunc("/signup", controller.UserSignUp)
	api.HandleFunc("/addfields", controller.AddFields)
	api.HandleFunc("/listusers", controller.ListUsers)
	api.HandleFunc("/user/{id:[0-9]+}", controller.User).Methods("GET", "DELETE")
	api.HandleFunc("/user", controller.User).Methods("POST", "PUT")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"content-type"},
	})
	handler := c.Handler(r)

	log.Print("Serving on port 9000")
	log.Fatal(http.ListenAndServe(":9000", handler))
}
