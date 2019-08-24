package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ReturnType struct {
	Time        int64   `json:"time"`
	IP          string  `json:"ip"`
	DnsTime     float64 `json:"time_dns"`
	TcpTime     float64 `json:"time_tcp"`
	SslTime     float64 `json:"time_ssl"`
	FirstTime   float64 `json:"time_firstly"`
	LoadTime    float64 `json:"time_load"`
	RequestTime float64 `json:"time_total"`
}

type TestURL struct {
	IP string `json:"ip"`
}

func main() {
	http.HandleFunc("/test", Hand)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func Hand(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var need_ip TestURL
	err = json.Unmarshal(body, &need_ip)
	if err != nil {
		panic(err)
	}
	curtime := time.Now().Unix()
	returntype := ReturnType{
		Time:        curtime,
		IP:          need_ip.IP,
		DnsTime:     0.33,
		TcpTime:     0.23,
		SslTime:     0.454,
		FirstTime:   0.01,
		LoadTime:    0.3,
		RequestTime: 0.13,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(returntype); err != nil {
		panic(err)
	}
}
