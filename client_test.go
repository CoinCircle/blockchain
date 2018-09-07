package blockchain

import (
	"strings"
	"testing"
)

// Test recovery from a 504 timeout error
func TestRecoverFrom504(t *testing.T) {
	c, err := New(&Options{
		UseTestnet: true,
		APIRoot:    "https://httpstat.us",
	})
	if err != nil {
		t.Fatal(err)
	}
	rsp := &Block{}
	e := c.loadResponse("/504?sleep=15000", rsp, false)
	if e != nil {
		t.Log(e.Error())
	}

	if strings.Index(e.Error(), "Client.Timeout exceeded") > 0 {
		// This was a timeout success
		t.Log("Got timeout error")
	} else {
		// This was an error, but not a timeout
		t.Fail()
	}
}
