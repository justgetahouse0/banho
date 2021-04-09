package packets

import (
	"log"
	"strings"

	"github.com/justgetahouse0/banho/bancho"
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
	//username := data[0]
	//password := data[1]

	session := (*bancho.Session).New(nil)
	b.Sessions[session.Token] = session

	buff := LoginReply(LoginNeedUpdate)

	log.Println(string(input))
	return buff, "no"
}
