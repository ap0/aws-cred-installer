package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/go-ini/ini"
)

func main() {
	awsCredPath := flag.String("aws-cred-path", path.Join(os.Getenv("HOME"), ".aws", "credentials"), "Absolute path to AWS credentials file")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: aws-cred-installer [options] account\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	sectionToUse := flag.Arg(0)
	if sectionToUse == "" {
		flag.Usage()
		os.Exit(1)
	}

	iniFile, err := ini.Load(*awsCredPath)
	if err != nil {
		log.Println("Error parsing INI file: ", err.Error())
		os.Exit(1)
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
