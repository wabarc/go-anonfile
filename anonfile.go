// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package anonfile // import "github.com/wabarc/go-anonfile"

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/wabarc/helper"
)

const (
	ENDPOINT = "https://api.anonfiles.com/upload"
)

type Anonfile struct {
	Client *http.Client
}

type size struct {
	Bytes    int
	Readable string
}

type metadata struct {
	ID, Name string
	Size     size
}

type url struct {
	Full, Short string
}

type file struct {
	Metadata metadata
	URL      url
}

type data struct {
	File file
}

type erro struct {
	Message, Type string

	Code int
}

type Anonfiles struct {
	Data   data
	Status bool
	Error  erro
}

func NewAnonfile(client *http.Client) *Anonfile {
	if client == nil {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &Anonfile{
		Client: client,
	}
}

func (anon *Anonfile) Upload(path string) (*Anonfiles, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if size := helper.FileSize(path); size > 5368709120 {
		return nil, fmt.Errorf("File too large, size: %d GB", size/1024/1024/1024)
	}

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		field := "file"
		part, err := m.CreateFormFile(field, filepath.Base(file.Name()))
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := anon.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parse(resp)
}

func (anon *Anonfile) Info(path string) (string, error) {
	return "TODO", nil
}

func (anon *Anonfiles) Full() string {
	if anon != nil {
		return anon.Data.File.URL.Full
	}
	return ""
}

func (anon *Anonfiles) Short() string {
	if anon != nil {
		return anon.Data.File.URL.Short
	}
	return ""
}

func parse(resp *http.Response) (anon *Anonfiles, err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &anon); err != nil {
		return nil, err
	}
	if !anon.Status {
		return nil, fmt.Errorf(`Upload failed, message: %s`, anon.Error.Message)
	}

	return anon, nil
}
