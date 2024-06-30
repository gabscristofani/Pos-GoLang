package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Producs []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// create product
	db.Create(&Product{
		Name:       "Panela",
		Price:      200.00,
		Categories: []Category{category, category2},
	})

	fmt.Printf("Product created successfully\n")

	var categories []Category
	erro := db.Model(&Category{}).Preload("Producs").Find(&categories).Error
	if erro != nil {
		panic(erro)
	}

	// for _, category := range categories {
	// 	fmt.Println(category.Name, ":")
	// 	for _, product := range category.Producs {
	// 		fmt.Println("- ", product.Name)
	// 	}
	// }
}
