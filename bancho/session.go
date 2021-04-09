package bancho

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	Token string
}

func (s *Session) New() *Session {
	token := uuid.NewV4()
	return &Session{
		Token: fmt.Sprint(token),
	}
}
