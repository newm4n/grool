package pkg

import (
	"fmt"
	"github.com/juju/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Resource interface {
	Load() ([]byte, error)
	String() string
}

func NewReaderResource(reader io.Reader) Resource {
	return &ReaderResource{Reader: reader}
}

type ReaderResource struct {
	Reader io.Reader
}

func (res *ReaderResource) Load() ([]byte, error) {
	return ioutil.ReadAll(res.Reader)
}

func (res *ReaderResource) String() string {
	return "Reader resource. Source unknown."
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
			return nil, errors.Trace(err)
		}
		res.Bytes = data
		return res.Bytes, nil
	}
}

func (res *FileResource) String() string {
	return fmt.Sprintf("File resource at %s", res.Path)
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

func (res *BytesResource) String() string {
	return fmt.Sprintf("Byte array resources %d bytes", len(res.Bytes))
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

func (res *UrlResource) String() string {
	return fmt.Sprintf("URL resource at %s", res.Url)
}

func (res *UrlResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	} else {
		resp, err := http.Get(res.Url)
		if err != nil {
			return nil, errors.Trace(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Trace(err)
		}
		res.Bytes = data
		return res.Bytes, nil
	}
}
