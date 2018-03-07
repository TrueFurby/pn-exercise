package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	portscanner "github.com/anvie/port-scanner"
)

func main() {
	flag.Parse()

	first := flag.Arg(0)
	var last string
	if flag.NArg() > 1 {
		last = flag.Arg(1)
	}

	for _, ip := range rangeToIPs(first, last) {
		file := filepath.Join("/tmp", fmt.Sprintf("scandata-%s", ip))
		prev := loadData(file)
		if data := scanTarget(ip); data == prev {
			if prev != "" {
				fmt.Println("no new data found in last scan")
			}
			continue
		} else {
			storeData(file, data)
			fmt.Println(data)
		}
	}
}

func loadData(file string) string {
	if _, err := os.Stat(file); err != nil {
		return ""
	}
	data, _ := ioutil.ReadFile(file)
	return string(data)
}

func storeData(file, data string) {
	fileHandle, _ := os.Create(file)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprint(writer, data)
	writer.Flush()
}

func scanTarget(ip string) string {
	fmt.Printf("=> Target: %v\n", ip)

	ps := portscanner.NewPortScanner(ip, 1*time.Second, 10)

	openedPorts := ps.GetOpenedPort(20, 100)

	var out string
	for _, port := range openedPorts {
		out += fmt.Sprint(" ", port, " [open]")
		out += fmt.Sprintln("  -->  ", ps.DescribePort(port))
	}
	return out
}

func rangeToIPs(first, last string) (ips []string) {
	if last == "" {
		return []string{first}
	}

	f := ip2int(net.ParseIP(first))
	l := ip2int(net.ParseIP(last))
	for i := f; i <= l; i++ {
		ips = append(ips, int2ip(i).String())
	}

	return
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
