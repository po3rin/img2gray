package img2gray_test

import (
	"testing"

	"github.com/po3rin/img2gray"
)

func TestTraverseGray(t *testing.T) {
	tests := []struct {
		name    string
		dirPath string
	}{
		{
			name:    "target testimg",
			dirPath: "testimg",
		},
		{
			name:    "target testimg/",
			dirPath: "testimg/",
		},
		{
			name:    "target .",
			dirPath: ".",
		},
	}
	for _, tt := range tests {
		err := img2gray.TraverseGray(tt.dirPath)
		if err != nil {
			t.Fatalf("test %v failed. unexpected error: %v", tt.name, err)
		}
	}
}
