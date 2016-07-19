package filegroup_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	chk "gopkg.in/check.v1"

	. "Tests/body-file"
	"Tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type FileSuite struct{}

var _ = chk.Suite(&FileSuite{})

var filesClient = getFileClient()

func getFileClient() FilesClient {
	c := NewFilesClient()
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

func (s *FileSuite) TestGetFile(c *chk.C) {
	res, err := filesClient.GetFile()
	c.Assert(err, chk.IsNil)
	b, err := ioutil.ReadAll(res.Body)
	defer func() { (res.Body).Close() }()
	c.Assert(err, chk.IsNil)
	c.Assert(len(b), chk.Equals, 8725)
}

func (s *FileSuite) TestGetEmptyFile(c *chk.C) {
	res, err := filesClient.GetEmptyFile()
	c.Assert(err, chk.IsNil)
	b, err := ioutil.ReadAll(res.Body)
	defer func() { (res.Body).Close() }()
	c.Assert(err, chk.IsNil)
	c.Assert(len(b), chk.Equals, 0)
}

func (s *FileSuite) TestGetFileLarge(c *chk.C) {
	res, err := filesClient.GetFileLarge()
	c.Assert(err, chk.IsNil)
	b, err := ioutil.ReadAll(res.Body)
	defer func() { (res.Body).Close() }()
	c.Assert(err, chk.IsNil)
	c.Assert(len(b), chk.Equals, 3000*1024*1024)
}
