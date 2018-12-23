package main

import (
	"fmt"
	"os"
	"text/template"
)

func doAgentConf(argv []string) (err error) {
	if len(argv) < 2 {
		return fmt.Errorf("install_path or mackerel_api_key aren't specified")
	}
	installPath := argv[0]
	mackerelAPIKey := argv[1]
	confFile := "./mackerel-agent.conf"
	confTmpl := "./mackerel-agent.conf.tmpl"

	tmpl, err := template.ParseFiles(confTmpl)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(confFile, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer func() {
		if e := f.Close(); e != nil {
			err = e
		}
	}()

	return tmpl.Execute(f, struct{ InstallPath, MackerelAPIKey string }{
		InstallPath:    installPath,
		MackerelAPIKey: mackerelAPIKey,
	})
}
