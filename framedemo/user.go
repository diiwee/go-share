package main

import (
	"fmt"
	"framedemo/framework"
	"net/http"
	"time"
)

func Login(c *framework.Context) error {
	time.Sleep(2 * time.Second)
	c.Json(http.StatusOK, fmt.Sprintf("URl:%s  Meth:%s", c.GetRequest().URL, c.GetRequest().Method))
	return nil
}

func Create(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("URl:%s  Meth:%s", c.GetRequest().URL, c.GetRequest().Method))
	return nil
}

func Update(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("URl:%s  Meth:%s", c.GetRequest().URL, c.GetRequest().Method))
	return nil
}

func Delete(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("URl:%s  Meth:%s", c.GetRequest().URL, c.GetRequest().Method))
	return nil
}
