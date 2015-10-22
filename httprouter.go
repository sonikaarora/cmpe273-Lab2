package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

//Request structure for POST
type Request struct {
  Name string
}

// Response structure for POST
type Response struct  {
   Greeting  string
}

// Method invoked for GET request
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

//Method invoked for POST request
func postHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

   request := Request{}
   response := Response{}

   json.NewDecoder(req.Body).Decode(&request)

   greetingsString := "Hello, " + request.Name + "!"
   response.Greeting = greetingsString

   jsonResponse, _ := json.Marshal(response)
   rw.Header().Set("Content-Type", "application/json")
   rw.WriteHeader(200) // Status code for success
   fmt.Fprintf(rw, "%s", jsonResponse)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", postHello)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
