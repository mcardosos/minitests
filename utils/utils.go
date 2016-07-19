package utils

import (
	"os/exec"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
)

func ToDateTimeRFC1123(s string) date.TimeRFC1123 {
	t, _ := time.Parse(time.RFC1123, strings.ToUpper(s))
	return date.TimeRFC1123{t}
}

func ToDateTime(s string) date.Time {
	t, _ := time.Parse(time.RFC3339, strings.ToUpper(s))
	return date.Time{t}
}

func GetBaseURI() string {
	/*
		if strings.HasPrefix(baseURI, "https") {
			baseURI = "http://localhost:3000"
		} else {
			baseURI += ":3000"
		}
		return baseURI
	*/
	return "http://localhost:3000"
}

func StartTestServer() (*exec.Cmd, error) {
	testServerPath := "../../dev/TestServer/server"
	install := exec.Command("npm", "install")
	install.Dir = testServerPath
	server := exec.Command("node", "startup/www")
	server.Dir = testServerPath
	if err := install.Run(); err != nil {
		return nil, err
	}
	if err := server.Start(); err != nil {
		return nil, err
	}
	return server, nil
}
