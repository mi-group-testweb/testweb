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

type ReturnType struct {
	Time        string  `json:"time"`
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
	s2 := rand.NewSource(time.Now().Unix())
	r2 := rand.New(s2)
	curtime := time.Now().Format("2006-01-02 15:04:05")
	returntype := ReturnType{
		Time:        curtime,
		IP:          need_ip.IP,
		DnsTime:     r2.Float64(),
		TcpTime:     r2.Float64(),
		SslTime:     r2.Float64(),
		FirstTime:   r2.Float64(),
		LoadTime:    r2.Float64(),
		RequestTime: r2.Float64(),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(returntype); err != nil {
		panic(err)
	}
}
