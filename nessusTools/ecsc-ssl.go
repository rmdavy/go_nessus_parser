package nessusTools

import (
	"errors"
	"fmt"
	"os"
)

//Hosts is a list of host objects
//var Hosts []Host

//Ciphers combines a bunch of issues into output in either the terminal or a file
func Ciphers(filePath string, writeToFile bool, Hosts []Host) ([]Host, error) {
	issueIDs := []string{"20007", "65821", "78479", "71549", "26928", "94437", "104743", "81606", "83875", "58751", "51192", "95631", "69551", "15901"}
	issuesIDMap := make(map[string]bool)
	for _, b := range issueIDs {
		issuesIDMap[b] = true
	}
	search := "SSL"

	for a, b := range Hosts {
		Hosts[a] = b.RemoveDuplicateIssues()
	}

	f, err := os.Create("output") //write to a file called output

	for _, a := range Hosts {
		hostPrinted := 0 //Change to 1 when a hostname is printed. (resets on each new host)
		for _, b := range a.issues {
			if issuesIDMap[b.issueID] == true {
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

//SslIssues lists findings by issue and only the findings with ID's in the issueID's slice.
func SslIssues(filePath string, writeToFile bool, Hosts []Host) error {
	issueIDs := []string{"20007", "65821", "78479", "71549", "26928", "94437", "104743", "81606", "83875", "58751", "51192", "95631", "69551", "15901"}
	issuesIDMap := make(map[string]bool)
	var issues []issue
	err := errors.New("")

	for _, b := range issueIDs {
		issuesIDMap[b] = true
	}

	if len(Hosts) == 0 {
		err = errors.New("error: No hosts have been parsed")
	}

	for _, a := range Hosts {
		for _, b := range a.issues {
			if issuesIDMap[b.issueID] == true {
				var thisIssue issue
				thisIssue.issueID = b.issueID
				thisIssue.issueName = b.issueName
				thisIssue.port = b.port
				thisIssue.output = append(thisIssue.output, a.name+":"+thisIssue.port)
				if len(issues) < 1 {
					issues = append(issues, thisIssue)
				} else {
					issueFound := false
					for c := 0; c < len(issues); c++ {
						if thisIssue.issueID == issues[c].issueID {
							issues[c].output = append(issues[c].output, thisIssue.output...)
							issueFound = true
						}
					}
					if issueFound == false {
						issues = append(issues, thisIssue)
					}
				}
			}
		}
	}
	for _, a := range issues {
		fmt.Println(a.issueName)
		for _, b := range a.output {
			fmt.Println("    ", b)
		}
	}

	return err
}

//Summary prints a list of issues from the supplied reports with associated ID's and open ports
func Summary(filePath string, writeToFile bool, Hosts []Host) error {
	issueIDs := []string{"20007", "65821", "78479", "71549", "26928", "94437", "104743", "81606", "83875", "58751", "51192", "95631", "69551", "15901"}
	issuesIDMap := make(map[string]bool)
	var issues []issue
	err := errors.New("")

	for _, b := range issueIDs {
		issuesIDMap[b] = true
	}

	if len(Hosts) == 0 {
		err = errors.New("error: No hosts have been parsed")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\n----------------------------------------------\nOpen ports (Nessus scan only):\n\n")
	openPorts := make(map[string]string)
	for _, a := range Hosts {
		for _, b := range a.issues {
			if b.issueName == "Nessus SYN scanner" {
				if _, ok := openPorts[a.name]; ok {
					openPorts[a.name] += ", " + b.port
				} else {
					openPorts[a.name] = b.port
				}
			}
		}
	}

	for a, b := range openPorts {
		fmt.Printf("%s: %s\n", a, b)
	}

	fmt.Printf("\n\n----------------------------------------------\nHosts Affected by SSL/TLS vulnerabilities\n\n")
	vulnerableHosts := []string{}
	for _, a := range Hosts {
		vulnerable := false
		for _, b := range a.issues {
			if issuesIDMap[b.issueID] == true {
				vulnerable = true
			}
		}
		if vulnerable == true {
			vulnerableHosts = append(vulnerableHosts, a.name)
		}
	}

	for a, b := range vulnerableHosts {
		if a == len(vulnerableHosts)-1 {
			fmt.Printf("%s\n", b)
		} else {
			fmt.Printf("%s, ", b)
		}
	}

	fmt.Printf("\n\n----------------------------------------------\nSSL issues listed by host:port\n\n")
	for _, a := range Hosts {
		for _, b := range a.issues {
			if issuesIDMap[b.issueID] == true {
				var thisIssue issue
				thisIssue.issueID = b.issueID
				thisIssue.issueName = b.issueName
				thisIssue.port = b.port
				thisIssue.output = append(thisIssue.output, a.name+":"+thisIssue.port)
				if len(issues) < 1 {
					issues = append(issues, thisIssue)
				} else {
					issueFound := false
					for c := 0; c < len(issues); c++ {
						if thisIssue.issueID == issues[c].issueID {
							issues[c].output = append(issues[c].output, thisIssue.output...)
							issueFound = true
						}
					}
					if issueFound == false {
						issues = append(issues, thisIssue)
					}
				}
			}
		}
	}
	for _, a := range issues {
		fmt.Println(a.issueName)
		for _, b := range a.output {
			fmt.Println("    ", b)
		}
	}

	return err
}
