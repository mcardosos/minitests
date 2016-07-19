package custombaseurlgroup_test

import (
	"fmt"
	"os"
	"testing"

	chk "gopkg.in/check.v1"

	. "Tests/custom-baseurl"
	"Tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type CustomBaseURLGroupSuite struct{}

var _ = chk.Suite(&CustomBaseURLGroupSuite{})

var custombaseuriClient = getCustomBaseURIClient()

func getCustomBaseURIClient() PathsClient {
	c := NewPathsClient("local")
	c.BaseURI = utils.GetBaseURI()
	return c
}

func TestMain(m *testing.M) {
	server, err := utils.StartTestServer()
	if err != nil {
		fmt.Println(err)
	}
	exitCode := m.Run()
	server.Process.Kill()
	os.Exit(exitCode)
}

func (s *CustomBaseURLGroupSuite) TestGetEmpty(c *chk.C) {
	_, err := custombaseuriClient.GetEmpty()
	c.Assert(err, chk.IsNil)
}
