package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	Option_Local    string
	Option_Remote   string
	Option_Protocol string
	Option_Tls      bool
	Option_Help     bool
)

func init() {
	flag.BoolVar(&Option_Help, "help", false, "this help")
	flag.StringVar(&Option_Local, "local", "0.0.0.0:8080", "set local tcp protocol listen address.")
	flag.StringVar(&Option_Remote, "remote", "", "set remote tcp protocol proxy address.")
	flag.StringVar(&Option_Protocol, "protocol", "tcp", "set remote tcp protocol tcp or tcp6.")
	flag.BoolVar(&Option_Tls, "tls", false, "enable remote tcp with tls encryption.")
}

func IoCopy(c *sync.WaitGroup, up bool, in net.Conn, out net.Conn) {
	defer c.Done()
	size, _ := io.Copy(out, in)
	if size > 0 {
		if up {
			StatUpdate(size, 0)
		} else {
			StatUpdate(0, size)
		}
	}
	in.Close()
	out.Close()
}

func TcpProxy(local_conn net.Conn, remote_conn net.Conn) {
	log.Printf("start %s connect to %s\n", local_conn.RemoteAddr(), remote_conn.RemoteAddr())

	syncSem := new(sync.WaitGroup)
	syncSem.Add(2)

	go IoCopy(syncSem, true, local_conn, remote_conn)
	go IoCopy(syncSem, false, remote_conn, local_conn)

	syncSem.Wait()

	log.Printf("close %s connect to %s\n", local_conn.RemoteAddr(), remote_conn.RemoteAddr())
}

func main() {
	flag.Parse()
	if Option_Help {
		flag.Usage()
		return
	}

	listen, err := net.Listen("tcp", Option_Local)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	log.Printf("listen %s successfully\n", Option_Local)

	var tls_config *tls.Config

	if Option_Tls {
		tls_config, err = TlsConfigClient(Option_Remote)
		if err != nil {
			log.Println(err.Error())
		}
	}

	for {
		local_conn, err := listen.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		remote_conn, err := net.Dial(Option_Protocol, Option_Remote)
		if err != nil {
			log.Println(err.Error())
			local_conn.Close()
			continue
		}

		if tls_config != nil {
			remote_conn = tls.Client(remote_conn, tls_config)
		}

		go TcpProxy(local_conn, remote_conn)
	}
}
