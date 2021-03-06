package durationgroup_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	chk "gopkg.in/check.v1"

	. "Tests/Generated/body-duration"
	"Tests/go-acceptance-tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type DurationSuite struct{}

var _ = chk.Suite(&DurationSuite{})

var durationClient = getDurationClient()

func getDurationClient() DurationClient {
	c := NewDurationClient()
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

func (s *DurationSuite) TestGetInvalidDuration(c *chk.C) {
	res, err := durationClient.GetInvalid()
	if err != nil {
		c.SucceedNow()
	}
	_, err = time.ParseDuration(*res.Value)
	c.Assert(err, chk.NotNil)
}

func (s *DurationSuite) TestGetNullDuration(c *chk.C) {
	res, err := durationClient.GetNull()
	c.Assert(err, chk.IsNil)
	c.Assert(res.Value, chk.IsNil)
}

func (s *DurationSuite) TestGetPositiveDuration(c *chk.C) {
	res, err := durationClient.GetPositiveDuration()
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, "P3Y6M4DT12H30M5S")
}

func (s *DurationSuite) TestPutPositiveDuration(c *chk.C) {
	_, err := durationClient.PutPositiveDuration("P123DT22H14M12.011S")
	c.Assert(err, chk.IsNil)
}
