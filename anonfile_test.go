// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package anonfile // import "github.com/wabarc/go-anonfile"

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	content := make([]byte, 5000)
	tmpfile, err := ioutil.TempFile("", "go-anonfile-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}

	if _, err := NewAnonfile(nil).Upload(tmpfile.Name()); err != nil {
		t.Fatal(err)
	}
}
