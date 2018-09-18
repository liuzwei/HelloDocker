package main

import (
	"net/http"
	"fmt"
)

func helloDocker(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello Docker")
}

func main() {
	http.HandleFunc("/hello", helloDocker)
	http.ListenAndServe(":8080", nil)

}
