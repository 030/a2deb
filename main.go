package main

import (
	"flag"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := flag.String("app", "", "The name of the Golang application")
	description := flag.String("description", "", "The description of the Golang application")
	maintainer := flag.String("maintainer", "", "The maintainer of the to be created debian package, e.g. user \"<user@some-domain>\"")
	version := flag.String("version", "", "The version of the Golang application")

	flag.Parse()

	fmt.Println("app:", *app)
	fmt.Println("description:", *description)
	fmt.Println("maintainer:", *maintainer)
	fmt.Println("version:", *version)

	out, err := exec.Command("bash", "-c", "mkdir -p "+*app+"_"+*version+"-0/usr/local/bin && cp "+*app+" "+*app+"_"+*version+"-0/usr/local/bin/"+*app+" && mkdir "+*app+"_"+*version+"-0/DEBIAN && echo -e 'Package: "+*app+"\nVersion: "+*version+"\nArchitecture: amd64\nMaintainer: "+*maintainer+"\nDescription: "+*app+"\n "+*description+"' > "+*app+"_"+*version+"-0/DEBIAN/control && dpkg-deb --root-owner-group --build "+*app+"_"+*version+"-0").CombinedOutput()
	outString := string(out)

	log.Info(outString)

	if err != nil {
		log.Fatal(err)
	}
}
