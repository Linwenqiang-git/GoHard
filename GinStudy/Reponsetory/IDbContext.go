package reponsetory

type IDbContext interface {
	NonQueryReturnRowCount(sql string) int64

	NonQueryReturnLastId(sql string) int64

	Query(sql string, obj interface{}) []interface{}
}
