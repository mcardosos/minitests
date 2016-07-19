package formdatagroup_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	chk "gopkg.in/check.v1"

	. "Tests/body-formdata"
	"Tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type FormdataSuite struct{}

var _ = chk.Suite(&FormdataSuite{})

var formdataClient = getFormdataClient()

func getFormdataClient() FormdataClient {
	c := NewFormdataClient()
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

func (s *FormdataSuite) TestUploadFileViaBody(c *chk.C) {
	f, err := ioutil.ReadFile("../sample.png")
	c.Assert(err, chk.IsNil)
	res, err := formdataClient.UploadFileViaBody(ioutil.NopCloser(bytes.NewReader(f)))
	c.Assert(err, chk.IsNil)
	buf := new(bytes.Buffer)
	buf.ReadFrom(*res.Value)
	b := buf.Bytes()
	c.Assert(len(b), chk.Equals, len(f))
	c.Assert(string(b), chk.Equals, string(f))
}

func (s *FormdataSuite) TestUploadFile(c *chk.C) {
	f, err := ioutil.ReadFile("../sample.png")
	c.Assert(err, chk.IsNil)
	res, err := formdataClient.UploadFile(ioutil.NopCloser(bytes.NewReader(f)), "samplefile")
	c.Assert(err, chk.IsNil)
	buf := new(bytes.Buffer)
	buf.ReadFrom(*res.Value)
	b := buf.Bytes()
	c.Assert(len(b), chk.Equals, len(f))
	c.Assert(string(b), chk.Equals, string(f))
}
