package main

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"
import "time"

func main() {
	db, err := gorm.Open("mysql", "eleme:eleMe@tcp(172.16.10.24:3306)/payment")
	if nil != err {
		fmt.Print(err)
		fmt.Println("")
	}
	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	db.SingularTable(true)

	var result merchant
	fmt.Println("--------------- struct / orm ---------------")
	// db.Debug().Last(&result)
	query := db.First(&result)
	if query.Error != nil {
		fmt.Print(query.Error)
		fmt.Println("")
	}
	fmt.Println("records", query.RowsAffected)

	fmt.Println("id: ", result.id)
	fmt.Println("merchant_name: ", result.merchant_name)
	fmt.Println("merchant: ", result)

	var merchants []merchant
	query = db.Find(&merchants)
	fmt.Println("records", len(merchants))

	for i, mer := range merchants {
		fmt.Println("i = ", i)
		fmt.Println("id: ", mer.id)
		fmt.Println("merchant_name: ", mer.merchant_name)
		fmt.Println("merchant: ", mer)
	}

	fmt.Println("--------------- end of struct / orm ---------------")

	fmt.Println("")
	fmt.Println("--------------- raw sql --------------------------")
	rows, err := db.Raw("select id, merchant_name, auth_key, secret_key from merchant limit 10").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
	}

	var id int64
	var merchant_name string
	var auth_key string
	var secret_key string

	// Fetch rows
	for rows.Next() {
		rows.Scan(&id, &merchant_name, &auth_key, &secret_key)
		fmt.Println("id: ", id)
		fmt.Println("merchant_name: ", merchant_name)
		fmt.Println("auth_key: ", auth_key)
		fmt.Println("secret_key: ", secret_key)
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// fmt.Print(result)
}

type merchant struct {
	id                     int64
	merchant_name          string
	auth_key               string
	secret_key             string
	permit_payment_methods string
	is_valid               bool
	created_at             time.Time
	updated_at             time.Time
}
