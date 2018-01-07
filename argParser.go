package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	addReport        = false
	addIssue         = false
	addIssueContains = false
	addHost          = false
	search           = false
)

type argList struct {
	reports        []string
	issues         []string
	issuesContains []string
	hosts          []string
	searchString   string
	file           bool
	list           bool
	help           bool
	ciphers        bool
	sslIssues      bool
	summary        bool
}

func (a argList) parse() argList {
	var thisArgList argList
	for _, b := range os.Args {
		if strings.HasPrefix(b, "--") {
			addReport = false
			addIssue = false
			addIssueContains = false
			addHost = false
			search = false
			if b == "--r" || b == "--reports" {
				addReport = true
			} else if b == "--i" || b == "--issues" {
				addIssue = true
			} else if b == "--ic" || b == "--issuesContains" {
				addIssueContains = true
			} else if b == "--h" || b == "--hosts" {
				addHost = true
			} else if b == "--f" || b == "--file" {
				thisArgList.file = true
			} else if b == "--s" || b == "--search" {
				search = true
			} else if b == "--l" || b == "--list" {
				thisArgList.list = true
			} else if b == "--help" {
				thisArgList.help = true
			} else if b == "--ciphers" {
				thisArgList.ciphers = true
			} else if b == "--sslIssues" {
				thisArgList.sslIssues = true
			} else if b == "--summary" {
				thisArgList.summary = true
			} else {
				fmt.Println("Invalid argument:", b)
				thisArgList.help = true
			}
		} else if addReport == true {
			thisArgList.reports = append(thisArgList.reports, b)
		} else if addIssue == true {
			thisArgList.issues = append(thisArgList.issues, b)
		} else if addIssue == true {
			thisArgList.issuesContains = append(thisArgList.issuesContains, b)
		} else if addHost == true {
			thisArgList.hosts = append(thisArgList.hosts, b)
		} else if search == true {
			thisArgList.searchString = b
		}
	}
	return thisArgList
}

func (a argList) helpText() {
	fmt.Println("This program is used to combine and edit Nessus files. Only the .nessus extension is supported. When combining reports, removing issues or removing hosts a new nessus file will be created and saved in the current directory.")
	fmt.Println("\n\nOptions:")
	fmt.Println("	--r, --reports")
	fmt.Println("		location of nessus reports to be worked on.")
	fmt.Println("	--i, --issues")
	fmt.Println("		issues to be removed (must be pluginItem value as found in nessus file)")
	fmt.Println("	--ic, --issuesContains")
	fmt.Println("		issues to be removed (removes all issues continaing supplied strings)")
	fmt.Println("	--h, --host")
	fmt.Println("		Hosts to be removed.")
	fmt.Println("	--s, --search")
	fmt.Println("		display all output where pluginItem contains the supplied string (i.e. --search SSL will return all ssl issues.")
	fmt.Println("	--f, --file")
	fmt.Println("		save issue search/combo text to a file, saves to a file named output in current directory.")
	fmt.Println("	--l, --list")
	fmt.Println("		List type of issue in the supplied reports.")
	fmt.Println("	--ciphers")
	fmt.Println("		Print a selection of common ssl issues along with their associated ciphers.")
	fmt.Println("	--sslIssues")
	fmt.Println("		Print a selection of common ssl issues and the hosts affected.")
	fmt.Println("	--summary")
	fmt.Println("		Print a summary of useful information from the supplied reports.")
	fmt.Println("	--help")
	fmt.Println("		Display this help page.\n\n\n")
	fmt.Println("Example usage:")
	fmt.Println("	./nessusToolCLI --reports ~/Desktop/*nessus (combine all listed reports into a single nessus file)")
	fmt.Println("	./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f (List all output from SSL issues and save to file")
	fmt.Println("	./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f --issues SSL\\ Session\\ Resume\\ Supported --hosts 127.0.0.1 (as above but also removes ssl session resume issues and host 127.0.0.1)")
}
