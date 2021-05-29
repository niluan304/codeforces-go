// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`1
20
2
10 20`,
			`20 20`,
		},
		{
			`3
3 2 2
5
1 5 7 7 9`,
			`3 1`,
		},
		{
			`4
1 3 5 7
4
7 5 3 1`,
			`1 1`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
