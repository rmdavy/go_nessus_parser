package nessusTools

import (
	"fmt"
	"os"
	"strings"

	"github.com/beevik/etree"
)

//RemoveIssuesContains removes all issues from a nessus file where the plugin name contians the supplied strings
func RemoveIssuesContains(filePath string, issues []string) (string, error) {
	fmt.Println("Removing Issues:")
	for _, b := range issues {
		fmt.Println("	", b)
	}

	//Open document and parse for issues
	report := etree.NewDocument()
	if err := report.ReadFromFile(filePath); err != nil {
		return "", err
	}

	root := report.SelectElement("NessusClientData_v2")
	for _, report := range root.SelectElements("Report") { //Select report branch to iterate over
		for _, host := range report.SelectElements("ReportHost") { //Select Reporthost branch to iterate over
			for _, item := range host.SelectElements("ReportItem") { //Select ReportItem branch to iterate over
				for _, b := range issues {
					if strings.Contains(item.SelectAttrValue("pluginName", "Not Found"), b) {
						host.RemoveChild(item)
					}
				}
			}
		}
	}
	//create a random filename and write to disk.
	newFilePath := TempFileName()
	report.WriteToFile(newFilePath)
	os.Remove(filePath)

	return newFilePath, nil
}
