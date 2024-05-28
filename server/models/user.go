package models

type User struct {
	UserId    int    `db:"UserId" json:"userId"`
	Username  string `db:"Username" json:"username"`
	Password  string `db:"Password" json:"password"`
	Email     string `db:"Email" json:"email"`
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	Sex       string `db:"Sex" json:"sex"`
	Country   string `db:"Country" json:"country"`
	Role      string `db:"Role" json:"role"`
}
