package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Clothing struct {
	Code string `gorm:"primaryKey"`
	Size string
	Price int64
	Type string
}

type WareHouseInfo struct {
	Code string `gorm:"primaryKey"`
	Capacity int64
}

type SupplyInfo struct {
	ClothingCode string
	VendorCode int64
	QualityLevel string
}
func main(){
	db,err:=gorm.Open("mysql", "root:ROOT@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Clothing{},&WareHouseInfo{},&SupplyInfo{})

	//查询服装尺码为'S'且销售价格在100以下的服装信息
	var C Clothing
	db.Where("Size = ? AND Price <= ?", "S", "22").Find(&C)
	fmt.Println("查询服装尺码为'S'且销售价格在100以下的服装信息")
	fmt.Printf("%v\n",C)

	//查询仓库容量最大的仓库信息。
	var W WareHouseInfo
	db.Order("Capacity DESC").Limit(1).Find(&W)
	fmt.Println("查询仓库容量最大的仓库信息")
	fmt.Printf("%v\n",W)

	//查询A类服装的库存总量。
	var C1  WareHouseInfo
	db.Select("capacity").Where("code = ?","A").Find(&C1)
	fmt.Println("查询A类服装的库存总量")
	fmt.Printf("%v\n",C1)

	//查询服装编码以‘A’开始开头的服装。
	var C2 []Clothing
	db.Where("code like ?","c"+"%").Find(&C2)
	fmt.Println("查询服装编码以‘A’开始开头的服装")
	fmt.Printf("%v\n",C2)

	//查询服装质量等级有不合格的供应商信息。
	var s []SupplyInfo
	db.Select("vendor_code").Where(" quality_level > ?","C").Find(&s)
	fmt.Println("查询服装质量等级有不合格的供应商信息")
	fmt.Printf("%v\n",s)

	//把服装尺寸为'S'的服装的销售价格均在原来基础上提高10%
	db.Model(&Clothing{}).Where("Size = ?","S").Update("Price",gorm.Expr("Price * 1.1"))

	//删除所有服装质量等级不合格的供应情况
	db.Delete(&SupplyInfo{}," quality_level > ?","C")

	//向每张表插入一条记录。
	A1:=WareHouseInfo{Code: "1"}
	A2:=SupplyInfo{VendorCode: 1}
	A3:=Clothing{Code: "A"}
	db.Create(&A1)
	db.Create(&A2)
	db.Create(&A3)
}