package gorm

import (
	"Go_Test/util"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price *uint
}

func TestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	var p uint = 100
	db.Create(&Product{Code: "D42", Price: &p})

	// Read
	var product Product
	//db.First(&product, 1)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(util.Marshal(product))

	// Update - 将 product 的 price 更新为 200
	p = 200
	db.Model(&product).Update("Price", 200)
	fmt.Println(util.Marshal(product))

	// Update - 更新多个字段
	p = 0
	db.Model(&product).Updates(Product{Price: &p, Code: "F42"}) // 仅更新非零值字段
	fmt.Println(util.Marshal(product))
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println(util.Marshal(product))

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
