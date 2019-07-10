package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type TimeStamp time.Time
func init() {
	fmt.Print("Product at...")
}

// Products modal table
type Products = struct {
	tableName  struct{}  `sql:"products"`
	ProductID  int64     `sql:"pk, product_id"`
	CategoryID int64     `sql:"category_id"`
	Name       string    `sql:",name"`
	Image      string    `sql:",image"`
	Price      float64   `sql:"price, type:real" `
	Body       string    `sql:"body"`
	CreatedAt  int `sql:"created_at, default:now()"`
	UpdatedAt  int `sql:"updated_at, default:now()" `
	IsActive   bool      `sql:"is_active"`
	Features   struct {
		Name string
		Body string
	} `sql:"features, type:jsonb"`
}

// Categories modal table
type Categories = struct {
	tableName  struct{}  `sql:"categories"`
	CategoryID int64     `sql:",category_id,pk"`
	Name       string    `sql:",unique"`
	Body       string    `sql:"body"`
	CreatedAt  int `sql:"created_at, default:now()"`
	UpdatedAt  int `sql:"updated_at, default:now()"`
}

// CreateProductTables creates all table for product service
func CreateProductTables(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	fmt.Println("create table products...", &Products{})
	errProduct := db.CreateTable(&Products{}, options)
	if errProduct != nil {
		return errProduct
	}

	fmt.Println("create table categories...")
	errCategory := db.CreateTable(&Categories{}, options)
	if errCategory != nil {
		return errCategory
	}

	return nil
}

// InsertFakeData inserts categories and product, the number is the product amount
func InsertFakeData(db *pg.DB, number int) error {

	var err error

	categoryArr := []string{"Điện thoại- máy tính", "Điện tử", "Phụ kiện", "Thời trang", "Làm đẹp", "Đồ chơi", "điện gia dụng"}

	for i := 0; i < len(categoryArr); i++ {
		errCat := db.Insert(&Categories{
			Name: categoryArr[i],
		})

		if errCat != nil {
			fmt.Println("errCat...", errCat)
		}
	}

	return err
}
