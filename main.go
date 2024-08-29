package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

const defaultListenAddress =":5001"
type Config struct{
	ListenAddress string
}

type Server struct{
	Config
	peers map[*Peer] bool
	ln net.Listener
	addPchan chan *Peer
	quitchan chan struct{}
	msgCh chan []byte
}

func NewServer (C Config) *Server{
	if len(C.ListenAddress)==0{
		C.ListenAddress=defaultListenAddress
	}
	return &Server{
		Config: C,
		peers:make(map[*Peer]bool),
		addPchan: make(chan *Peer), 
		quitchan: make(chan struct{}),
		msgCh: make(chan []byte),
	}
} 

func (S *Server) Start() error{
	ln,err := net.Listen("tcp",S.ListenAddress)
	if err !=nil{
		return err
	}
	S.ln=ln
	go S.loop()

	slog.Info("server running","listenAddr",S.ListenAddress)
	return S.acceptLoop()
}

func (S *Server) acceptLoop() error{
	for{
		conn,err:=S.ln.Accept()
		if err!=nil{
			slog.Error("error accepted","err",err)
			continue
		}
		go S.handleConn(conn)
	}
}

func (S *Server) handleMessage(rawMsg []byte) error{
	fmt.Println(string(rawMsg))
	return nil
}

func (S *Server) loop(){
	for{
		select{
		case rawMsg:=<-S.msgCh:
			if err :=S.handleMessage(rawMsg); err!=nil{
				slog.Error("error in RawMessage","err",err)
			}
			fmt.Println(rawMsg)
		case <-S.quitchan:
			return 
		case peer:=<-S.addPchan:
			S.peers[peer] =true
	
		}
	}
}

func (S *Server) handleConn(conn net.Conn){
	peer:=NewPeer(conn,S.msgCh)
	S.addPchan <-peer
	slog.Info("new peer connected","remoteAddress",conn.RemoteAddr())
	if err:=peer.readLoop();err!=nil{
		slog.Error("Error in peer","err",err,"remoteAddr",conn.RemoteAddr())
	}
}

func main()  {
	s:=NewServer(Config{})
	log.Fatal(s.Start())
	fmt.Println("HEllo go")
}
 