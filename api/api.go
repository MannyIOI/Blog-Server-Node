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
	enableCors(&w)
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
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	server.NodeServer.NotifyMasterUser(user)
	json.NewEncoder(w).Encode(user)
}

func (server Server) createBlog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var blog models.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	server.NodeServer.NotifyMasterBlogCreate(blog)
	json.NewEncoder(w).Encode(blog)
}

func (server Server) updateBlog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var blog models.Blog
	var reply models.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	server.NodeServer.NotifyMasterBlogUpdate(blog, reply)
	json.NewEncoder(w).Encode(reply)
}

func (server Server) getBlog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	blogID := ""
	// var err error
	if val, ok := pathParams["blogIdentifier"]; ok {
		blogID = val
	}

	w.WriteHeader(http.StatusOK)

	val, err := strconv.Atoi(blogID)
	if err == nil {
		var reply models.Blog
		server.Database.GetBlog(uint(val), &reply)
		json.NewEncoder(w).Encode(&reply)
	}

}

func (server Server) getAllBlogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all blogs called remotely")
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var reply []models.Blog
	server.Database.GetAllBlogs("nil", &reply)
	json.NewEncoder(w).Encode(&reply)
}

func enableCors(w *http.ResponseWriter) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

// StartAPI comment
func StartAPI(db models.DBHandler, address string, port int, node network.ServerNode) {
	server := Server{Database: &db, NodeServer: node}

	// server.NodeServer
	r := mux.NewRouter()
	api := r.PathPrefix("").Subrouter()

	api.HandleFunc("/user/{username}/", server.getUser).Methods("GET")

	api.HandleFunc("/user/", server.createUser).Methods("POST")

	api.HandleFunc("/blog/{blogIdentifier}", server.getBlog).Methods("GET")
	api.HandleFunc("/blog/", server.createBlog).Methods("POST")
	api.HandleFunc("/updateBlog/", server.updateBlog).Methods("POST")

	api.HandleFunc("/blogs/", server.getAllBlogs).Methods("GET")

	fmt.Println("Listening on ", address+":"+strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
