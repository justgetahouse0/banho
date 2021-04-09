package bancho

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Bancho struct {
	Sessions map[string]*Session
	Database *r.Session
}

func (b *Bancho) New() (Bancho) {
	return Bancho {
		Sessions: make(map[string]*Session),
	}
}