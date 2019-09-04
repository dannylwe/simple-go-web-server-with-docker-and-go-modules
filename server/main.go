package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

//HelloWeb function returns a welcome message
func HelloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to Go")
	log.Info("success, 200")
}

func main() {
	fmt.Println("server running on port 8585")
	http.HandleFunc("/", HelloWeb)
	http.ListenAndServe(":8585", nil)
	
}
