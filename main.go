package main

type Address struct {
	AddressId uint   `db:"address_id"`
	Street    string `db:"street"`
	City      string `db:"city"`
	Region    string `db:"region"`
	Zip       string `db:"zip"`
	Country   string `db:"country"`
}

type User struct {
	UserId    uint   `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Address   Address
}
