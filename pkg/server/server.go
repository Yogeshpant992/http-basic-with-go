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
 
// user represent the value that are sent as a response to a request 
type user struct{
	Name string `json:"Name"`
	Email string `json:"Email"`
	Age int `json:"Age"`

}
type Server struct {
	users map[string]userinfo
}

// userinfo is the information that is stored per user
type userinfo struct{
	email string;
	roll_no int;
	age int;
}


//HandleIndex handle the index path "/"
func  ( s *Server) HandleIndex(w  http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK) // send the request which it is either requested or stuats ok and all basically provide you the status of the code.
		w.Write([]byte(indexPage))

} 


func (s *Server) HandleCreaterUsers(w http.ResponseWriter, r *http.Request)
{
	switch r.Method{
	case http.MethodPost, http.MethodPut:
          currentType := r.HEader.Get("Content-Type");
		  contentType != 
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) //HTTP 415
	}
}
// HandleIndex handle the user path "/users/create"
 
