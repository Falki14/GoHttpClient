package main

import (
	"github.com/Falki14/GoHttpClient/httpclient"
)

func main() {

	//client := httpclient.New("dfdsfds")
	baum := httpclient.New(5)
	baum.GetToken()
	baum.SetToken("DONG")
	print(baum.GetToken())
	//resp := client.Get("https://onfalkenberg.de")
	
	//fmt.Print(resp)
}