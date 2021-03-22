package async

import (
	"strconv"
	"testing"
)

func TestAsync(t *testing.T) {
	//TODO: write tests for async interface and functionality
}

var tests = []func(t *testing.T){
	TestAsync,
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
