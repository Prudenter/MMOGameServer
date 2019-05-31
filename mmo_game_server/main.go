/**
* @Author: ASlowPerson  
* @Date: 19-5-31 下午5:11
*/
package main

import (
	"ZinxServerFramework/zinx/zinxNet"
)


func main() {
	//创建一个zinx server对象
	server := zinxNet.NewServer()

	//注册一些创建链接之后或者销毁链接之前的Hook钩子函数

	//启动服务
	server.Run()
}
