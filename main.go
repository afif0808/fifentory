package main

import (
	"database/sql"
	customerresthandler "fifentory/customer/handler/rest"
	productresthandler "fifentory/product/handler/rest"
	skuoutresthandler "fifentory/skuout/handler/rest"
	skuoutgroupresthandler "fifentory/skuoutgroup/handler/rest"
	stockresthandler "fifentory/stock/handler/rest"
	skuresthandler "fifentory/stockkeepingunit/handler/rest"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/fiprosys?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	ee := echo.New()

	productresthandler.InjectProductRESTHandler(conn, ee)
	skuresthandler.InjectSKURESTHandler(conn, ee)
	skuoutresthandler.InjectSKUOutHandler(conn, ee)
	customerresthandler.InjectCustomerRESTHandler(conn, ee)
	skuoutgroupresthandler.InjectSKUOutGroupRESTHandler(conn, ee)
	stockresthandler.InjectStockRESTHandler(conn, ee)
	ee.Start(":555")
}
