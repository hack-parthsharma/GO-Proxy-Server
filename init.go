package goproxyserver

import (
  "net"
  "fmt"
  "log"
)

func (p *GoProxyServer)Init(srcAddress, srcPort, dstAddress, dstPort string)  {
  if (srcAddress == "") || (srcPort == "") || (dstAddress == "") || (dstPort == ""){
    log.Fatal("[-] Provide correct arguments")
  }
  var err error
  p.srcAddress = srcAddress
  p.dstAddress = dstAddress
  p.srcPort = srcPort
  p.dstPort = dstPort

  srcSocket := fmt.Sprintf("%s:%s", srcAddress, srcPort)
  dstSocket := fmt.Sprintf("%s:%s", dstAddress, dstPort)

  p.srcSocket, err = net.ResolveTCPAddr("tcp", srcSocket)
  termError("[-] Can't resolve src address", err)

  p.dstSocket, err = net.ResolveTCPAddr("tcp", dstSocket)
  termError("[-] Can't resolve dst address", err)
}
