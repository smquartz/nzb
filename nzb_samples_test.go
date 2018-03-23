package nzb

import "testing"

func TestMoreNzb(t *testing.T) {
	notNzbPath := "./samples/NotNzb.nzb"
	noFilesNzbPath := "./samples/NoFiles.nzb"

	// not nzb
	_, err := FromFile(notNzbPath)
	if err == nil {
		t.Errorf("Failed to error on invalid NZB file", err)
	}

	// no files nzb
	n, err := FromFile(noFilesNzbPath)
	if err != nil {
		t.Errorf("Failed to open noFiles test NZB: %v", err)
	}

	if len(n.Files) != 0 {
		t.Errorf("Wrong number of files: %d", len(n.Files))
	}
}
