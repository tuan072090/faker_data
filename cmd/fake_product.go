package cmd

import (
	"faker/database"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init(){

}

func fakeData() *cobra.Command{
	return &cobra.Command{
		Use:   "fake [data name]",
		Short: "Fake data với tên service nhập vào",
		Long: `Fake data với tên service nhập vào, ví dụ "fake product"`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dbConnect := database.GetConn()

			fmt.Println("dbConnect...", dbConnect)
			if dbConnect == nil{
				fmt.Println("Connect db failed.")
				os.Exit(1)
			}

			//	create all tables
			err := database.CreateProductTables(dbConnect);
			if err != nil{
				fmt.Println("CreateProductTables failed.", err)
				os.Exit(2)
			}

			fakeDataErr := database.InsertFakeData(dbConnect, 10)
			if fakeDataErr != nil{
				fmt.Println("fakeDataErr failed.", fakeDataErr)
			}
		},
	}
}

func Execute(){
	if err := fakeData().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}