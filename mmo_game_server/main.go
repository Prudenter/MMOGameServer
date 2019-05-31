/**
* @Author: ASlowPerson  
* @Date: 19-5-31 下午5:11
*/
package main

import (
	"ZinxServerFramework/zinx/zinxNet"
	"ZinxServerFramework/zinx/zinxInterface"
	"fmt"
)

//当前客户端建立链接之后触发Hook函数
func OnConnectionAdd(conn zinxInterface.InterfaceConnection) {
	fmt.Println("conn Add..")

}

func main() {
	//创建一个zinx server对象
	server := zinxNet.NewServer()

	//注册一些创建链接之后或者销毁链接之前的Hook钩子函数
	server.AddOnConnStart(OnConnectionAdd)

	//启动服务
	server.Run()
}
