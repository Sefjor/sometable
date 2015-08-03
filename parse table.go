package main

import (
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cnt?charset=utf8")
	err = db.Ping() // Open doesn't open a connection. Validate DSN data:
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT Артикул, Наименование, Цена, Склад1, Склад2, Склад3, Склад1 + Склад2 + Склад3 as t FROM Товар")
	data := [][]string{}
	var num string
		var name string
		var price string
		var s1, s2, s3, sTotal string
	for rows.Next() {
		rows.Scan(&num, &name, &price, &s1, &s2, &s3, &sTotal) // Тут можно просто выводить значения в консоль, но тогда уползёт таблица
		data = append(data, []string{num, name, price, s1, s2, s3, sTotal})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Артикул", "Наименование", "Цена", "Склад1", "Склад2", "Склад3", "Всего на складах"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}