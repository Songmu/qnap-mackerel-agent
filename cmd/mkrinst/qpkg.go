package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func doQpkgConf(argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("no install_path specified")
	}
	installPath := argv[0]
	confFile := "/etc/config/qpkg.conf"
	confTmpl := "./qpkg.conf.tmpl"

	tmpl, err := template.ParseFiles(confTmpl)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, struct{ InstallPath string }{InstallPath: installPath})

	confAllBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return err
	}
	resultStr := updateConf(string(confAllBytes), buf.String())
	if resultStr == string(confAllBytes) {
		return nil
	}
	tmp, err := ioutil.TempFile(filepath.Dir(confFile), "mackerel-agent")
	if err != nil {
		return err
	}
	defer func(fname string) {
		tmp.Close()
		os.Remove(fname)
	}(tmp.Name())

	st, err := os.Stat(confFile)
	if err != nil {
		return err
	}
	if err := os.Chmod(tmp.Name(), st.Mode()); err != nil {
		return err
	}
	fmt.Fprint(tmp, resultStr)
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), confFile)
}

var reg = regexp.MustCompile(`\[mackerel-agent\][^[]+`)

func updateConf(from, confStr string) string {
	if !strings.HasPrefix(confStr, "[mackerel-agent]") {
		log.Fatalf("invalid confStr: %s", confStr)
	}
	if !strings.HasSuffix(confStr, "\n") {
		confStr += "\n"
	}
	if !strings.Contains(from, "[mackerel-agent]") {
		if !strings.HasSuffix(from, "\n") {
			from += "\n"
		}
		return from + confStr
	}
	return reg.ReplaceAllString(from, confStr)
}
