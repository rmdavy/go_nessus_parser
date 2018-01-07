package main

import (
	"errors"
	"fmt"
	"log"
	"nessusToolCLI/nessusTools"
	"os"
)

var Hosts []nessusTools.Host

func main() {

	var arg argList

	if len(os.Args) < 2 {
		arg.helpText()
		os.Exit(0)
	}

	arg = arg.parse()

	path := ""
	err := errors.New("")

	if arg.help == true {
		arg.helpText()
		os.Exit(0)
	}

	if len(arg.reports) > 1 {
		path, err = nessusTools.CombineReports(arg.reports)
		if err != nil {
			log.Fatal(err)
		}
	} else if len(arg.reports) <= 0 {
		fmt.Println("Error: No arguments supplied for --reports")
		arg.helpText()
		os.Exit(0)
	} else {
		path = arg.reports[0]
	}

	if arg.list == true {
		nessusTools.ListIssues(path)
	}

	if arg.issues != nil {
		path, err = nessusTools.RemoveIssues(path, arg.issues)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.issuesContains != nil {
		path, err = nessusTools.RemoveIssuesContains(path, arg.issuesContains)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.hosts != nil {
		path, err = nessusTools.RemoveHosts(path, arg.hosts)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.ciphers == true {
		Hosts, err = nessusTools.HostListBuilder(path, arg.file)
		nessusTools.Ciphers(path, arg.file, Hosts)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.searchString != "" {
		Hosts, err = nessusTools.HostListBuilder(path, arg.file)
		nessusTools.SearchIssues(path, arg.searchString, arg.file, Hosts)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.sslIssues == true {
		Hosts, err = nessusTools.HostListBuilder(path, arg.file)
		nessusTools.SslIssues(path, arg.file, Hosts)
		if err != nil {
			fmt.Println(err)
		}
	}

	if arg.summary == true {
		Hosts, err = nessusTools.HostListBuilder(path, arg.file)
		nessusTools.Summary(path, arg.file, Hosts)
		if err != nil {
			fmt.Println(err)
		}
	}
}
