package main
import (
	"net"
	"bytes"
	"encoding/binary"
	//"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"os"
	"time"
//	"github.com/sparrc/go-ping"
	//"flag"
	//"io"
	//"math"
	//"code.google.com/p/mahonia"
	"log"
	//"github.com/json-iterator/go"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/wonderivan/logger"
)

type (
    Dist struct{
    Available string  `json:"Available"`
	Size string    `json:"Size"`
}
 cipan struct{
	Cached int  `json:"cached"`
	Free int   `json:"free"`
	Total int   `json:"total"`
	Used int    `json:"used"`
}
 Cpu struct{
	Time string   `json:"time"`
	Use float64   `json:"use"`
}
Data struct{ 
  CPU []Cpu   `json:"CPU"`
  Disk Dist   `json:"Disk"`
  Mem  cipan   `json:"Mem"`
}
  receive struct{
    DATA Data   `json:"data"`
    //code int  
    Message string  `json:"message"`
}
)
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}
func usage() {
	msg := `
Need to run as root!

Usage:
	goping host

	Example: ./goping www.baidu.com`

	fmt.Println(msg)
	os.Exit(0)
}

func convertToBin(num int) string {
    s := ""
    
    if num == 0 {
        return "0"
    }    
 
    // num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
    for ;num > 0 ; num /= 2 {
        lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
        s = strconv.Itoa(lsb) + s
    }
    return s
}

func main() {

	httpGet()
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
func httpGet() {
	fileName:="ll.log"
	logFile,err:=os.Create(fileName)
	defer logFile.Close()
	if err!=nil{
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile,"[info]",log.LstdFlags)
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	var shuju receive
	var (
		icmp ICMP
	laddr    = net.IPAddr{IP: net.ParseIP("192.168.201.117")}
	//raddr=net.res
	raddr = net.IPAddr{IP: net.ParseIP("192.168.201.123")}
	
	)
	conn, err := net.DialIP("ip4:icmp", &laddr, &raddr)	
	

	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()
	if err != nil {
		fmt.Printf("%s",err)
		return 
	}
	icmp.Type = 8
	icmp.Code = 0
	icmp.CheckSum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)
	fmt.Printf("\n正在 Ping %s 具有 0 字节的数据:\n", raddr.String())
	recv := make([]byte, 1024)
	sended_packets:=0

	
	//f, err := os.OpenFile(name, os.O_WRONLY, 0644)
	TimerDemo := time.NewTimer(time.Duration(3) * time.Second)
	f,err:=os.OpenFile("E:/zuixin/testwritefile.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if err!=nil{
		fmt.Printf("FAIL to open %s\n",err)
	}
	//i:=1
	for {
		resp, err := http.Get("http://192.168.201.123:8002/esb/supervision_handler/devinfo")
		if err != nil {
            resp.Body.Close()
        }
		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
		  fmt.Printf("FAIL%s\n",err)
		}
		select {
		case <-TimerDemo.C:
			if _, err := conn.Write(buffer.Bytes()); err != nil {
				fmt.Println(err.Error())
				return
			}
			sended_packets++
			t_start:=time.Now()

conn.SetReadDeadline((time.Now().Add(time.Second * 2)))
_, err := conn.Read(recv)

		if err != nil {
			fmt.Println("请求超时")
			f.WriteString("发送失败")
		}else {
			t_end:=time.Now()
			dur:=t_end.Sub(t_start).Nanoseconds()
			fmt.Printf("发送成功%s 的回复时间为 = %dms\n", raddr.String(), dur/1000000)
			k:=raddr.String()+" 的回复时间为 = "
			p:=strconv.Itoa(int(dur/1000000))
			c:="ms"	
			f.WriteString(k)
			f.WriteString(string(p))
			f.WriteString(c)
//i++
		}
		if err := json.Unmarshal(body, &shuju);err !=nil{
			fmt.Println(err)
			return
			}
			flag:=0
			for key,_:=range shuju.DATA.CPU{
				flag=key
			}
			shuju.DATA.CPU[0]=shuju.DATA.CPU[flag]
			shuju.DATA.CPU=append(shuju.DATA.CPU[:1])
		   jsonstu,err:= json.Marshal(shuju)
		   if err!=nil{
			fmt.Printf("FAIL%s\n",err)
		  }
//value:=gjson.Get(string(jsonstu),"data.CPU[0].time")
//f.WriteString("时间为:")
//f.WriteString(shuju.DATA.CPU[0].Time)
//flag1:=gjson.Get(string(jsonstu),"data.CPU[0].use")
s2 := strconv.FormatFloat(shuju.DATA.CPU[0].Use, 'E', -1	, 64)
//f.WriteString("	CPU:")
//f.WriteString(s2)
flag2:=gjson.Get(string(jsonstu),"data.Disk.Available")
//f.WriteString("	磁盘:")
//f.WriteString(flag2.String())
flag3:=gjson.Get(string(jsonstu),"data.Disk.Size")
//f.WriteString("/")
//f.WriteString(flag3.String())
flag4:=gjson.Get(string(jsonstu),"data.Mem.free")
//f.WriteString("	内存:")
//f.WriteString(flag4.String())
//f.WriteString("/")
flag5:=gjson.Get(string(jsonstu),"data.Mem.total")
//f.WriteString(flag5.String()+"	")
logger.Info("CPU："+s2+"  磁盘:"+flag2.String()+"/"+flag3.String()+"  内存:"+flag4.String()+"/"+flag5.String()+"	")
debugLog.Println("CPU："+s2+"  磁盘:"+flag2.String()+"/"+flag3.String()+"  内存:"+flag4.String()+"/"+flag5.String()+"	")
		/*  f.WriteString(string(jsonstu))
	fmt.Fprintln(f,"")*/
		//	f.Write([]byte(shuju))
			//if err != nil {
			//	log.Fatal(err)
				//   }
			//超时后重置定时器
			TimerDemo.Reset(time.Duration(3) * time.Second)
		}
	}
	//fmt.Println(string(body))
	//f1 := &multitypeTest{
	//  One:"a",
	//  Two:"b",
	//  }
	//f1.Showmul()
	//fjson1, err := json.Marshal(f1)
	//fmt.Println(string(fjson1))
}
func getICMP(seq uint16) ICMP {
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()

	return icmp
}
func CheckError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
