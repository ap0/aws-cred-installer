package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		log.Println("must specify an argument")
		os.Exit(1)
	}

	sectionToUse := os.Args[1]

	awsPath := path.Join(os.Getenv("HOME"), ".aws")
	if os.Getenv("AWS_CRED_PATH") != "" {
		awsPath = os.Getenv("AWS_CRED_PATH")
	}
	b, err := ioutil.ReadFile(path.Join(awsPath, "credentials"))
	if err != nil {
		log.Println("Error reading credential file: %s", err.Error())
		os.Exit(1)
	}

	iniFile, err := ini.Load(b)
	if err != nil {
		log.Println("Error parsing INI file: %s", err.Error())
	}

	for _, section := range iniFile.Sections() {
		if section.Name() == sectionToUse {
			for key, value := range section.KeysHash() {
				fmt.Printf("export %s=%s\n", strings.ToUpper(key), value)
			}
			return
		}
	}

	log.Printf("Section '%s' not found!\n", sectionToUse)
	os.Exit(1)

}
