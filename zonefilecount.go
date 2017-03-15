package main

import (
	"bufio"
	"fmt"
	"github.com/miekg/dns"
	"os"
	"strings"
)

type Domains struct {
	count    int
	idn      int
	previous string
	suffix   int
}

var rrParsed int

/* Return rdata */
func rdata(RR dns.RR) string {
	return strings.Replace(RR.String(), RR.Header().String(), "", -1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Help:\n")
		fmt.Println("     zonefilecount <zonefile>\n")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	domains := new(Domains)

	ns := map[string]int{}
	signed := map[string]int{}

	zoneFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("ERROR: Can't open zone file.")
	}

	zone := dns.ParseZone(bufio.NewReader(zoneFile), "", "")

	var rrtypes [100]int

	for parsedLine := range zone {
		if parsedLine.RR != nil {
			rrtypes[parsedLine.RR.Header().Rrtype]++

			switch parsedLine.RR.Header().Rrtype {
			case dns.TypeDS:
				/* Increment Signed Domains counter */
				signed[parsedLine.RR.Header().Name]++
			case dns.TypeNS:
				/* Increment NS counter */
				ns[rdata(parsedLine)]++

				if parsedLine.RR.Header().Name != domains.previous { // Unique domain

					/* Increment Domain counter */
					domains.count++
					domains.previous = parsedLine.RR.Header().Name

					/* Check if the domain is an IDN */

					if strings.HasPrefix(strings.ToLower(parsedLine.RR.Header().Name), "xn--") {
						domains.idn++
					}

					/* Display progression */
					if domains.count%1000000 == 0 {
						fmt.Printf("*")
					} else if domains.count%100000 == 0 {
						fmt.Printf(".")
					}
				}
			}
		} else {
			fmt.Println("ERROR: A problem occured while parsing the zone file.")
		}

		/* Increment number of resource records parsed */
		rrParsed++
	}

	/* Don't count origin */
	domains.count--

	fmt.Println("A;AAAA;CNAME;NS;MX")
	fmt.Println(rrtypes[dns.TypeA], ";", rrtypes[dns.TypeAAAA], ";", rrtypes[dns.TypeCNAME],";", rrtypes[dns.TypeNS], ";", ";", rrtypes[dns.TypeMX])
}
