package main

import (
	"flag"
	"fmt"
	"scanfactory"
	. "types"
	"os"
)

var ipRange *string = flag.String("ip", "127.0.0.1", `IP range of scan. Example:
		scah your-ip`)
var demo *string = flag.String("d", "scan ip=127.0.0.1 port=1-1024", `go run nscan.go -ip=your-ip`)

var help *string = flag.String("h", "help doc", "help doc")

func PortRange() (ports []Port) {
	s:=1024
	for i:=1;i!=s+1;i++{
		ports = append(ports, Port(i))
	}
	return
}

func getIp(ipRange *string) (ips []Ip, err error) {
	ips = append(ips, Ip(*ipRange))
	return
}

func main() {
	flag.Parse()
	if len(os.Args) > 1 && os.Args[1] == "-h" ||
		len(os.Args) == 1 {
		flag.Usage()
		os.Exit(-1)
	}
	ips, err := getIp(ipRange)
	if err != nil {
		os.Exit(-1)
	}
	ports:= PortRange()
	fmt.Printf("current scan IP :%s:",ips[0])
	fmt.Println("\n")
	scan := scanfactory.NewScan(ips, ports)
	scan.Scan()
}
