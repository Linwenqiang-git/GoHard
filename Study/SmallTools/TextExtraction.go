package smalltools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var BasePaht string = "C:\\work5.0\\xbot-v3\\src\\client\\"

//专用于替换api的error部分
func GetDicFile(fileName string) map[string]string {
	sourceFile := BasePaht + fileName
	f, err := os.Open(sourceFile)
	if err != nil {
		return nil
	}
	defer f.Close()
	fileScanner := bufio.NewReader(f)
	currentKey := ""
	var result map[string]string = make(map[string]string)
	reg := regexp.MustCompile("(\".*[\u4E00-\u9FA5]+)|([\u4E00-\u9FA5]+.*\")")
	for {
		data, _, err := fileScanner.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(data)
		if strings.Contains(line, "public") {
			value, ok := result[currentKey]
			if ok {
				if value == "%%" {
					delete(result, currentKey)
				}
			}
			//去掉async
			line = strings.Replace(line, "async ", "", -1)
			lines := strings.Split(line, " ")
			funcIndex := 0
			for index, name := range lines {
				if name == "public" {
					funcIndex = index + 2
				}
			}
			funcName := strings.Split(lines[funcIndex], "(")[0]
			result[funcName] = "%%"
			currentKey = funcName
		}
		//没有替换的部分
		if strings.Contains(line, "error = $") && !strings.Contains(line, "Strings.RobotApi_") {
			matchData := reg.FindAllStringSubmatch(line, -1)
			if len(matchData) > 0 {
				result[currentKey] = matchData[0][0][1:]
			}
		}
	}
	print(len(result))
	print("\n")
	return result
}

func UpdateSourceFile(dic map[string]string, fileType, fileName, filePath string) {
	paht := BasePaht + filePath
	data := ReadAll(paht)
	prefix := ""
	if fileType == "cs" {
		prefix = "{Strings." + fileName + "_"
	} else {
		prefix = "{x:Static resx:Strings." + fileName + "_"
	}
	for k, v := range dic {
		key := prefix + k + "}"
		data = strings.Replace(data, v, key, -1)
	}
	err := WriteToFile(paht, data)
	if err != nil {
		fmt.Print(err)
	}
}

//设置首字母大写
func setFisrtLetters(data string) string {
	b := data[0] - 32
	result := string(b) + data[1:]
	return result
}
