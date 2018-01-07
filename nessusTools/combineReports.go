package nessusTools

import (
	"os"
	"strings"

	"github.com/beevik/etree"
)

func CombineReports(filePath []string) (string, error) {

	//report for hosts to be added too
	newReport := etree.NewDocument()
	if err := newReport.ReadFromFile(filePath[0]); err != nil {
		return "", err
	}

	newRoot := newReport.SelectElement("NessusClientData_v2")

	for i := 0; i < len(filePath); i++ {
		//current report hosts are being added from
		report := etree.NewDocument()
		if err := report.ReadFromFile(filePath[i]); err != nil {
			return "", err
		}

		root := report.SelectElement("NessusClientData_v2")

		var hosts []*etree.Element

		for _, report := range root.SelectElements("Report") { //Select report branch to iterate over
			for _, host := range report.SelectElements("ReportHost") { //Select Reporthost branch to iterate over
				hosts = append(hosts, host)
			}
		}

		for _, report := range newRoot.SelectElements("Report") {
			for _, b := range hosts {
				report.AddChild(b)
			}
		}
	}

	//create a random filename and write to disk.
	newFilePath := TempFileName()
	newReport.WriteToFile(newFilePath)
	os.Remove(strings.Join(filePath, ""))

	return newFilePath, nil
}
