package initsql

import (
	"database/sql"
	"log"

	. "github.com/go-sql-driver/mysql"
	. "github.linwenqiang.com/GinStudy/Entity"
)

//原生连接mysql
func ConnetSql() *sql.DB {
	connstr := "root:LWQlwq123@@tcp(127.0.0.1:3306)/golangstudy"
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return db
}

/*常用数据库操作*/
func NonQuery(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}

func Query(db *sql.DB, sql string) []*UserEntity {
	result := make([]*UserEntity, 10)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err.Error())
		return result
	}
	//相当于while 循环
scan:
	if rows.Next() {
		userEntity := new(UserEntity)
		err := rows.Scan(&userEntity.UserId, &userEntity.UserName, &userEntity.UserPhoto)
		if err != nil {
			log.Fatal(err.Error())
			return result
		}
		result = append(result, userEntity)
		goto scan
	}
	return result
}
