package test

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"testing"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func TestJava(t *testing.T) {

	params := []string{"-jar", "ocicat-afpj-master.jar", "-p", "afpj.properties", "s"}
	//params2 := []string{"-jar", "ocicat-afpj-master.jar", "-p", "afpj.properties", "i"}
	cmd := exec.Command("java", params...)
	cmd.Dir = "D:\\78240749\\ocicat\\"
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("error:", err.Error())
	}
	//fmt.Printf("end: %q\n", string(out))
	garbledStr := ConvertByte2String(out, GB18030)
	fmt.Println("-----------------")
	fmt.Println(garbledStr)

	if strings.Contains(garbledStr, "命中次数hit: 20001") {
		fmt.Println("YES")
	}

	//清空文件夹
	//stdout, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//_ = cmd.Start()
	//in := bufio.NewScanner(string(stdout))
	//for in.Scan() {
	//	cmdRe:=ConvertByte2String(in.Bytes(),"GB18030")
	//	fmt.Println(cmdRe)
	//}
	//_ = cmd.Wait()
}

//func main() {
//	command := "ping"
//	params := []string{"127.0.0.1","-t"}
//	cmd := exec.Command(command, params...)
//
//}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
