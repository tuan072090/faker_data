package main

import (
	"faker/cmd"
	"fmt"
)

/*
- function tạo table
- function gen data cho table
*/
func main() {
	cmd.Execute()
	//	Connect database
	// dbConnect := database.GetConn()

	// fmt.Println("dbConnect...", dbConnect)
	// if dbConnect == nil {
	// 	fmt.Println("Connect db failed.")
	// 	os.Exit(1)
	// }

	// //	create all tables
	// err := database.CreateProductTables(dbConnect)
	// if err != nil {
	// 	fmt.Println("CreateProductTables failed.", err)
	// 	os.Exit(2)
	// }

	// fakeDataErr := database.InsertFakeData(dbConnect, 20000)
	// if fakeDataErr != nil {
	// 	fmt.Println("fakeDataErr failed.", fakeDataErr)
	// }

	fmt.Println("faker start done")
}
