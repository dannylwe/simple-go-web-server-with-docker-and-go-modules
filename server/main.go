package main

import (
	"fmt"
	"net/http"
)

//HelloWeb function returns a welcome message
func HelloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to Go")
}

func main() {
	fmt.Println("server running on port 8090")
	http.HandleFunc("/", HelloWeb)
	http.ListenAndServe(":8090", nil)
	
}
