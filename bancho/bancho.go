package bancho

type Bancho struct {
	Sessions map[string]*Session
}

func (b *Bancho) New() (Bancho) {
	return Bancho {
		Sessions: make(map[string]*Session),
	}
}