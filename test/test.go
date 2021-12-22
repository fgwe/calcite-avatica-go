package main

import (
	"conn-lindorm/mapper"
	"fmt"
	_ "github.com/fgwe/calcite-avatica-go/v5"
)

func main() {
	var u mapper.UserManager = &mapper.UserManagerImpl{}
	//err := u.Connect("http://ld-t4n52235c22n4dym9-proxy-lindorm-pub.lindorm.rds.aliyuncs.com:30060/test?authentication=DIGEST&avaticaUser=root&avaticaPassword=root")
	err := u.Connect("http://localhost:30060/test?user=root&password=root")
	if err != nil {
		fmt.Println("connect error:", err.Error())
		return
	}
	_ = u.Add(1, "zhangsan", 17)
	_ = u.Add(2, "lisi", 18)
	_ = u.Add(3, "wanger", 19)
	_ = u.Delete(2)
	_ = u.Update(1, "zhangsan2", 27)
	fmt.Println(u.Get(1))
}
