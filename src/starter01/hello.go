package main

import (
	"fmt"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func main() {

	http.HandleFunc("/", sayHelloName)
	http.ListenAndServe(":9000", nil)

}
