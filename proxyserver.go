package goproxyserver

import (
  "net"
)

type GoProxyServer struct {
  srcAddress string
  dstAddress string
  srcPort string
  dstPort string
  srcSocket *net.TCPAddr
  dstSocket *net.TCPAddr
}
