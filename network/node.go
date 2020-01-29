package network

import (
	"blogServerNode/models"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// NODE ROLES MASTER AND SLAVE

// ServerMaster comment
type ServerMaster struct {
	Address   string
	Port      string
	SlaveList []ServerNode
	Database  models.DBHandler
}

// ServerNode comment
type ServerNode struct {
	Address       string
	Port          string
	MasterAddress string
}

// ExposeToMaster comment
func (server ServerNode) ExposeToMaster(db1 *models.DBHandler) {

	err := rpc.Register(db1)

	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":"+server.Port)

	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %s", server.Port)
	// Start accept incoming HTTP connections
	http.Serve(listener, nil)
}

// NotifyMasterUser comment
func (server ServerNode) NotifyMasterUser(user models.User) {
	var reply models.DBHandler
	fmt.Println("NotifyMaster with address", server.MasterAddress)
	client, err := rpc.DialHTTP("tcp", server.MasterAddress)
	if err != nil {
		log.Fatal("error")
	}
	client.Call("ServerMaster.NotifyNodesUser", user, &reply)
}

// NotifyMasterBlogCreate comment
func (server ServerNode) NotifyMasterBlogCreate(blog models.Blog) {
	var reply models.DBHandler
	client, err := rpc.DialHTTP("tcp", server.MasterAddress)
	if err != nil {
		log.Fatal("error")
	}
	client.Call("ServerMaster.NotifyNodesBlogCreate", blog, &reply)
}

// NotifyMasterBlogUpdate comment
func (server ServerNode) NotifyMasterBlogUpdate(blog models.Blog, reply models.Blog) {
	client, err := rpc.DialHTTP("tcp", server.MasterAddress)
	if err != nil {
		log.Fatal("error")
	}
	client.Call("ServerMaster.NotifyNodesBlogUpdate", blog, &reply)
}

// NotifyMasterBlogUpdateTitle comment
func (server ServerNode) NotifyMasterBlogUpdateTitle(blog models.Blog, reply models.Blog) {
	client, err := rpc.DialHTTP("tcp", server.MasterAddress)
	if err != nil {
		log.Fatal("error")
	}
	client.Call("ServerMaster.NotifyNodesBlogTitleUpdate", blog, &reply)
}

// ConnectToMaster comment
func (server ServerNode) ConnectToMaster() {
	var reply string
	client, err := rpc.DialHTTP("tcp", server.MasterAddress)
	if err != nil {
		log.Fatal("error")
	}
	client.Call("ServerMaster.AddNode", server, &reply)
}
