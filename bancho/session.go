package bancho

import (
	"fmt"

	"github.com/justgetahouse0/banho/common"
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	Token string
	User  *common.User
}

func (s *Session) New() *Session {
	uuid := uuid.NewV4()
	return &Session{
		Token: fmt.Sprint(uuid),
		User:  nil,
	}
}
