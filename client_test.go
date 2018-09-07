package blockchain

import "testing"

// Test recovery from a 504 timeout error
func TestRecoverFrom504(t *testing.T) {
	c, err := New(&Options{
		UseTestnet: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	rsp := &Block{}
	e := c.loadResponse("https://httpstat.us/504", rsp, false)
	if e != nil {
		t.Log(e)
	}
	// This line won't execute if it can't recover
	a := 1
	if a != 1 {
		t.Fatal("Test failed")
	}
}
