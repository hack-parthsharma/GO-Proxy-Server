package goproxyserver

import (
  "log"
)

func termError(msg string, err error){
  if err != nil {
    log.Fatal(msg, ":", err.Error())
  }
}

func passError(msg string, err error){
  if err != nil{
    log.Println(msg,":",err.Error())
  }
}
