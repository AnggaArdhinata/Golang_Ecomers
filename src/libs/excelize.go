package libs

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/AnggaArdhinata/indochat/src/models"
)

func GenerateXls(data []models.OrderJoin) {
	xlsx := excelize.NewFile()

	sheet1Name := "Order"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Order_Id")
	xlsx.SetCellValue(sheet1Name, "B1", "Customer_Name")
	xlsx.SetCellValue(sheet1Name, "C1", "Order_Date")
	xlsx.SetCellValue(sheet1Name, "D1", "Total_Price")
	xlsx.SetCellValue(sheet1Name, "E1", "Status")

	err := xlsx.AutoFilter(sheet1Name, "A1", "E1", "")
	if err != nil {
		log.Fatal("error filter", err.Error())
	}

	for i , v := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), v.Id)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), v.Cust_Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), v.Order_Date)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), v.Price)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), v.Status)
	}

	err = xlsx.SaveAs("./Order_report.xlsx")
	if err != nil {
		log.Println(err)
	}
}