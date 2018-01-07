package nessusTools

import (
	"errors"
	"strings"

	"github.com/beevik/etree"
)

//SearchIssues collects all issues and output from a nessus report, prunes repreating information and presents output to a user based on a specific search term.
func HostListBuilder(filePath string, writeToFile bool) ([]Host, error) {
	//current report hosts are being added from
	var Hosts []Host
	err := errors.New("")
	counter := 0

	report := etree.NewDocument()
	if err := report.ReadFromFile(filePath); err != nil {
		return Hosts, err
	}

	root := report.SelectElement("NessusClientData_v2")
	for _, report := range root.SelectElements("Report") { //Select report branch to iterate over
		for _, hostName := range report.SelectElements("ReportHost") { //Select Reporthost branch to iterate over
			var thisHost Host
			thisHost.name = hostName.SelectAttrValue("name", "")
			for _, item := range hostName.SelectElements("ReportItem") { //Select ReportItem branch to iterate over
				var thisIssue issue
				id := item.SelectAttrValue("pluginID", "Not Found")
				port := item.SelectAttrValue("port", "Not Found")
				thisIssue.issueID = id
				thisIssue.port = port
				thisIssue.issueName = string(item.SelectAttrValue("pluginName", ""))
				output := item.SelectElement("plugin_output")
				counter = 0
				if counter == 0 {
					thisIssue.output = append(thisIssue.output, "") //Initialise output slice
				}
				if output != nil {
					for _, char := range output.Text() {
						thisIssue.output[counter] += string(char)
						if char == '\n' {
							thisIssue.output[counter] = strings.TrimSpace(thisIssue.output[counter]) //Append the output from a duplicate issue to the first instance of the issue
							thisIssue.output = append(thisIssue.output, "")
							counter++
						}
					}
				}
				thisHost.issues = append(thisHost.issues, thisIssue)
			}
			Hosts = append(Hosts, thisHost)
		}
	}
	return Hosts, err
}
