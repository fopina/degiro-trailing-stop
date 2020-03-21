package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fopina/degiro-trailing-stop/degiro"
	flag "github.com/spf13/pflag"
)

var version string = "DEV"
var date string

func main() {
	versionPtr := flag.BoolP("version", "v", false, "display version")
	usernamePtr := flag.StringP("username", "u", LookupEnvOrString("DEGIRO_USERNAME", ""), "degiro username")
	passwordPtr := flag.StringP("password", "p", LookupEnvOrString("DEGIRO_PASSWORD", ""), "degiro password")

	flag.Parse()

	if *versionPtr {
		fmt.Println("Version: " + version + " (built on " + date + ")")
		return
	}

	c := degiro.NewAPIClient(*usernamePtr, *passwordPtr)
	fmt.Println(c.Login())
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}
