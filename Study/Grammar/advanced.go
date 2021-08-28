package basekgrammar

import (
	"encoding/json"
	. "fmt"
	"reflect"
)

/*
空接口、反射
对应c#的泛型、反射
*/

/*空接口  interface{} 万能类型 对应c#里的泛型 或者 object类型  所有类型的父类*/
func CallCommonType(arg interface{}) {
	Println("Call CallCommonType...")
	Println("arg type is ", reflect.TypeOf((arg)))   //反射获取对象类型
	Println("arg value is ", reflect.ValueOf((arg))) //反射获取对象的值
	//interface{} 提供 类型断言 机制
	value, ok := arg.(string)
	if ok {
		Println("arg is string value = ", value)
	}
}

/*
反射
pair<type,value>
变量的结构： {
	type：{
		static type: int,string...
		concrete type ： interface所指向的具体类型
	}，
	value：
}
*/
type Read interface {
	ReadBook()
}
type Write interface {
	WriteBook()
}
type Essay struct {
	Name string
	Page int
}

func (this *Essay) ReadBook() {
	Println("ReadBook...")
}
func (this *Essay) WriteBook() {
	Println("WriteBook...")
}

//反射遍历类里面的方法和字段
func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	Println("input type :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	Println("input value :", inputValue)

	//获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	//获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		Printf("%s : %v\n", m.Name, m.Type)
	}
}

/*
	结构体标签 struct tag
	key:value 形式  主要用于包说明
	encoding/json 用于结构体(类)转换成json格式数据

*/
type rensume struct {
	Name string `json:"name" doc:"名字"`
	Age  int    `json:"age"`
	Info string `json:"info"`
}

func FindTag(str rensume) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagstr := t.Field(i).Tag.Get("json")
		if tagstr != "" {
			Println("json:", tagstr)
		}

	}
}
func GetJsonData(str interface{}) {
	jsonstr, err := json.Marshal(str)
	if err != nil {
		Println("json.Marshal error :", err)
	}
	Printf("%s\n", jsonstr)
	JsonToStruct(jsonstr)
}
func JsonToStruct(jsonstr []byte) {
	myresume := rensume{}
	err := json.Unmarshal(jsonstr, &myresume)
	if err != nil {
		Println("json.Unmarshal error :", err)
	}
	Printf("%v\n", myresume)
}

func init() {
	// Println("advance init...")
	// //断言
	// CallCommonType("sss")

	// //b pair<type:Essay,value:Essay{}地址>
	// b := &Essay{Name: "张三", Page: 10}

	// //r pair<type:,value:>
	// var r Read
	// //r pair<type:MathBook,value:MathBook{}地址>
	// r = b

	// r.ReadBook()
	// Println("======================================================")
	// //DoFieldAndMethod(Essay{Name: "李四", Page: 10})
	// var tag rensume = rensume{Name: "zhangsan", Age: 10, Info: "json info"}
	// //FindTag(&tag)
	// GetJsonData(tag)
	// Println("======================================================")

	// var w Write
	// w = r.(Write) //此处能断言 是因为 r和w类型一致

}
