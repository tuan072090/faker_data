package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
	"time"
	"github.com/icrowley/fake"
	"github.com/bxcodec/faker"
)

type TimeStamp time.Time

func init() {
	fmt.Print("Product at...")
}

// Products modal table
type Products = struct {
	tableName  struct{} `sql:"products"`
	ProductID  int64    `sql:"product_id,pk"`
	CategoryID int64    `sql:"category_id"`
	Name       string   `sql:",name"`
	Image      string   `sql:",image"`
	Price      float64  `sql:"price" `
	Body       string   `sql:"body"`
	CreatedAt  int32    `sql:"created_at,notnull"`
	UpdatedAt  int32    `sql:"updated_at,notnull" `
	IsActive   bool     `sql:"is_active"`
	Features   struct {
		Name string
		Body string
	} `sql:"features, type:jsonb"`
}

// Categories modal table
type Categories = struct {
	tableName  struct{} `sql:"categories"`
	CategoryID int64    `sql:",category_id,pk"`
	Name       string   `sql:",unique"`
	Body       string   `sql:"body"`
	CreatedAt  int32    `sql:"created_at, notnull"`
	UpdatedAt  int32    `sql:"updated_at, notnull"`
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
	timeStart := int(time.Now().Unix())

	categoryArr := []string{"Điện thoại- máy tính", "Điện tử", "Phụ kiện", "Thời trang", "Làm đẹp", "Đồ chơi", "điện gia dụng"}

	// insert categories
	for i := 0; i < len(categoryArr); i++ {
		errCat := db.Insert(&Categories{
			Name:      categoryArr[i],
			CreatedAt: int32(time.Now().Unix()),
			UpdatedAt: int32(time.Now().Unix()),
		})

		if errCat != nil {
			fmt.Println("errCat...", errCat)
		}
	}

	if number > 0 {
		//	Insert products
		for i := 0; i < number; i++ {
			price, err := strconv.ParseFloat(faker.AmountWithCurrency, 64)

			if err != nil {
				price = 0.0
			}

			errProduct := db.Insert(&Products{
				CategoryID: 1,
				Name: fake.ProductName(),
				Image: fake.TopLevelDomain(),
				Price: price,
				Body: fake.Paragraph(),
				CreatedAt: int32(time.Now().Unix()),
				UpdatedAt: int32(time.Now().Unix()),
				IsActive: true,
			})

			if errProduct != nil {
				fmt.Println("errProduct...", errProduct)
			}
		}
	}

	timeEnd := int(time.Now().Unix())

	fmt.Println("Spend ", strconv.Itoa(timeEnd-timeStart)+" seconds")
	return err
}
