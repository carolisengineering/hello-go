
package main 

import (
	"fmt" 
	"net/http"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "i made this http server!")
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "hello and welcome!")
    })
 
    http.ListenAndServe(":5050", nil)
}
