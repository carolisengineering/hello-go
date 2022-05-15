
package main 

import (
	"fmt" 
	"net/http"
)

func IMadeThis() string {
	return "i made this http server!"
}

func Hello() string {
	return "hello and welcome!"
}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, IMadeThis())
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, Hello())
    })
 
    http.ListenAndServe(":5050", nil)
}
