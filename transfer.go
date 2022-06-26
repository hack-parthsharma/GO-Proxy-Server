package goproxyserver

import (
  "io"
  "net"
)

func copyData(srcCon, dstCon *net.TCPConn)  {
   _, err := io.Copy(dstCon, srcCon)
   termError("[-] Error Transfer", err)
}
