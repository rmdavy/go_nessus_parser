package nessusTools

import (
	"fmt"
	"os"
	"strings"
)

//SearchIssues collects all issues and output from a nessus report, prunes repreating information and presents output to a user based on a specific search term.
func SearchIssues(filePath string, search string, writeToFile bool, Hosts []Host) ([]Host, error) {
	for a, b := range Hosts {
		Hosts[a] = b.RemoveDuplicateIssues()
	}

	f, err := os.Create("output") //write to a file called output

	for _, a := range Hosts {
		hostPrinted := 0 //Change to 1 when a hostname is printed. (resets on each new host)
		for _, b := range a.issues {
			if strings.Contains(b.issueName, search) {
				if hostPrinted == 0 { //check if host has been printed, if not, print
					if writeToFile == true {
						f.Write([]byte(a.name + "\n"))
					} else {
						fmt.Println(a.name)
					}
					hostPrinted = 1 //change to 1 now hostname has been printed.
				}
				if writeToFile == true {
					f.Write([]byte("	" + b.issueName + "\n"))
				} else {
					fmt.Println("	" + b.issueName)
				}
				list := []string{}   //list for containing and manipulating nessus pluginOutput data
				if search == "SSL" { //copy this block to add in new search specific criteria
					for _, c := range b.output {
						list = append(list, sslSearch(c)...)
					}
					list = RemoveDuplicates(list)
					for _, a := range list {
						if writeToFile == true {
							f.Write([]byte("		" + a + "\n"))
						} else {
							fmt.Println("		" + a)
						}
					}
				} else {
					for _, c := range b.output {
						if writeToFile == true {
							f.Write([]byte(c + "        \n"))
						} else {
							fmt.Println(c + "        \n")
						}
					}
				}
			}
		}
	}
	return Hosts, err
}
