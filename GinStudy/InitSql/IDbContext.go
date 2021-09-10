package initsql

type IDbContext interface {
	//设置数据库上下文
	SetDbConnect()

	NonQueryReturnRowCount(sql string) int64

	NonQueryReturnLastId(sql string) int64

	Query(sql string, obj interface{}) []interface{}
}
