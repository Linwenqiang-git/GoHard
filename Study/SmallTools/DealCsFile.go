package smalltools

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type row struct {
	key_zh         string
	key_en         string
	prefix         string
	file_full_path string
	mark           string
}

//写入excel
func write_to_excel(rows []*row, sheetName string) {
	for _, row := range rows {
		print(row.key_zh + "  " + row.prefix + "\n")
	}
	Rows_to_excel(rows, sheetName)
}

//获取该文件的中文部分
func get_file_zh(fileName string) []*row {
	no_repeat := make(map[string]string)
	file, err := os.Stat(fileName)
	if err != nil {
		println("读取文件信息失败")
		return nil
	}
	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()
	fileScanner := bufio.NewReader(f)
	ret := []*row{}
	reg := regexp.MustCompile("(\".*[\u4E00-\u9FA5]+)|([\u4E00-\u9FA5]+.*\")")
	for {
		data, _, err := fileScanner.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(data)
		if strings.Contains(line, "///") {
			continue
		}
		if strings.Contains(line, "//") {
			line = strings.Split(line, "//")[0]
		}
		matchData := reg.FindAllStringSubmatch(line, -1)
		if matchData != nil && len(matchData) > 0 {
			zh := strings.Replace(matchData[0][0], "\"", "", -1)
			prefix := strings.Replace(file.Name(), ".cs", "", -1) + "_"
			_, ok := no_repeat[prefix+zh]
			if !ok {
				rowInfo := &row{
					key_zh:         zh,
					prefix:         prefix,
					key_en:         "",
					file_full_path: fileName,
					mark:           "",
				}
				if strings.Contains(line, "Logging") {
					rowInfo.mark = "Logging"
				}
				ret = append(ret, rowInfo)
				no_repeat[prefix+zh] = zh
			}
		}
	}
	return ret
}

//遍历该项目下所有的文件
func Fraversal_folder_all_file(pathname, sheetName string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := Fraversal_folder_all_file(fullname, sheetName)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			if !strings.Contains(fullname, "obj") {
				ext := filepath.Ext(fullname)
				if ext == ".cs" && !strings.Contains(fi.Name(), "Strings") {
					result = append(result, fullname)
				}
			}
		}
	}
	rows := []*row{}
	for _, file := range result {
		rows = append(rows, get_file_zh(file)...)
	}
	write_to_excel(rows, sheetName)
	return result, nil
}

//将整理好的excel写入原始文件
func UpdateCSSourceFile() {
	WaitDealFile := map[string]string{
		//"ShadowBot.UIAutomation": "UIAutomation",
		"ShadowBot.UIAutomation.Tools": "Tools",
	}
	for _, v := range WaitDealFile {
		rows := ReadExcel(v)
		for _, data := range rows {
			content := ReadAll(data.file_full_path)
			if !strings.Contains(content, "using ShadowBot.UIAutomation.Tools.LocalizationResources;") {
				content = "using ShadowBot.UIAutomation.Tools.LocalizationResources;\n" + content
			}
			to_replace := data.key_zh
			replace_text := ""
			if data.mark == "Logging" {
				replace_text = "$$Log" + data.key_en //这里打一个特殊标记，方便后面修改
			} else {
				replace_text = "{Strings." + data.key_en + "}"
			}
			new_content := strings.Replace(content, to_replace, replace_text, -1)
			err := WriteToFile(data.file_full_path, new_content)
			if err != nil {
				fmt.Print("文件写入出错" + data.file_full_path)
			}
		}
	}
	print("finish")
}
