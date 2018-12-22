package main

import "testing"

func TestUpdateConf(t *testing.T) {
	testCases := []struct {
		name, from, to string
	}{
		{
			name: "append",
			from: `[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
`,
			to: `[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
[mackerel-agent]
Name = mackerel-agent
Version = 0.58.2
Author = mackerelio
Date = 2018-11-27
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
`,
		},

		{
			name: "replace last",
			from: `[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
[mackerel-agent]
Name = mackerel-agent
Version = 0.57.0
Author = mackerelio
Date = 2018-10-10
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
`,
			to: `[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
[mackerel-agent]
Name = mackerel-agent
Version = 0.58.2
Author = mackerelio
Date = 2018-11-27
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
`,
		},
		{
			name: "replace first",
			from: `[mackerel-agent]
Name = mackerel-agent
Version = 0.57.0
Author = mackerelio
Date = 2018-10-10
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
`,
			to: `[mackerel-agent]
Name = mackerel-agent
Version = 0.58.2
Author = mackerelio
Date = 2018-11-27
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
[git]
Name = git
Class = null
Status = complete
Display_Name = Git
Sys_App = 1
Version = 1.8.4.2-1
Author = Open
QPKG_File = git.qpkg
Date = 2015-05-04
Shell = /share/MD0_DATA/.qpkg/git/git.sh
Install_Path = /share/MD0_DATA/.qpkg/git
Enable = TRUE
`,
		},
	}

	confStr := `[mackerel-agent]
Name = mackerel-agent
Version = 0.58.2
Author = mackerelio
Date = 2018-11-27
Shell = /share/MD0_DATA/.mackerel-agent/run.sh
Install_Path = /share/MD0_DATA/.mackerel-agent
QPKG_File = /share/MD0_DATA/.mackerel-agent/DUMMY
Enable = TRUE
`
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := updateConf(tc.from, confStr)
			if out != tc.to {
				t.Errorf("something went wrong.\nout:\n%s\nexpect:\n%s\n", out, tc.to)
			}
		})
	}
}
