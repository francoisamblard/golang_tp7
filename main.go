package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	// curl -X POST -d '{"id":"will-be-omitted","title":"awesomeness"}' http://localhost:3333/articles
)

type Mouton struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Age    float32 `json:"age"`
	Weight float32 `json:"weight"`
}

// Database représente la base de données en mémoire.
var Database = make(map[int]Mouton)
var mutex sync.RWMutex
var currentID int

func GetMoutons(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	moutons := make([]Mouton, len(Database))
	i := 0
	for _, mouton := range Database {
		moutons[i] = mouton
		i++
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moutons)
}

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
