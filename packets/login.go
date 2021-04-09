package packets

import (
	"strings"

	"github.com/justgetahouse0/banho/bancho"
	"github.com/justgetahouse0/banho/common"
	"golang.org/x/crypto/bcrypt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

const (
	LoginFailed        = -1
	LoginNeedUpdate    = -2
	LoginBanned        = -4
	LoginError         = -5
	LoginNeedSupporter = -6
)

func LoginReply(code int32) Packet {
	return Make(5, 4, code)
}

func Login(input []byte, b bancho.Bancho) ([]byte, string) {
	data := strings.Split(string(input), "\n")
	if len(data) != 4 {
		return nil, ""
	}
	username := data[0]
	password := data[1]

	session := (*bancho.Session).New(nil)
	b.Sessions[session.Token] = session

	var buff Packet

	cursor, err := r.Table("users").Filter(r.Row.Field("Username").Eq(username)).Run(b.Database)
	if err != nil {
		buff = LoginReply(LoginError)
		return buff, "no"
	}

	var user *common.User
	if err := cursor.One(&user); err == r.ErrEmptyResult {
		buff = LoginReply(LoginFailed)
		return buff, "no"
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		buff = LoginReply(LoginFailed)
		return buff, "no"
	}

	if user.IsBanned() {
		buff = LoginReply(LoginBanned)
		return buff, "no"
	}

	session.User = user

	// i dunno what to do next help pls

	return buff, session.Token
}
