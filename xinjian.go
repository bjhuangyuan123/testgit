package main
import "fmt"
/*
#include<string.h>
#include<stdlib.h>
#include<stdio.h>
#include"videoma_api.h"
*/
import "C"

func main(){
	var a C.ConfigParam
	var c C.GbConfigParam
	a.gbCfgParam=c
   c.szSvrID[0]=C.char('1')
   fmt.Println(c.szSvrID[0])
   c.heartbeatCycleTime=3600
   c.maxHeartbeatCount=3
   //c.szLocalIP=C.string ("192.168.201.1")
	b:="192.168.201.1"
	//str:=[]rune(b)
	var p [17]C.char
	var str *C.char
	*str=p[0]
	C.strcpy(str,C.CString(b))
	//C.strcpy(p,"192.168.201.1")
	c.szLocalIP=p
}
