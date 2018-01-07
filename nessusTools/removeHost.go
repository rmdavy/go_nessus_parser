package nessusTools

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func RemoveHosts(filePath string, hosts []string) (string, error) {
	fmt.Println("Removing Hosts:")
	for _, b := range hosts {
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
			for _, b := range hosts {
				if host.SelectAttrValue("name", "Not Found") == b {
					report.RemoveChild(host)
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
