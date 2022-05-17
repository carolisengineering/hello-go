
package main 

import (
	"fmt" 
	"net/http"
)

func IMadeThis() string {
	return "i made this http server!"
}

func Hello(language string) string {
	if language == "french" {
		return "bonjour et bienvenue!"
	}
	if language == "spanish" {
		return "hola y bienvenidos!"
	}
	return "hello and welcome!"
}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, IMadeThis())
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, Hello("french"))
    })
 
    http.ListenAndServe(":5050", nil)
}
