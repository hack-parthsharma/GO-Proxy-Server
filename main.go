package main

import (
		"github.com/cyberkhalid/gosec/goproxyserver/goproxyserver"
		"flag"
)

func main()  {
	targetIp := flag.String("targetip","","Ip address of the target host")
	targetPort:= flag.String("targetport","","Port number of the target host")
	proxyIp := flag.String("proxyip","","Ip address of proxy")
	proxyPort := flag.String("proxyport","","Port number to be used by proxy")
	flag.Parse()

	var goProxyServer goproxyserver.GoProxyServer
	goProxyServer.Init(*proxyIp,*proxyPort,*targetIp,*targetPort)
	goProxyServer.Connect()
}
