package scanfactory

import (
	"fmt"
	. "github.com/cen9jing/nscan/src/types"
)

type Addrs struct {
	ip   []Ip
	portList []Port
}

type Scan struct {
	Addrs
	ch      chan Result
	ok      chan bool
	results []Result
	sumAddr int
}

func NewScan(ip []Ip, portList []Port) (scan *Scan) {
	sumAddr := len(ip) * len(portList)
	ch := make(chan Result, sumAddr)
	ok := make(chan bool, sumAddr)
	return &Scan{Addrs: Addrs{ip, portList}, ch: ch, ok: ok, sumAddr: sumAddr}
}

func (s *Scan) Scan() {
	for _, ip := range s.ip {
		for _, port := range s.portList {
			addr := Addr{ip, port}
			go connect(addr, s.ch, s.ok)
		}
	}
	s.waitResults()
}

func (s *Scan) waitResults() {
	var result Result
	for {
		select {
		case result = <-s.ch:
			s.results = append(s.results, result)
			if result.Open {
				fmt.Printf("port: %d open\n", result.Addr.Port)
			}
		case <-s.ok:
			s.sumAddr--
			if s.sumAddr == 0 {
				fmt.Println("scan over")
				goto LOOP
			}
		}
	}
LOOP:
}

func (s *Scan) GetResults() []Result {
	return s.results
}
