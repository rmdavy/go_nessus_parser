package nessusTools

//Host model is used for storing nessus issues associated with a specific host
type Host struct {
	name   string
	issues []issue
}

type issue struct {
	issueName string
	issueID   string
	port      string
	output    []string // if Listing by issue name, output is used to store the list of affected hosts
}

func (h Host) deleteIssue(i int) Host {
	if i < len(h.issues) {
		h.issues = append(h.issues[:i], h.issues[i+1:]...)
	}
	return h
}

//RemoveDuplicateIssues remoevs duplicate issues from the issues []string of a specific host.
func (h Host) RemoveDuplicateIssues() Host {
	encountered := map[string]int{}
	var newHost Host
	newHost.name = h.name

	for a, b := range h.issues {
		if encountered[b.issueName] != 0 {
			//h.issues[encountered[b.issueName]].output = append(h.issues[encountered[b.issueName]].output, b.output...)
			// Append duplicate issue.output to first sighting of issueName
			for c, d := range newHost.issues {
				if b.issueName == d.issueName {
					newHost.issues[c].output = append(newHost.issues[c].output, b.output...)
				}
			}
			h = h.deleteIssue(a)
		} else {
			// Record this element as an encountered element.
			encountered[b.issueName] = a
			newHost.issues = append(newHost.issues, b)
		}
	}

	for a, b := range h.issues {
		output := []string{}
		for _, c := range b.output {
			output = append(output, c)
		}
		h.issues[a].output = RemoveDuplicates(output)
	}

	return newHost
}
