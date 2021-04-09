package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/justgetahouse0/banho/bancho"
	"github.com/justgetahouse0/banho/packets"
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
	bancho := bancho.Bancho{
		Sessions: make(map[string]*bancho.Session),
	}
	http.HandleFunc("/", Handler(bancho))
	log.Fatalln(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
