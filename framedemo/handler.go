package main

import (
	"context"
	"fmt"
	"framedemo/framework"
	"log"
	"net/http"
	"time"
)

func TestHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.FrameContext(), 20*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(http.StatusOK, "ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.Lock()
		defer c.Unlock()
		log.Println(p)
		c.Json(http.StatusInternalServerError, "Panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.Lock()
		defer c.Unlock()
		c.Json(http.StatusGatewayTimeout, "Time Out")
		c.SetTimeout()
	}
	return nil

}
