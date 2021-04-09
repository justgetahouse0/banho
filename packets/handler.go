package packets

import (
	"io"
	"log"

	"github.com/justgetahouse0/banho/bancho"
)

func Handle(token string, input []byte, output io.Writer, bancho bancho.Bancho) string {
	if token == "" {
		res, token := Login(input, bancho)
		output.Write(res)
		log.Println(res, output)
		return token
	}
	return ""
}
