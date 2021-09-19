package dependentinjection

import (
	. "github.linwenqiang.com/GinStudy/InitSql"
	"go.uber.org/dig"
)

type Config struct {
	/*内嵌了dig.In之后，dig会将该类型中的其它字段看成对象的依赖，
	  创建Object类型的对象时，会先将依赖的Arg1/Arg2/Arg3/Arg4创建好。*/
	dig.In
	MySQL *MySqlDbConfig
}
