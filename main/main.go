package main

import (
	"fmt"
	"landlords/logic"
	"log"
	"net/http"
)

func initLogic() (err error) {
	var roomConf []*logic.RoomConf
	for i := 0; i < 5; i++ {
		conf := &logic.RoomConf{
			Id:           fmt.Sprintf("%d", i),
			Name:         fmt.Sprintf("room_%d", i),
			DeskNum:      16,
			MaxPlayerNum: 300,
		}

		roomConf = append(roomConf, conf)
	}

	err = logic.Init(roomConf, 1000)
	return
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
	return
}

func ws(w http.ResponseWriter, r *http.Request) {

	err := logic.HandleConn(w, r)
	if err != nil {
		return
	}
	return
}

func main() {

	err := initLogic()
	if err != nil {
		fmt.Printf("init logic failed, err:%v\n", err)
		return
	}

	fmt.Println("init logic succ")
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws(w, r)
	})
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
