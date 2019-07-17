package pkg

import (
	"io/ioutil"
	"net/http"
)

type Resource interface {
	Load() ([]byte, error)
}

func NewFileResource(path string) Resource {
	return &FileResource{
		Path: path,
	}
}

type FileResource struct {
	Path  string
	Bytes []byte
}

func (res *FileResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	} else {
		data, err := ioutil.ReadFile(res.Path)
		if err != nil {
			return nil, err
		}
		res.Bytes = data
		return res.Bytes, nil
	}
}

func NewBytesResource(bytes []byte) Resource {
	return &BytesResource{
		Bytes: bytes,
	}
}

type BytesResource struct {
	Bytes []byte
}

func (res *BytesResource) Load() ([]byte, error) {
	return res.Bytes, nil
}

func NewUrlResource(url string) Resource {
	return &UrlResource{
		Url: url,
	}
}

type UrlResource struct {
	Url   string
	Bytes []byte
}

func (res *UrlResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	} else {
		resp, err := http.Get(res.Url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		res.Bytes = data
		return res.Bytes, nil
	}
}
