package proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()
	t.Log(res)

	if res != "pre:real:after" {
		t.Fail()
	}
}
