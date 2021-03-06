package reponsetory

import (
	"database/sql"
	"encoding/json"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	tool "github.linwenqiang.com/GinStudy/Tool"
	"go.uber.org/dig"
)

type DbContext struct {
	dig.In
	Context *sql.DB
}

/*============================================继承IDbContext============================================*/
func (db *DbContext) SetConnect(conStr string) {
	context, err := sql.Open("mysql", conStr)
	tool.PanicError(err)
	db.Context = context
}

/*常用数据库操作*/
//返回受影响行数
func (this *DbContext) NonQueryReturnRowCount(sql string) int64 {
	result, err := this.Context.Exec(sql)
	if checkErr(err) {
		return 0
	}
	rowCount, err := result.RowsAffected()
	if checkErr(err) {
		return 0
	}
	return rowCount
}

//返回自增ID
func (this *DbContext) NonQueryReturnLastId(sql string) int64 {
	result, err := this.Context.Exec(sql)
	if checkErr(err) {
		return 0
	}
	lastId, err := result.LastInsertId()
	if checkErr(err) {
		return 0
	}
	return lastId
}

//查询类，需要指定返回的实体
func (this *DbContext) Query(sqlstr string, obj interface{}) []interface{} {

	rows, err := this.Context.Query(sqlstr)
	checkErr(err)
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	checkErr(err)
	//sql.RawBytes是byte[]类型的别名，columnValues得到的就是[]byte类型的切片
	columnValues := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columnValues))

	for i := range columnValues {
		scanArgs[i] = &columnValues[i]
	}
	result := []interface{}{}
	for rows.Next() {
		// get RawBytes from data
		checkErr(rows.Scan(scanArgs...))

		entityValue := reflect.ValueOf(obj).Elem()

		for i, col := range columnValues {
			columnName := columns[i]
			if col == nil {
				continue
			} else {
				fieldTypeKind := entityValue.FieldByName(columnName).Type().Kind()
				//反射根据字段类型设置值
				if fieldTypeKind == reflect.String {
					entityValue.FieldByName(columnName).SetString(string(col))
				} else if fieldTypeKind == reflect.Bool {
					var value bool
					json.Unmarshal(col, &value)
					entityValue.FieldByName(columnName).SetBool(value)
				}
			}
		}
		//基于初始大小后面append，所以切片默认大小为0即可
		result = append(result, entityValue.Interface())
	}
	return result
}

/*============================================继承IDbContext============================================*/

/*============================================结构体内私有方法============================================*/
//错误异常处理
func checkErr(err error) bool {

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return false
}

//通过类型来动态创建实例
func getTypeInstance(_type reflect.Type) interface{} {
	_value := reflect.New(_type) //创建一个指向_type类型的指针,本身为val类型
	vl := _value.Elem()          //获取指针指向的_value的值
	//vl.Field(0).SetString("lalala")	//给第一个值赋值
	_instance := vl.Interface() //将当前对象转换成一个通用类型并返回
	return _instance
}

//判断参数类型是否是切片
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

//生成[]interface{}切片，因为interface{}是通用类型的，但是[]interface{}不是
func createAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := isSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

/*============================================结构体内私有方法============================================*/
