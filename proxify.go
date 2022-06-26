package goproxyserver

import (
  "net"
  "fmt"
)

func proxify(dstCon *net.TCPConn, srcSocket, dstSocket *net.TCPAddr){
  srcCon, err := net.DialTCP("tcp", nil, dstSocket)
  if err != nil {
    fmt.Println("[-] Can't connect to the target", dstSocket,err.Error())
    return
  }
  defer srcCon.Close()

  fmt.Println("[+] Proxy Status: ", srcSocket,"<-->",dstSocket)
  //copy from src to dst
  go copyData(srcCon, dstCon)

  //copy from dst to src
  copyData(dstCon, srcCon)

}
