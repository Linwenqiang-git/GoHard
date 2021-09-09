package tool

import "reflect"

//两个struct 相同字段映射
func Mapper(source, target interface{}) interface{} {
	source_type := reflect.TypeOf(source) //获取type
	target_type := reflect.TypeOf(target)

	target_result := reflect.ValueOf(target).Elem() //获取value
	source_result := reflect.ValueOf(source).Elem()
	//双循环，对相同名字对字段进行赋值
	for i := 0; i < target_type.NumField(); i++ {
		for j := 0; j < source_type.NumField(); j++ {
			if target_type.Field(i).Name == source_type.Field(j).Name {
				target_result.Field(i).Set(source_result.Field(j))
			}
		}
	}
	return target_result
}
