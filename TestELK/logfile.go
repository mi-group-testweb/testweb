package main
import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"syscall"
	"time"
)

func writeLog(elkItem interface{}, logDir string) {
	logTime := time.Now().Format("20060102")

	//创建文件夹
	err := os.MkdirAll(logDir, 0777)
	if err != nil {
		log.Fatalln("logFile dir create fail:", err)
	}
	filename := logDir + logTime +"app.log"
	//打开/新建日志文件
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer func() {
		err := logfile.Close()
		if err != nil {
			log.Println("logfile close fail:", err)
		}
	}()

	if err != nil {
		log.Println("logfile open fail:", err)
		return
	}

	// 文件锁
	err = syscall.Flock(int(logfile.Fd()), syscall.LOCK_EX)
	defer func() {
		err := syscall.Flock(int(logfile.Fd()), syscall.LOCK_UN)
		if err != nil {
			log.Println("unlock logfile fail:", err)
		}
	}()


	if err != nil {
		log.Println("add lock to logfile fail:", err)
		return
	}

	//将无效数据以json格式写入日志文件，阻止HTML特殊字符编码
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err = jsonEncoder.Encode(elkItem)
	if err != nil {
		log.Println("encode json fail: ", err)
		return
	}

	if _, err := logfile.WriteString(bf.String()); err != nil {
		log.Println("write to log file fail:", err)
		return
	}
}

func WriteLogToElk(timestamp int64, Num int64) {
	var testElk Test

	testElk.Timestamp = timestamp
	testElk.Num = Num

	logDir := "/home/work/log/"
	writeLog(testElk, logDir)
}
