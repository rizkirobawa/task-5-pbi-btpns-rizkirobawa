package app

// AuthorizedRegister merupakan struktur data yang digunakan untuk user melakukan  registrasi akun
type AuthorizedRegister struct {
	Id       string `valid:"required" json:"id"`
	Username string `valid:"required" json:"username"`
	Email    string `valid:"required" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}

// AuthorizedLogin merupakan struktur data yang digunakan untuk user melakukan login akun
type AuthorizedLogin struct {
	Email    string `valid:"required" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}
