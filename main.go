package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"webServer/tools"
)

func init() {
	// errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("打开日志文件失败：", err)
	// }

	// Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
}

func paserParam(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("id")
	w.Write([]byte(v))
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	loghelper.Info.Println("request:" + r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello world"))
}

func handleJSONRes(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("id")
	if v == "EFL" {
		w.Write(load())
	}
}

func load() []byte {
	data, err := ioutil.ReadFile("EFL.json")
	if err != nil {
		return nil
	}
	return data
	// fmt.Println("$s", data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHelloName)
	mux.HandleFunc("/p", paserParam)
	mux.HandleFunc("/json", handleJSONRes)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
