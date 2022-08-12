package garten

//keyword 关键字和保留字

const (
	define   = "define"       // 用来定义全局变量，会写入 heap中
	ret      = "return"       //用来返回全局变量
	retLocal = "return_local" //写入局部变量，写入当前Node中的stack中
	sleep    = "sleep"        //休眠一段时间

	//ssh 操作相关的
	ip   = "ip"
	sys  = "sys"
	pwd  = "pwd" //password
	port = "port"
	user = "user"
	//本地命令执行
	run = "run"
	//ssh远程命令执行
	sshRun = "ssh_run"

	//测试相关的关键字
	tool   = "tool"   // 代表使用工具运行
	name   = "name"   // 工具名称
	count  = "count"  // 运行次数
	ignore = "ignore" //屏蔽其它流程
	hook   = "hook"   //hook一些操作，在运行前后
	ip0    = "ip0"
	ip1    = "ip1"
	ip2    = "ip2"
	ip3    = "ip3"
	api    = "api"
)
