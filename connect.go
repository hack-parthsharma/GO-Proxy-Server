package goproxyserver

import (
  "net"
  "fmt"
)

func (p *GoProxyServer)Connect(){
  //var err error
  tcplisn, err := net.ListenTCP("tcp", p.srcSocket)
  termError("[-] Can't Listen on this socket :"+p.srcAddress+":"+p.srcPort, err)
  defer tcplisn.Close()
  fmt.Println("[+] Proxy Status: Listening At",p.srcSocket)
  for {
    con, err := tcplisn.AcceptTCP()
    passError("[-] Can't accept connection", err)
    defer con.Close()

    fmt.Println("[+] Proxy Status: Connected to", con.RemoteAddr().String())
    go proxify(con, p.srcSocket, p.dstSocket)
  }
}
