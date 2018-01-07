# go_nessus_parser
This program is used to combine and edit Nessus files. Only the .nessus extension is supported. When combining reports, removing issues or removing hosts a new nessus file will be created and saved in the current directory.


Options:    
        --r, --reports    
                location of nessus reports to be worked on.    
        --i, --issues    
                issues to be removed (must be pluginItem value as found in nessus file)    
        --ic, --issuesContains    
                issues to be removed (removes all issues continaing supplied strings)    
        --h, --host    
                Hosts to be removed.    
        --s, --search    
                display all output where pluginItem contains the supplied string (i.e. --search SSL will return all ssl issues.    
        --f, --file    
                save issue search/combo text to a file, saves to a file named output in current directory.    
        --l, --list    
                List type of issue in the supplied reports.    
        --ciphers    
                Print a selection of common ssl issues along with their associated ciphers.    
        --sslIssues    
                Print a selection of common ssl issues and the hosts affected.    
        --summary    
                Print a summary of useful information from the supplied reports.    
        --help    
                Display this help page.    



Example usage:    
        ./nessusToolCLI --reports ~/Desktop/*nessus (combine all listed reports into a single nessus file)    
        ./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f (List all output from SSL issues and save to file    
        ./nessusToolCLI --reports ~/Desktop/*nessus --search SSL --f --issues SSL\ Session\ Resume\ Supported --hosts 127.0.0.1 (as above but also removes ssl session resume issues and host 127.0.0.1)    
