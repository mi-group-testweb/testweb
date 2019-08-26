package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Assignment struct {
	RequestInterval float64 `json:"request_interval"`
	Tasks           []Task  `json:"tasks"`
}

type TaskId struct {
	Id               int    `json:"id"`
	MissionId        int    `json:"mission_id"`
	MissionIntanceId string `json:"mission_instance_id"`
	MissileId        int    `json:"missile_id"`
}

type Task struct {
	TaskId
	Type       string          `json:"type"`
	AssignedAt string          `json:"assigned_at"`
	RunAt      string          `json:"run_at"`
	IpVersion  int             `json:"ip_version"`
	Bullet     json.RawMessage `json:"bullet"`
}

type HTTPBullet struct {
	Url         string                 `json:"url"`
	Method      string                 `json:"method"`
	Header      map[string]interface{} `json:"header"`
	ContentType string                 `json:"content_type"`
	Body        string                 `json:"body"`
	MatchResult struct {
		MatchHeader  bool   `json:"match_header"`
		MatchBody    bool   `json:"match_body"`
		MatchContent string `json:"match_content"`
	} `json:"match_result"`
	ReturnHeader bool `json:"return_header"`
	Redirect     bool `json:"redirect"`
	Compress     bool `json:"compress"`
	Timeout      int  `json:"timeout"`
	Debug        struct {
		Ping string `json:"ping"`
		Mtr  string `json:"mtr"`
		Dig  string `json:"dig"`
	} `json:"debug"`
	Resolve struct {
		Enabled bool   `json:"enabled"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Address string `json:"address"`
	} `json:"resolve"`
	Ldns struct {
		Enabled bool   `json:"enabled"`
		Server  string `json:"server"`
	} `json:"ldns"`
}

type StackData struct {
	TaskId
	Err
	Type  string          `json:"type"`
	Raw   string          `json:"raw"`
	Track json.RawMessage `json:"track"`
}

type Err struct {
	Message   string `json:"msg"`
	Code      int    `json:"code"`
	ErrorType string `json:"error_type"`
	Error     bool   `json:"error"`
	ParseErr  string `json:"-"`
}

type DataArr struct {
	Data []StackData `json:"data"`
}

type HTTPTrack struct {
	HTTPTime
	HTTPTimeSlice
	ContentType  string  `json:"content_type"`
	Speed        float64 `json:"speed"`
	ConnectNum   float64 `json:"num_connects"`
	RedirectUrl  string  `json:"redirect_url"`
	RedirectNum  float64 `json:"num_redirects"`
	DownloadSize float64 `json:"size_download"`
	ServerIP     string  `json:"server_ip"`
	// UrlEffective string  `json:"url_effective"`
	Ldns         string `json:"ldns"`
	IsMatch      bool   `json:"is_match"`
	Header       string `json:"header"`
	ResponseCode int    `json:"rc"`
	HTTPDebug
	Load string `json:"load"`
	At
}

type HTTPTime struct {
	DNSResolveTime  float64 `json:"time_namelookup"`
	ConnectionTime  float64 `json:"time_connect"`
	RedirectTime    float64 `json:"time_redirect"`
	AppConnectTime  float64 `json:"time_appconnect"`
	PretransferTime float64 `json:"time_pretransfer"`
	FirstByteTime   float64 `json:"time_starttransfer"`
	ResponseTime    float64 `json:"time_total"`
}

type HTTPTimeSlice struct {
	Dnsrt float64 `json:"dnsrt"`
	Ct    float64 `json:"ct"`
	Rdt   float64 `json:"rdt"`
	Act   float64 `json:"act"`
	Fbt   float64 `json:"fbt"`
	Cdt   float64 `json:"cdt"`
	Rbt   float64 `json:"rbt"`
	Rt    float64 `json:"rt"`
}

type HTTPDebug struct {
	Url   string `json:"url"`
	Debug struct {
		Ping string `json:"ping"`
		Dig  string `json:"dig"`
		Mtr  string `json:"mtr"`
	} `json:"debug"`
}

type At struct {
	AssignedAt  string `json:"assigned_at"`
	RunAt       string `json:"run_at"`
	CollectedAt string `json:"collected_at"`
}

type TestURL struct {
	IP string `json:"ip"`
}

type ReturnType struct {
	Time          string  `json:"time"`
	IP            string  `json:"ip"`
	DnsTime       float64 `json:"time_dns"`
	TcpTime       float64 `json:"time_tcp"`
	SslTime       float64 `json:"time_ssl"`
	FirstByteTime float64 `json:"time_firstly"`
	LoadTime      float64 `json:"time_load"`
	TotalTime     float64 `json:"time_total"`
	DNSNote       string  `json:"note_dns"`
	TCPNote       string  `json:"note_tcp"`
	SSLNote       string  `json:"note_ssl"`
	FirstByteNote string  `json:"note_firstly"`
	LoadNote      string  `json:"note_load"`
	TotalNote     string  `json:"note_total"`
}

var urlstream chan string
var returnstream chan ReturnType

//var flag int

func main() {
	urlstream = make(chan string, 5)
	returnstream = make(chan ReturnType, 5)
	http.HandleFunc("/url", GetUrl)
	http.HandleFunc("/v2/assignments", SendMission)
	http.HandleFunc("/v2/tracks", GetResult)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8234", nil))
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var NeedIP TestURL
	//从请求中解析出需要测试的站点IP
	err = json.Unmarshal(body, &NeedIP)
	if err != nil {
		panic(err)
	}
	TestIP := NeedIP.IP
	urlstream <- TestIP
	returnvalue, _ := <-returnstream
	returnvalue.IP = TestIP
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(returnvalue); err != nil {
		panic(err)
	}
}

func SendMission(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get mission request")
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(10000)
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	assignment := Assignment{
		RequestInterval: 600,
	}
	task := Task{
		TaskId: TaskId{
			Id:               count,
			MissionId:        count,
			MissionIntanceId: "",
			MissileId:        count,
		},
		Type:      "HTTP",
		IpVersion: 4,
	}
	var bullet HTTPBullet
	TestIP, _ := <-urlstream
	bullet.Url = TestIP
	//bullet.Url = "https://www.xiaomi.com"
	bullet.Method = "POST"
	bullet.Redirect = false
	bullet.Timeout = 60
	bullet.ContentType = "raw"
	bullet.ReturnHeader = false
	byte1, err := json.Marshal(bullet)
	if err != nil {
		panic(err)
	}
	task.Bullet = byte1
	assignment.Tasks = append(assignment.Tasks, task)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(assignment); err != nil {
		panic(err)
	}
	//time.Sleep(30 * time.Second)
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Data")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var stack DataArr
	fmt.Println(string(body))
	err = json.Unmarshal(body, &stack)
	if err != nil {
		panic(err)
	}
	var HTTPResult HTTPTrack
	err = json.Unmarshal(stack.Data[0].Track, &HTTPResult)
	fmt.Printf(" DNSTime : %+v\n", HTTPResult.DNSResolveTime)
	var tmp ReturnType
	tmp.Time = time.Now().Format("2006-01-02 15:04:05")
	tmp.DnsTime = HTTPResult.DNSResolveTime
	tmp.TcpTime = HTTPResult.ConnectionTime
	tmp.SslTime = HTTPResult.AppConnectTime
	tmp.FirstByteTime = HTTPResult.FirstByteTime
	tmp.LoadTime = HTTPResult.ResponseTime - HTTPResult.PretransferTime
	tmp.TotalTime = HTTPResult.ResponseTime
	tmp.DNSNote = "通过域名解析服务（DNS），将指定的域名解析成IP地址的消耗时间"
	tmp.TCPNote = "建立到服务器的TCP连接所用的时间"
	tmp.SSLNote = "SSL握手时间"
	tmp.FirstByteNote = "在发出请求之后，Web 服务器返回数据的第一个字节所用的时间"
	tmp.LoadNote = "客户机从服务器下载数据所用的时间"
	tmp.TotalNote = "在发出请求之后，Web 服务器处理请求并开始发回数据所用的时间是"
	returnstream <- tmp
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
