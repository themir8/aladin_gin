package model

type UserInter interface {
	SaveUser(ulist *User)
	UpdateUser(id uint64, ulist *User) error
	DeleteUser(id uint64) error
	GetUser(id uint64) *User
	// Print()
}
