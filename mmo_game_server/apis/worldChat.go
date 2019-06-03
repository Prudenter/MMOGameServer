/**
* @Author: ASlowPerson  
* @Date: 19-6-3 下午10:08
*/

package apis

import (
	"ZinxServerFramework/zinx/zinxNet"
	"ZinxServerFramework/zinx/zinxInterface"
	"MMOGameServer/mmo_game_server/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
	"MMOGameServer/mmo_game_server/core"
)

/*
	世界聊天路由业务
*/
type WorldChat struct {
	zinxNet.ZinxRouter
}

//定义用来处理业务的方法
func (wc *WorldChat)Handle(request zinxInterface.InterfaceRequest){
	//1.解析客户端传递进来的protobuf数据
	proto_msg := &pb.Talk{}
	if err := proto.Unmarshal(request.GetMessage().GetMsgData(), proto_msg);err!=nil{
		fmt.Println("Talk message unmarshal error",err)
		return
	}
	//通过获取链接属性,得到当前的玩家ID
	pid,err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("get pid err:",err)
		return
	}

	//通过pid,来得到对应的player对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	//将当前的聊天数据广播给全部的在线玩家
	player.SendTaldMsgToAll(proto_msg.GetContent())
}