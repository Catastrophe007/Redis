package main

import (
	// "fmt"
	"net"
)

type Peer struct {
	conn net.Conn
	msgCh chan []byte
}

func NewPeer (c net.Conn,msgCh chan[] byte) *Peer{
	return &Peer{
		conn:c,
		msgCh:msgCh,
	}
}

func (p *Peer) readLoop() error{
	buf:=make([]byte,1024)
	for{
		n,err:=p.conn.Read(buf)
		if err!=nil{
			return err
		}
		msgBuf :=make([]byte,n)
		copy(msgBuf,buf[:n])  
		p.msgCh<-msgBuf
		// fmt.Print(string(msgBuf))
	}
}