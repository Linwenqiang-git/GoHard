package tool

import "fmt"

//将错误以panic的形式报出
func PanicError(err error, info ...string) {
	if err != nil {
		if info == nil {
			panic(err)
		} else {
			var tips string
			for _, str := range info {
				tips += str
			}
			panic(tips + ":" + err.Error())
		}
	}
}

//将错误信息打印出来
func PrintInfoError(err error, info ...string) {
	if err != nil {
		if info == nil {
			fmt.Println(err)
		} else {
			var tips string
			for _, str := range info {
				tips += str
			}
			fmt.Printf("%s:%s", tips, err.Error())
		}
	}
}
