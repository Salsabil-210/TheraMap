package main

import (
	"Backend/infrastructure" 
	"Backend/bootstrap"
)



func main() {
	infrastructure.ConnectDB() 

	bootstrap.StartServer() 
}
