package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From Here!"))
}

func listenAndServe(port string){	
	err := http.ListenAndServe(port, nil)
	if err!= nil{
		fmt.Printf("%+v", err)
	}

}


func HandlerFunc(w http.ResponseWriter, r* http.Request){
	c, err := websocket.Accept(w, r, nil)

	if err != nil{
		fmt.Printf("%+v", err)
	}
	defer c.CloseNow()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("received: %v", v)

	c.Close(websocket.StatusNormalClosure, "")

}

func main(){
	http.HandleFunc("/test", helloWorld)
	http.HandleFunc("/socket", HandlerFunc)
	fs := http.FileServer(http.Dir("./data"))
	http.Handle("/",  fs)
	
	listenAndServe(":3000")	
	
}

