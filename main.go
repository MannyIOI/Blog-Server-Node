package main

import (
	"blogServerNode/models"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	db1 := new(models.DBHandler)
	db1.Init()
	db1.Migrate()
	// var reply models.User

	// var user = models.User{Username: "username", Password: "password"}
	// fmt.Println(user)
	// db.CreateUser(user, &reply)

	// fmt.Println(reply)
	// api.StartAPI(db, "http://localhost", 8081)
	// name := "world"
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	// fmt.Printf("hello, %s\n", name)

	// myAddress := "localhost"
	// serverAddress := "localhost"
	// isMaster := os.Args[1]

	// db1 := new(models.DBHandler)
	err := rpc.Register(db1)

	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)

}
