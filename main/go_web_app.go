package main

import "net/http"

func main() {
	http.HandleFunc("/", helloFunc)
	http.ListenAndServe(":8080", nil)

}
func helloFunc(wr http.ResponseWriter, req *http.Request) {

}
