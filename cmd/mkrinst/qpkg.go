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

	f, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	origMode := fi.Mode()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	resultStr := updateConf(string(b), buf.String())
	if resultStr == string(b) {
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

	if err := tmp.Chmod(origMode); err != nil {
		return err
	}
	if _, err := fmt.Fprint(tmp, resultStr); err != nil {
		return err
	}
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
