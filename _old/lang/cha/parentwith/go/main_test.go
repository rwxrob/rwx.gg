package main

import "testing"

func Test_ParentWith(t *testing.T) {
	t.Logf("%v %v\n", "file1", "testing/foo/bar")
	if s := ParentWith("file1", "testing/foo/bar"); s != "" {
		t.Log(s)
		t.Fail()
	}
}
