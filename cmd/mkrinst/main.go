package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(argv []string) error {
	conffile := "/etc/config/qpkg.conf"
	confBytes, err := ioutil.ReadFile("./qpkg.conf.txt")
	if err != nil {
		return err
	}
	confAllBytes, err := ioutil.ReadFile(conffile)
	if err != nil {
		return err
	}
	resultStr := updateConf(string(confAllBytes), string(confBytes))
	if resultStr == string(confAllBytes) {
		return nil
	}
	return ioutil.WriteFile(conffile, []byte(resultStr), 0644)
}

var reg = regexp.MustCompile(`\[mackerel-agent\][^[]+`)

func updateConf(from, confStr string) string {
	if !strings.HasPrefix(confStr, "[mackerel-agent]") {
		log.Fatalf("invalid confStr: %s", confStr)
	}
	if !strings.Contains(from, "[mackerel-agent]") {
		if !strings.HasSuffix(from, "\n") {
			from += "\n"
		}
		return from + confStr
	}
	return reg.ReplaceAllString(from, confStr)
}
