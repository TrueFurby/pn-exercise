package main

import (
	"flag"
	"fmt"
	"time"

	portscanner "github.com/anvie/port-scanner"
)

func main() {
	flag.Parse()
	target := flag.Arg(0)

	fmt.Printf("* Target - %v *\n", target)

	ps := portscanner.NewPortScanner(target, 2*time.Second, 10)

	openedPorts := ps.GetOpenedPort(20, 100)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}

}
