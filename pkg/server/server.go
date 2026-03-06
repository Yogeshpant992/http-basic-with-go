package server

import "net/http"
// Server is an HTTP Server 
type Server struct{

}

//New is a new server 
func New() *Server{
	return &Server{}
}

var indexPage = `
<!DOCTYPE html>
<html>
    <body> 
	<h1> My first heading </h1>
	<p> My first paragraph </p>
	</body>
</html>`



//HandleIndex handle the index path "/"
func  ( s *Server) HandleIndex(w  http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK) // send the request which it is either requested or stuats ok and all basically provide you the status of the code.
		w.Write([]byte(indexPage))

} 


// HandleIndex handle the user path "/users/create"
func  ( s *Server) HandleUsers(w  http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK) // send the request which it is either requested or stuats ok and all basically provide you the status of the code.
		w.Write([]byte(userInfo))

} 
