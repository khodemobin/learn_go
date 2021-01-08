package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello.logger.Println("Hello world")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
