package main

import (
	"fmt"
	"reflect"
	"strconv"
	"tiche/db"
	"tiche/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func mainUI(w fyne.Window) fyne.CanvasObject {
	name := widget.NewEntry()
	// 基础项目输入
	heightEntry := widget.NewEntry()
	weightEntry := widget.NewEntry()
	stJumpEntry := widget.NewEntry()
	lgJumpEntry := widget.NewEntry()
	sitReachEntry := widget.NewEntry()
	balanceBeamEntry := widget.NewEntry()
	gripStrengthEntry := widget.NewEntry()
	obstacleRunEntry := widget.NewEntry()

	// var dataList []models.PhysicalInfo
	// db.Dao.Model(&models.PhysicalInfo{}).Find(&dataList)
	// 数据列表
	// list := widget.NewList(
	// 	func() int { return len(dataList) },
	// 	func() fyne.CanvasObject { return widget.NewLabel("") },
	// 	func(i int, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(fmt.Sprintf("%s - %.1fkg", dataList[i].Name, dataList[i].Weight))
	// 	},
	// )
	// list.OnSelected = func(id int) { showDetailWindow(dataList[id]) }
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "学生姓名", Widget: name, HintText: "中文名称"},
			{Text: "身高", Widget: heightEntry},
			{Text: "体重", Widget: weightEntry},
			{Text: "立定跳远", Widget: stJumpEntry},
			{Text: "双脚连续跳", Widget: lgJumpEntry},
			{Text: "坐位体前屈", Widget: sitReachEntry},
			{Text: "走平衡木", Widget: balanceBeamEntry},
			{Text: "握力", Widget: gripStrengthEntry},
			{Text: "15m绕障碍跑", Widget: obstacleRunEntry},
		},
	}
	form.OnCancel = func() {
		fmt.Println("Cancelled")
		for _, item := range form.Items {
			if entry, ok := item.Widget.(*widget.Entry); ok {
				entry.SetText("")
			}
		}
	}

	form.OnSubmit = func() {
		fmt.Println("Form submitted")
		// 收集数据
		heightFloat, _ := strconv.ParseFloat(heightEntry.Text, 64)
		weightFloat, _ := strconv.ParseFloat(weightEntry.Text, 64)
		stJumpFloat, _ := strconv.ParseFloat(stJumpEntry.Text, 64)
		lgJumpFloat, _ := strconv.ParseFloat(lgJumpEntry.Text, 64)
		sitReachFloat, _ := strconv.ParseFloat(sitReachEntry.Text, 64)
		balanceBeamFloat, _ := strconv.ParseFloat(balanceBeamEntry.Text, 64)
		gripStrengthFloat, _ := strconv.ParseFloat(gripStrengthEntry.Text, 64)
		obstacleRunFloat, _ := strconv.ParseFloat(obstacleRunEntry.Text, 64)
		info := models.PhysicalInfo{
			Name:         name.Text,
			Height:       heightFloat,
			Weight:       weightFloat,
			StandingJump: stJumpFloat,
			LongJump:     lgJumpFloat,
			SitReach:     sitReachFloat,
			BalanceBeam:  balanceBeamFloat,
			GripStrength: gripStrengthFloat,
			ObstacleRun:  obstacleRunFloat,
		}

		// 数据收集与存储逻辑
		result := db.Dao.Model(&models.PhysicalInfo{}).Create(&info)
		if result.Error != nil {
			fmt.Println(result.Error)
			dialog.ShowError(result.Error, w)
		} else {
			form.OnCancel()
			dialog.ShowInformation("提示", "数据保存成功", w)
		}
	}
	var dataList []models.PhysicalInfo
	db.Dao.Model(&models.PhysicalInfo{}).Find(&dataList)

	physicalInfoRef := reflect.TypeOf(models.PhysicalInfo{})

	table := widget.NewTable(
		func() (int, int) { return len(dataList), physicalInfoRef.NumField() },
		func() fyne.CanvasObject {
			return widget.NewLabel("test1,test2,test3,test4")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				fmt.Println(id.Row, id.Col)
			} else {
				// field := physicalInfoRef.Field(id.Col)
				label.SetText("")
				// label.SetText(dataList[id.Row].Field(id.Col).Interface())
			}
		})
	table.SetColumnWidth(0, 102)
	table.SetRowHeight(2, 50)
	return container.NewVBox(form, table)
}

// func validateNumberInput(entry *widget.Entry) (float64, error) {
// 	if val, err := strconv.ParseFloat(entry.Text, 64); err != nil {
// 		entry.SetValidationError(fmt.Errorf("请输入数字"))
// 		return 0, err
// 	} else {
// 		return val, nil
// 	}
// }

func main() {
	db.Init("./tiche.db")
	db.Dao.AutoMigrate(&models.PhysicalInfo{})

	myApp := app.New()
	myWindow := myApp.NewWindow("体测记录")
	title := widget.NewLabel("体测数据")
	content := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), mainUI(myWindow)), nil, nil, nil, nil)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.ShowAndRun()
}
