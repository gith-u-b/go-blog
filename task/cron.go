package main

import (
	"time"
	"log"

	"github.com/robfig/cron"

	"go-blog-step-by-step/models"
)


func main() {
	log.Println("Starting...")

	c := cron.New()
	err := c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
