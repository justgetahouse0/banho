package common

const (
	PermissionBanned = 1 << iota
)

type User struct {
	ID          int
	Username    string
	Password    []byte
	Permissions uint32
}

func (u *User) IsBanned() bool {
	return u.Permissions&PermissionBanned != 0
}
