package main

import ( // this import the library
	"log"      // this is use to print start server message and the error message in the terminal
	"net/http" // the core go library that handles http protocol
	"time"     // need to define the timeout

	"github.com/Yogeshpant992/http-basic-with-go/pkg/server"

)





func main() {

	address := ":8080"        // Set the port address to 8080. This : means that listen on all available netwrok interfaces.
	mux := http.NewServeMux() // mux is like the traffic controller. based on the url it decides which function should work with the given url.
    srv := server.New()
	mux.HandleFunc("/", srv.HandleIndex)  //'/' is ---> Listen for anyone visiting the root address
		
	mux.HandleFunc("/users/create", srv.HandleCreaterUsers) //'/' is ---> Listen for anyone visiting the root address
	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second, // tells that how long the server has to wait for the client(the browser) to finsh sending the request.
		WriteTimeout: 10 * time.Second, // tells that how long the server has to finish sending its response back to the client
		// if server takes long than 10 second to either recieve or send response kill the connection as soon as possible.
		MaxHeaderBytes: 1 << 20, // it prevent someone to send 100 MB header to protect RAM from crashing. The MaxHeader byte define the size of maximum header as such no HTTP request header size should not be greater than this .
	}
	log.Printf("Start server: %v", address) //Prints a message so that we know that the code is working.

	log.Fatal(s.ListenAndServe()) // This is the blocking function that wait for infinte time untill user to connect.

}
