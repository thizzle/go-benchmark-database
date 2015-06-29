package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

var (
	dbNative *sql.DB
	dbSqlx   *sqlx.DB
)

func init() {
	var err error

	dbNative, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	dbSqlx = sqlx.NewDb(dbNative, "postgres")

	fmt.Printf("Connected to database at %s", os.Getenv("DATABASE_URL"))
}

func cleandb() {
	var err error

	if _, err = dbNative.Exec("TRUNCATE TABLE users CASCADE"); err != nil {
		panic(err)
	}

	if _, err = dbNative.Exec("TRUNCATE TABLE addresses CASCADE"); err != nil {
		panic(err)
	}
}

func validate(user User) bool {
	if user.FirstName != "Tharsan" {
		return false
	}

	return true
}

func BenchmarkNative(b *testing.B) {
	// cleandb()

	for i := 0; i < b.N; i++ {
		rows, err := dbNative.Query("SELECT user_id, first_name, last_name, addresses.address_id, street, city, region, zip, country FROM users INNER JOIN addresses USING (address_id)")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			user := User{}
			err = rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Address.AddressId, &user.Address.Street, &user.Address.City, &user.Address.Region, &user.Address.Zip, &user.Address.Country)
			if err != nil {
				panic(err)
			}

			// if !validate(user) {
			// 	fmt.Println("Sorry user not valid")
			// 	b.Fail()
			// }
		}
	}
}

func BenchmarkSqlx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rows, err := dbSqlx.Queryx("SELECT user_id, first_name, last_name, addresses.address_id, street, city, region, zip, country FROM users INNER JOIN addresses USING (address_id)")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			user := User{}
			err = rows.StructScan(&user)
			if err != nil {
				panic(err)
			}

			// if !validate(user) {
			// 	fmt.Println("Sorry user not valid")
			// 	b.Fail()
			// }
		}
	}
}
