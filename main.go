package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	// curl -X POST -d '{"id":"will-be-omitted","title":"awesomeness"}' http://localhost:3333/articles
)

func main() {
	fmt.Println("111")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Post("/mouton", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("222")
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			fmt.Println(b)
			myString := string(b[:])
			fmt.Println(myString)
		}
	})

	http.ListenAndServe(":3333", r)
}

type Mouton struct {
	Id     int
	Name   string
	Age    float32
	Weight float32
}
