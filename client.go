package main

import (
	"blogServerNode/api"
	"blogServerNode/models"
	"blogServerNode/network"
)

func main() {

	db1 := new(models.DBHandler)
	db1.Init()
	db1.Migrate()

	server := network.ServerNode{Address: "localhost", Port: "1237", MasterAddress: "localhost:1234"}
	go server.ConnectToMaster()
	go server.ExposeToMaster(db1)

	api.StartAPI(*db1, "localhost", 1238, server)

}
