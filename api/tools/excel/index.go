package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

type ExcelUtils struct {
}

func (receiver ExcelUtils) CreateFile() *excelize.File {
	f := excelize.NewFile()
	return f
}

/**
    // 头部
	header := []map[string]string{
		{"label":"姓名","key":"name"},
		{"label":"学号","key":"student_id"},
		{"label":"性别","key":"sex"},
	}
	// 数据源
	data := []map[string]interface{}{
		{"name":"jack","student_id":18,"sex":"男"},
		{"name":"mary","student_id":28,"sex":"女"},

*/

func (receiver ExcelUtils) Export(header []map[string]string, data []map[string]interface{}) (*excelize.File, error) {
	f := receiver.CreateFile()
	index := f.NewSheet("Sheet1")

	for clumnNum, v := range header {
		sheetPosition := Div(clumnNum+1) + "1"
		f.SetCellValue("Sheet1", sheetPosition, v["label"])
	}

	for lineNum, v := range data {
		clumnNum := 0
		for _, v1 := range header {
			clumnNum++
			sheetPosition := Div(clumnNum) + strconv.Itoa(lineNum+2)
			key := v1["key"]
			value := v[key]
			switch value.(type) {
			case string:
				f.SetCellValue("Sheet1", sheetPosition, value.(string))
			case int:
				f.SetCellValue("Sheet1", sheetPosition, value.(int))
			case float64:
				f.SetCellValue("Sheet1", sheetPosition, value.(float64))
			default:
				f.SetCellValue("Sheet1", sheetPosition, value)
			}
		}

	}

	f.SetActiveSheet(index)
	if err := f.SaveAs("test.xlsx"); err != nil {
		return f, err
	}
	return f, nil
}

// Div 数字转字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
