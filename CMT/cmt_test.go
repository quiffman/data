package cmt_testing

import (
	"path/filepath"
	"testing"

	"github.com/GeoNet/data/CMT/list"
)

func TestCMTs_File(t *testing.T) {

	var cmts list.CMTs
	if err := list.ReadFile(filepath.Join(data, list.CMTsFile), &cmts); err != nil {
		t.Fatal(err)
	}
}

//  vim: set ts=4 sw=4 tw=0 :
