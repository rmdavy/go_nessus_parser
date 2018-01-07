package nessusTools

import (
	"strings"
)

func sslSearch(c string) []string {
	out := []string{}

	if strings.Contains(c, "This port supports") && !strings.Contains(c, "sessions") {
		a := strings.Split(c, " ")
		b := strings.Split(a[len(a)-1], "/")
		// the last element of each slice will have a trailing . this must be removed so duplicates can be correctly identified and deleted.
		c := b[len(b)-1] //
		d := c[:len(c)-1]
		b = append(b[:len(b)-1], string(d))
		out = b
	} else {

		list := strings.Split(c, " ")

		for _, c := range list { //The following if statement filters out irrelevant values and just grabs the ciphers from ssl issues
			if strings.Contains(c, "-") && len(c) > 1 && !strings.Contains(c, "Orga") && !strings.Contains(c, "Country") &&
				!strings.Contains(c, "Vali") && !strings.Contains(c, "bit") && !strings.Contains(c, "Sub") && !strings.Contains(c, "Issu") &&
				!strings.Contains(c, "Sig") && !strings.Contains(c, "=") && !strings.Contains(c, "Comm") {
				out = append(out, c)
			}
		}
	}

	return out
}
