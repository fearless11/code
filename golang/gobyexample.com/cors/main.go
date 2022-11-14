package main

import (
	"fmt"
	"net/http"
	"os"
)

var corsHeaders = map[string]string{
	"Access-Control-Allow-Headers":  "Accept, Authorization, Content-Type, Origin,X-Token",
	"Access-Control-Allow-Methods":  "GET, DELETE, OPTIONS, PUT",
	"Access-Control-Allow-Origin":   "*",
	"Access-Control-Expose-Headers": "Date",
	"Cache-Control":                 "no-cache, no-store, must-revalidate",
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	// cors的头部设置
	for h, v := range corsHeaders {
		w.Header().Set(h, v)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func run() int {

	listen := ":8080"
	// 设置多路路由
	mux := http.NewServeMux()
	mux.HandleFunc("/cors", corsHandler)

	srv := http.Server{Addr: listen, Handler: mux}

	srvc := make(chan struct{})

	go func() {

		fmt.Printf("start listening %v\n", listen)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("listen error ", err)
			close(srvc)
		}

		defer func() {
			if err := srv.Close(); err != nil {
				fmt.Printf("Error on closing the server %v", err)
			}
		}()

	}()

	for {
		select {
		case <-srvc:
			return 1
		}
	}
}

func main() {
	os.Exit(run())
}
