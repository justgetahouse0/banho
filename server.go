package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/justgetahouse0/banho/bancho"
	"github.com/justgetahouse0/banho/packets"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func Handler(bancho bancho.Bancho) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("cho-protocol", "19")
		w.Header().Set("connection", "keep-alive")
		w.Header().Set("keep-alive", "timeout=5, max=100")

		token := r.Header.Get("osu-token")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		buff := new(bytes.Buffer)
		token = packets.Handle(token, body, buff, bancho)
		if token != "" {
			w.Header().Set("cho-token", token)
		}

		io.Copy(w, buff)
	}
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Username: "admin",
		Password: "scierniskotzndomkukanqa",
		Database: "prudenit",
	})

	if err != nil {
		log.Fatalln(err)
	}

	bancho := bancho.Bancho{
		Sessions: make(map[string]*bancho.Session),
		Database: session,
	}
	http.HandleFunc("/", Handler(bancho))
	log.Fatalln(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
