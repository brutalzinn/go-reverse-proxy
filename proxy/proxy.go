package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/brutalzinn/go-reverse-proxy/config"
	"github.com/brutalzinn/go-reverse-proxy/utils"
	"github.com/sirupsen/logrus"
)

func StartProxy(route config.Routes) {
	inPath := route.IN
	outPath := route.OUT

	switch inPath.Protocol {
	case "http":
		startHTTPProxy(inPath, outPath)
	case "tcp":
		startTCPProxy(inPath, outPath)
	case "udp":
		startUDPProxy(inPath, outPath)
	default:
		logrus.Fatalf("Unsupported protocol: %s", inPath.Protocol)
	}
}

func startHTTPProxy(inPath, outPath config.ProxyPath) {
	http.HandleFunc(inPath.Path, func(w http.ResponseWriter, r *http.Request) {
		data := fmt.Sprintf("Received HTTP request to %s", r.RequestURI)
		logrus.Println(data)
		if outPath.Protocol == "tcp" {
			forwardToTCP(outPath, []byte(data))
		} else if outPath.Protocol == "udp" {
			forwardToUDP(outPath, []byte(data))
		} else {
			w.Write([]byte("Unsupported output protocol"))
		}
	})

	addr := fmt.Sprintf("%s:%s", inPath.Host, inPath.Port)
	logrus.Printf("Starting HTTP proxy on %s", addr)
	logrus.Fatal(http.ListenAndServe(addr, nil))
}

func startTCPProxy(inPath, outPath config.ProxyPath) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", inPath.Host, inPath.Port))
	if err != nil {
		logrus.Fatalf("Error starting TCP listener: %v", err)
	}
	defer listener.Close()

	logrus.Printf("Starting TCP proxy on %s", inPath.Host+":"+inPath.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Println("Error accepting TCP connection:", err)
			continue
		}
		go handleTCPConnection(conn, outPath)
	}
}

func handleTCPConnection(conn net.Conn, outPath config.ProxyPath) {
	defer conn.Close()

	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil && err != io.EOF {
		logrus.Println("Error reading from TCP connection:", err)
		return
	}

	if outPath.Protocol == "tcp" {
		forwardToTCP(outPath, data[:n])
	} else if outPath.Protocol == "udp" {
		forwardToUDP(outPath, data[:n])
	}
}

func startUDPProxy(inPath, outPath config.ProxyPath) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(inPath.Host),
		Port: utils.ParseINT(inPath.Port),
	})
	if err != nil {
		logrus.Fatalf("Error starting UDP listener: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			logrus.Println("Error reading from UDP connection:", err)
			continue
		}

		if outPath.Protocol == "tcp" {
			forwardToTCP(outPath, buf[:n])
		} else if outPath.Protocol == "udp" {
			forwardToUDP(outPath, buf[:n])
		}
	}
}

func forwardToTCP(outPath config.ProxyPath, data []byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", outPath.Host, outPath.Port))
	if err != nil {
		logrus.Println("Error connecting to TCP target:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		logrus.Println("Error writing to TCP target:", err)
	}
}

func forwardToUDP(outPath config.ProxyPath, data []byte) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", outPath.Host, outPath.Port))
	if err != nil {
		logrus.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		logrus.Println("Error connecting to UDP target:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		logrus.Println("Error writing to UDP target:", err)
	}
}
