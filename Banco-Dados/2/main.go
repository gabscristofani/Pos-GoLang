package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 3000.00,
	// })

	// 	var product []Product
	// 	// db.First(&product, "name = ?", "Notebook")
	// 	db.Find(&product)
	// 	fmt.Println(product)

	// 	for _, product := range product {
	// 		fmt.Println(product)
	// 	}
	// }

	// var products []Product
	// db.Where("price > ?", 2000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "Smartphone"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	// db.Delete(&p2)

}
