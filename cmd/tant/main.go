package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	// 1. 打开文件
	f, err := excelize.OpenFile("./cmd/tant/contant.xlsx")
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	// 2. 获取全部工作表的名字
	sheetMap := f.GetSheetMap()
	sheets := make([]string, 0, len(sheetMap))
	// 3. 打印工作表的名字
	for _, sheet := range sheetMap {
		if sheet == "总计" {
			continue
		}
		if strings.Contains(sheet, "模板") {
			continue
		}
		sheets = append(sheets, sheet)
	}
	fmt.Println(sheets)
	tants := make(map[string]*Constant, len(sheets))
	for i := 0; i < len(sheets); i++ {
		tant := DealValue(f, sheets[i])
		if tant != nil {
			tants[tant.uid] = tant
		}
	}
	total := make(map[string]int32)
	sortTotal := make(map[string][]int32)
	sortTotal["water"] = make([]int32, 0, len(tants))
	sortTotal["ele"] = make([]int32, 0, len(tants))
	for s, constant := range tants {
		fmt.Printf("s:%v \n", s)
		for i, detail := range constant.month {
			fmt.Printf("%v ,detail:%v", i, detail)
			ele := fmt.Sprintf("%v_ele", s)
			water := fmt.Sprintf("%v_water", s)
			total["totalEle"] += detail.element
			total[ele] += detail.element
			total[water] += detail.water
			total["totalWater"] += detail.water
		}
		fmt.Printf("\n")
	}
	fmt.Printf("total:%+v \n", total)
	for s, v := range total {

	}
	//// 保存结果到新文件
	//if err := f.SaveAs("结果文件.xlsx"); err != nil {
	//	log.Fatalf("保存文件失败: %v", err)
	//}
}

type Constant struct {
	uid   string               // 代号
	month map[int32]*MonDetail // 每月
}
type MonDetail struct {
	waterRecord int32 // 水刻度
	water       int32 // 水费
	eleRecord   int32 // 电刻度
	element     int32 // 电费
	monthCost   int32 // 房租

}

func DealValue(f *excelize.File, sheet string) *Constant {
	temp := make(map[int32]*MonDetail)
	constants := &Constant{
		uid: sheet,
	}
	fn := func(str string) int32 {
		if str == "" {
			return 0
		}
		n, err := strconv.Atoi(str)
		if err == nil {
			k := int32(n)
			if k < 0 {
				return 0
			}
			return k
		}
		if strings.Contains(str, "月") {
			strs := strings.Split(str, "月")
			if len(strs) > 1 {
				n, err = strconv.Atoi(strs[0])
				if err == nil {
					return int32(n)
				}
			}
		}
		if strings.Contains(str, ".") {
			floatVal, err := strconv.ParseFloat(str, 32)
			if err != nil {
				fmt.Println("Error converting string to float:", err)
				return 0
			}

			// 将 float64 转换为 float32
			floatVal32 := float32(floatVal)

			// 向下取整 float32 值
			roundedDown := math.Floor(float64(floatVal32))

			// 再次转换为 float32
			down := int32(roundedDown)
			if down < 0 {
				return 0
			}
			return down
		}
		fmt.Println(err)
		return 0
	}
	for i := 4; i < 12; i++ {
		month := fn(f.GetCellValue(sheet, fmt.Sprintf("A%d", i)))
		eleRecord := fn(f.GetCellValue(sheet, fmt.Sprintf("B%d", i)))
		eleCost := fn(f.GetCellValue(sheet, fmt.Sprintf("E%d", i)))
		waterRecord := fn(f.GetCellValue(sheet, fmt.Sprintf("F%d", i)))
		waterCost := fn(f.GetCellValue(sheet, fmt.Sprintf("I%d", i)))
		switch sheet {
		case "二楼":
			eleRecord = fn(f.GetCellValue(sheet, fmt.Sprintf("B%d", i)))
			eleRecord += fn(f.GetCellValue(sheet, fmt.Sprintf("C%d", i)))
			eleCost = fn(f.GetCellValue(sheet, fmt.Sprintf("F%d", i)))
			waterRecord = fn(f.GetCellValue(sheet, fmt.Sprintf("G%d", i)))
			waterCost = fn(f.GetCellValue(sheet, fmt.Sprintf("J%d", i)))
		default:
		}

		temp[month] = &MonDetail{
			water:       waterCost,
			element:     eleCost,
			eleRecord:   eleRecord,
			waterRecord: waterRecord,
			monthCost:   0,
		}
	}

	constants.month = temp
	return constants
}
