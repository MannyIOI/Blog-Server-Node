package api

import (
	"blogServerNode/models"
	"blogServerNode/network"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Server comment
type Server struct {
	Database   *models.DBHandler
	NodeServer network.ServerNode
}

func (server Server) getUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	username := ""
	// var err error
	if val, ok := pathParams["username"]; ok {
		username = val
	}

	w.WriteHeader(http.StatusOK)

	var reply models.User
	server.Database.GetUser(username, &reply)
	json.NewEncoder(w).Encode(&reply)
}

func (server Server) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("------>", server.NodeServer)
	server.NodeServer.NotifyMasterUser(user)
	// var reply models.User
	// server.Database.CreateUser(user, &reply)
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(&reply)

}

// StartAPI comment
func StartAPI(db models.DBHandler, address string, port int, node network.ServerNode) {
	server := Server{Database: &db, NodeServer: node}

	// server.NodeServer
	r := mux.NewRouter()
	api := r.PathPrefix("").Subrouter()

	api.HandleFunc("/user/{username}/", server.getUser).Methods("GET")

	api.HandleFunc("/user/", server.createUser).Methods("POST")

	fmt.Println("Listening on ", address+":"+strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
