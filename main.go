package main

import (
	"database/sql"
	"encoding/json"
	customerresthandler "fifentory/customer/handler/rest"
	productresthandler "fifentory/product/handler/rest"
	skuinresthandler "fifentory/skuin/handler/rest"
	skuingroupresthandler "fifentory/skuingroup/handler/rest"
	supplierresthandler "fifentory/supplier/handler/rest"

	skuoutresthandler "fifentory/skuout/handler/rest"
	skuoutgroupresthandler "fifentory/skuoutgroup/handler/rest"
	stockresthandler "fifentory/stock/handler/rest"
	skuresthandler "fifentory/stockkeepingunit/handler/rest"

	imagehandler "fifentory/image/handler"
	skuimageresthandler "fifentory/skuimage/handler/rest"
	skuspreadsheethandler "fifentory/stockkeepingunit/handler/spreadsheet"
	variantresthandler "fifentory/variant/handler/rest"
	variantgroupresthandler "fifentory/variantgroup/handler/rest"

	skupriceresthandler "fifentory/skuprice/handler/rest"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	config := map[string]interface{}{}
	configFile, err := os.Open("../src/fifentory/config.json")
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(configFile).Decode(&config)

	conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/"+fmt.Sprint(config["database"])+"?parseTime=true")
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
	skuinresthandler.InjectSKUInRESTHandler(conn, ee)
	skuingroupresthandler.InjectSKUInGroupRESTHandler(conn, ee)
	supplierresthandler.InjectSupplierRESTHandler(conn, ee)
	variantresthandler.InjectVariantRESTHandler(conn, ee)
	variantgroupresthandler.InjectVariantGroupRESTHandler(conn, ee)
	skuspreadsheethandler.InjectSKUSpreadSheetHandler(conn, ee)
	skupriceresthandler.InjectSKUPriceRESTHandler(conn, ee)
	imagehandler.InjectImagesHandler(conn, ee)
	skuimageresthandler.InjectSKUImageRESTHandler(conn, ee)
	ee.Start(":555")
}
