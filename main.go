package main

import (
	"math/rand"
	"net/http"
	"time"
)

func randomChoice(s []string) string{
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var s []string
		s = append(s, "Trick")
		s = append(s, "Treat")
		writer.Write([]byte(randomChoice(s)))
	})

	http.ListenAndServe(":8080", nil)
}