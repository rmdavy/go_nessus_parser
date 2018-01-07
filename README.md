# go_nessus_parser
This program is used to combine and edit Nessus files. Only the .nessus extension is supported. When combining reports, removing issues or removing hosts a new nessus file will be created and saved in the current directory.


Options:    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--r, --reports    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;location of nessus reports to be worked on.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--i, --issues    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;issues to be removed (must be pluginItem value as found in nessus file)    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--ic, --issuesContains    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;issues to be removed (removes all issues continaing supplied strings)    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--h, --host    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Hosts to be removed.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--s, --search    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;display all output where pluginItem contains the supplied string (i.e. --search SSL will return all ssl issues.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--f, --file    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;save issue search/combo text to a file, saves to a file named output in current directory.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--l, --list    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;List type of issue in the supplied reports.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--ciphers    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Print a selection of common ssl issues along with their associated ciphers.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--sslIssues    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Print a selection of common ssl issues and the hosts affected.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--summary    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Print a summary of useful information from the supplied reports.    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;--help    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Display this help page.    



Example usage:    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;./nessusToolCLI --reports ~/Desktop/*nessus (combine all listed reports into a single nessus file)    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f (List all output from SSL issues and save to file    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f --issues SSL\ Session\ Resume\ Supported --hosts 127.0.0.1 (as above but also removes ssl session resume issues and host 127.0.0.1)    
