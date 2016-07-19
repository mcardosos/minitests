package autorestreport_test

import (
	"Tests/utils"
	. "Tests/zautorestreport"
	"fmt"
	"os"
	"testing"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type AutorestReportSuite struct{}

var _ = chk.Suite(&AutorestReportSuite{})

func TestMain(m *testing.M) {
	server, err := utils.StartTestServer()
	if err != nil {
		fmt.Println(err)
	}
	exitCode := m.Run()
	server.Process.Kill()
	os.Exit(exitCode)
}

func (s *AutorestReportSuite) TestGetReport(c *chk.C) {
	GetReport()
}
