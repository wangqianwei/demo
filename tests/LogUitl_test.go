package test

import (
	"demo/models/utils"
	"testing"
)

func TestError(t *testing.T) {

	utils.LogInfo("%s", "Hello World")
	utils.LogError("%s", "Hello World")
}

func TestInfo(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
