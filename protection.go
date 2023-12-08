// layer7 ddos protection

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Started")
	for {
		time.Sleep(1 * time.Second)
		command := "netstat -an | grep 443 | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n"

		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			fmt.Printf("%s", err)
		}

		output := string(out[:])

		for _, line := range strings.Split(output, "\n") {
			if line != "" {
				fields := strings.Fields(line)
				if len(fields) == 2 {
					connections, _ := strconv.Atoi(fields[0])
					ip := fields[1]

					if connections > 10 {
						go blockip(ip, connections)
					}
				}
			}
		}
	}
}

func blockip(ip string, connections int) {
	fmt.Println("Blocking " + ip + " with " + strconv.Itoa(connections) + " connections")
	command := "iptables -A INPUT -s " + ip + " -j DROP"
	exec.Command("bash", "-c", command).Output()
}

func unblockipeveryhour() {
	for {
		time.Sleep(1 * time.Hour)
		command := "iptables -F"
		exec.Command("bash", "-c", command).Output()
	}
}
