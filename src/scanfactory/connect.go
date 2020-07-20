package scanfactory

import (
	. "types"
	"net"
	"strconv"
	"time"
)

func connect(addr Addr, c chan Result, ok chan bool) {
	remote := string(addr.Ip) + ":" + strconv.Itoa(int(addr.Port))
	conn, err := net.DialTimeout("tcp", remote, 5*time.Second)
	var result Result
	if err == nil {
		conn.Close()
		result = Result{addr, true}
	} else {
		result = Result{addr, false}
	}
	c <- result
	ok <- true
}
