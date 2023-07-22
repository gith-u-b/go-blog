package main

import (
	"context"
	"fmt"
	"go-blog-step-by-step/models"
	"go-blog-step-by-step/pkg/gredis"
	"go-blog-step-by-step/pkg/logging"
	"go-blog-step-by-step/pkg/setting"
	"go-blog-step-by-step/pkg/upload"
	"go-blog-step-by-step/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init(){
	log.SetFlags(log.Ldate|log.Lshortfile)
}

func main() {

	setting.Setup()
	models.Setup()
	logging.Setup()
	upload.Setup()
	gredis.Setup()


	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	//if err := s.ListenAndServe(); err != nil {
	//	log.Printf("Listen: %s\n", err)
	//}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	log.Println("serving...")
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	models.CloseDB()

	log.Println("Server exiting")

}

