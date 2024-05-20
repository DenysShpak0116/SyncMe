package models

type User struct {
	UserId    int    `db:"UserId"`
	Username  string `db:"Username"`
	Password  string `db:"Password"`
	Email     string `db:"Email"`
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
	Sex       string `db:"Sex"`
	Country   string `db:"Country"`
	Role      string `db:"Role"`
}
