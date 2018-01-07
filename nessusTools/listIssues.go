package nessusTools

import (
	"fmt"

	"github.com/beevik/etree"
)

func ListIssues(filePath string) ([]string, error) {
	//current report hosts are being added from
	report := etree.NewDocument()
	if err := report.ReadFromFile(filePath); err != nil {
		return nil, err
	}
	issues := []string{}

	root := report.SelectElement("NessusClientData_v2")
	for _, report := range root.SelectElements("Report") { //Select report branch to iterate over
		for _, host := range report.SelectElements("ReportHost") { //Select Reporthost branch to iterate over
			for _, item := range host.SelectElements("ReportItem") { //Select ReportItem branch to iterate over
				var nameID string
				nameID = item.SelectAttrValue("pluginName", "Not Found")
				spaces := 70 - len(nameID)
				for i := 0; i < spaces; i++ {
					nameID += " "
				}
				nameID += "| ID: " + item.SelectAttrValue("pluginID", "Not Found")
				issues = append(issues, nameID) //Select the value of the plugin name attribute(NOT FOUND is the default value)
			}
		}
	}
	issues = RemoveDuplicates(issues)
	for a, b := range issues {
		fmt.Printf("%d | Issue: %s\n", a, b)
	}
	return issues, nil
}
