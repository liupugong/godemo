package main

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"
import "time"

func main() {
	db, err := gorm.Open("mysql", "root:Password@1@tcp(127.0.0.1:3306)/sakila")
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

	var result city
	fmt.Println("--------------- struct / orm ---------------")
	// db.Debug().Last(&result)
	query := db.First(&result)
	if query.Error != nil {
		fmt.Print(query.Error)
		fmt.Println("")
	}
	fmt.Println("records", query.RowsAffected)

	fmt.Println("id: ", result.city_id)
	fmt.Println("name: ", result.city)
	fmt.Println("country_id: ", result.country_id)

	// var merchants []city
	// query = db.Find(&merchants).Limit(10)
	// fmt.Println("records", len(merchants))

	// for i, mer := range merchants {
	// 	fmt.Println("i = ", i)
	// 	fmt.Println("id: ", mer.city_id)
	// 	fmt.Println("name: ", mer.city)
	// 	fmt.Println("city: ", mer)
	// }

	fmt.Println("--------------- end of struct / orm ---------------")

	fmt.Println("")
	fmt.Println("--------------- raw sql --------------------------")
	rows, err := db.Raw("select city_id, city, country_id, last_update from city limit 10").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
	}

	var city_id int
	var city string
	var country_id int
	var last_update time.Time

	// Fetch rows
	for rows.Next() {
		rows.Scan(&city_id, &city, &country_id, &last_update)
		fmt.Println("id: ", city_id)
		fmt.Println("merchant_name: ", city)
		fmt.Println("auth_key: ", country_id)
		fmt.Println("secret_key: ", last_update)
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// fmt.Print(result)
}

type city struct {
	city_id     int
	city        string
	country_id  int
	last_update time.Time
}
