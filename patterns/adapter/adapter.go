package main

import "fmt"

type Mobile interface {
	PrintName()
}

type Huawei struct {
}

func (h *Huawei) Name() {
	fmt.Println("Huawei 没有PrintName方式 只有Name方法")
	fmt.Println("China-made mobile phones:华为")
}

type HuaweiAdapter struct {
	Hw *Huawei
}

func (h *HuaweiAdapter) PrintName() {
	fmt.Println("封装一层适配器，让请求端自动翻译")
	h.Hw.Name()
}

type Iphone struct {
}

func (i *Iphone) PrintName() {
	fmt.Println("USA-made mobile phones:苹果")
}

type Client struct {
}

func (c *Client) Name(mobile Mobile) {
	mobile.PrintName()
}

func main() {
	hw := &Huawei{}
	ip := &Iphone{}
	hwAdapter := &HuaweiAdapter{Hw: hw}

	client := &Client{}

	client.Name(hwAdapter)
	client.Name(ip)

}
