package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

type PINGBullet struct {
	Host      string `json:"host"`
	Count     int    `json:"count"`
	DataBytes int    `json:"data_bytes"`
	Interval  int    `json:"interval"`
	Debug     struct {
		Mtr string `json:"mtr"`
	} `json:"debug"`
}

type TestURL struct {
	IP string `json:"ip"`
}

func main() {
	http.HandleFunc("/v2/mission_instance", Save_Mission)
	//http.HandleFunc("/v2/assignments", Send_Mission)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func Save_Mission (w http.ResponseWriter, r *http.Request) {
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
	_, _ = fmt.Fprintln(w, "Unmarshal successefully!")
	task := Task{
		TaskId: TaskId{
			Id:               0,
			MissionId:        0,
			MissionIntanceId: "",
			MissileId:        0,
		},
		Type: "PING",
	}
	var bullet PINGBullet
	bullet.Host = NeedIP.IP
	bullet.Count = 4
	byte1, err := json.Marshal(bullet)
	if err != nil {
		panic(err)
	}
	task.Bullet = byte1
	err = json.Unmarshal(task.Bullet, &bullet)
	_, _ = fmt.Fprintf(w, "type : %+v\n", task)
	_, _ = fmt.Fprintf(w, "bullet : %+v\n", bullet)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

