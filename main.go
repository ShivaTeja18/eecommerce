package main

import (
	"ecommerce/MajorProject"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	a, err := os.Create("log.text")
	custerr := fmt.Errorf("failed to create file %v", a)
	if err != nil {
		fmt.Println(custerr, err)
	} else {
		log.SetOutput(a)
	}

}
func main() {

	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Println("Failed to load", err)
	//}
	//Host := os.Getenv("HOST")
	//User := os.Getenv("USER")
	//
	////dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	////
	//////dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//////os.Env
	////fmt.Println(DBURL)
	//_, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	//if err != nil {
	//	log.Println("Failed to connect db", err)
	//
	//}
	//fmt.Println(DBURL)
	//var (
	//	Host     = os.Getenv("host")
	//	User     = os.Getenv("user")
	//	Password = os.Getenv("password")
	//	DBname   = os.Getenv("dbname")
	//)
	//var DBURL = fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=prefer", Host, User, Password, DBname)
	h := MajorProject.Dbhand()
	//var DNS = os.Getenv("dns")
	//dbc.Dbinit(DNS)
	var r = gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.53:8000"})
	//dbc.r.GET("/fetch/:customer_number", h.Fetch)
	//dbc.r.GET("/fbyonum/:order_number", h.FbyOrderNum)
	r.POST("/newEmp", h.CreateEmp)
	r.POST("/newProd", h.NewProd)
	r.DELETE("/remove/:employee_number", h.Delet)
	r.GET("/fetchemps", h.FetchEmp)
	r.POST("/payments", h.Paybyid)
	r.DELETE("/rmpay", h.Remv)
	r.GET("/fpaybyid", h.FetchPay)
	r.GET("/fromtotill", h.FromtoTill)
	r.GET("/fetchprod", h.Fproducts)
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
