package zautorestreport_test

import (
	"Tests/go-acceptance-tests/utils"
	. "Tests/go-acceptance-tests/zautorestreport"
	"fmt"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type AutorestReportSuite struct{}

var _ = check.Suite(&AutorestReportSuite{})

func TestMain(m *testing.M) {
	server, err := utils.StartTestServer()
	if err != nil {
		fmt.Println(err)
	}
	exitCode := m.Run()
	server.Process.Kill()
	os.Exit(exitCode)
}

func (s *AutorestReportSuite) TestGetReport(c *check.C) {
	GetReport()
}
