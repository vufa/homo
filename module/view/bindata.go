// Code generated by go-bindata.
// sources:
// assets/component/homo.js
// assets/component/styles/input.css
// assets/component/styles/reply.css
// assets/component/styles/says.css
// assets/component/styles/setup.css
// assets/component/styles/typing.css
// assets/index.html
// DO NOT EDIT!

package view

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _componentHomoJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x3a\xcb\x72\xdc\xb8\x76\xeb\xab\xaf\x38\x66\x55\x24\xb6\xd5\xea\x96\xe7\xd6\xdc\x45\xcb\x2d\x95\x2d\xdb\xb1\x53\xbe\xf6\x8c\xe5\x94\x6b\x4a\x51\xcd\xa0\xc9\xd3\x4d\x44\x6c\x80\x06\xc0\x96\xda\x23\x2d\xb2\x48\x55\xf6\xf9\x86\xac\xb3\xcb\x32\x7f\x93\xa9\x7c\x46\x0a\x00\x1f\x00\x09\xb6\xa5\xd4\xe5\x42\x62\x13\x38\x0f\xe0\xbc\x0f\x30\x9d\x42\xc2\x05\xc2\xb2\x64\x89\xa2\x9c\xed\xd5\x2f\xf0\x96\xaf\x79\x9c\x70\xa6\x08\x65\x28\xc6\x20\x31\x5f\x8e\x81\x17\x7a\x50\x8e\xe0\xf7\x3d\x00\x80\xe9\xb4\xfe\x62\x7e\x56\xef\x30\x07\xb5\x2d\x90\x2f\x9b\x0f\x4f\xe6\x73\x88\x4a\x96\xe2\x92\x32\x4c\x23\x38\x6b\x46\x66\xf0\xfb\xfd\x89\x01\x26\x8c\xae\x89\xfe\xf8\x99\xae\x11\xe6\xf5\x8c\x89\xff\xfd\xee\x0e\x7e\x38\x3e\x3e\xd1\x94\x33\x7e\x03\x39\x67\x2b\xa0\x0a\x14\xb9\x46\x09\x8a\x57\x58\x10\x92\x8c\x28\x58\x94\x8b\x45\x8e\x63\x20\xb9\xe4\x20\x51\x01\x65\x70\x7e\x71\x61\xc8\x69\x0e\x2f\x0a\xc4\xd4\x21\xd5\x7e\xbb\xbb\x83\x1f\x0d\x91\x14\x73\xb2\x85\x02\x85\xc6\x28\x48\xa2\xf4\x5e\x28\x0e\x92\xae\xcb\x5c\x13\x52\x19\xc2\x9a\x24\x19\x65\x08\x91\xda\x16\x94\xad\x22\x43\xe0\x86\xa6\x28\x5e\x6e\x1d\xf4\xf5\x17\xbd\x06\x83\x9c\xa4\x29\x10\xc8\xa9\x52\x39\x02\xde\x2a\x41\x34\x94\xca\x34\x01\xcb\xbb\x59\xd3\x9a\x5c\x23\xc8\x52\x18\x62\x5b\x48\x39\x3b\x50\xb0\x10\x48\xae\x0d\x21\x49\x53\xfc\x89\xa4\x29\x65\x2b\x87\x98\xfb\xf5\xee\x0e\xfe\x62\x08\x16\xd5\x07\xce\x60\xc1\x55\x66\x40\x25\xf0\xa5\xbb\x5d\x56\x94\x02\x13\x92\xe7\xef\x98\x42\xbd\xe8\x4a\xaa\x35\xee\xc0\xe0\xdd\x1d\x58\xa9\xb0\x72\xbd\x40\xa1\x71\x52\x77\x5c\xaf\x08\x41\xe0\x1a\xf5\x30\xa6\x40\x58\x0a\x0b\xc1\xcb\x55\xa6\x60\x41\x92\x6b\x28\x0b\xce\x40\xa0\x54\x44\x28\xc3\x02\x65\x45\xa9\xce\x49\x9e\xeb\xe1\x37\xcc\xa1\xdf\x1d\xb9\xbb\x83\x25\xc9\x25\x1a\x06\x64\xc6\xcb\x3c\x85\x1b\x84\x94\xca\x42\x4b\x8f\x30\x8b\x0b\x96\x14\xf3\xf4\x6c\xcf\x60\xdf\x10\x01\x52\x11\xa6\xf7\xe3\x05\x93\x37\x28\x60\x0e\x11\x4d\x30\x32\x58\x6a\x4e\xe1\x26\x43\xbd\xf1\xbc\x66\x0d\x12\xce\x36\x1c\x68\xb5\x3e\x51\x16\x0a\xd3\x16\xe5\xaf\x76\x78\xae\xd5\x5a\xe3\xc9\x79\x42\x72\x58\xe3\x9a\x8b\x2d\x2c\xb9\x30\xe0\x28\xa4\xd1\x68\xf8\x87\x8b\x8f\x1f\x80\x2f\xfe\x19\x13\x55\xd9\xd3\xd1\xd1\x29\x7c\xf8\xf8\xf9\x35\x28\x2d\x11\x95\x51\x59\x8d\x83\x7e\x63\xf9\x16\x88\x94\x74\xc5\x30\x05\xce\x12\x1c\x1b\xd5\x94\x28\xa5\x46\xa7\xf7\x34\xe5\x28\x81\x71\xa5\x45\xca\x56\x68\x68\x6a\x34\xb5\xbd\xfe\xe9\x4f\x09\x67\x52\x89\x32\x51\x5c\x00\x23\x6b\x84\xb4\x14\x46\x29\x0a\x64\x35\xaa\xc9\x5e\x3d\xdf\x2e\x40\x2a\x2e\x48\x85\xcd\x4a\x5f\x43\xb8\x6b\x91\x7d\x01\xea\xfd\x30\xe0\x17\x16\xfa\x3c\xc3\xe4\x1a\xe6\x8d\xb3\x81\xb8\xf6\x23\xf5\x6c\x85\x52\x69\x31\x68\x75\x3c\xb2\xea\x78\x54\x91\x3e\xd2\x63\xd1\x49\x33\x5d\x89\xad\x03\xac\x1f\x97\xd4\x44\xa2\x7a\xa7\x70\x1d\x6b\xa8\xb1\xc1\x3b\x3a\x19\x9e\x2d\x70\xcd\x37\xd8\x00\x74\xa6\x0a\x54\xa5\x60\xa0\x44\x89\xcd\xf7\x7b\x48\x88\x4a\x32\x88\x51\x08\x2e\x46\x1d\x56\xf4\x16\xf3\x1c\x27\x66\x30\xf6\x86\xf4\x13\xfd\xc2\x4b\x2d\x35\xb1\x41\xd1\x0a\x8c\xe4\x39\xbf\x31\x3b\xad\xf7\x36\x25\x8a\x58\x26\xf3\xed\x04\xfe\xca\xa5\x82\x9c\x5e\x63\xbe\x05\xaa\x0e\x24\x2c\x30\x21\xa5\x44\xd8\xf2\xf2\x60\x83\x46\x76\x98\x5a\x7d\x29\x8c\xa0\x04\x5f\xeb\x41\x01\x19\x11\xe9\x51\x2a\xe8\x06\x27\xf0\x86\xdb\x3d\xd6\x04\xb6\xbc\x84\x84\x30\x6d\x25\x64\x91\xa3\x9d\xbc\x10\xfc\x46\xa2\x38\x90\x20\x31\x29\x05\x55\x5b\xe0\xc6\x4e\x84\x82\x8a\x9d\x4c\xb3\x82\x6c\x43\x05\x67\x6b\x64\x6a\x12\x79\xeb\x0b\xef\x9d\x31\xcf\x76\xf3\xcc\x5b\xe5\xf7\xbb\x5a\xf2\x62\x43\x68\x6e\x38\x9a\xf7\xb5\x27\x1e\xc1\xfe\x7e\xc8\x39\x9d\xc2\x71\x8b\xce\xf5\x3c\xef\x2f\xba\x0a\xe5\x8e\x46\x61\xa0\xb7\x54\x8b\x61\x0b\xf3\x86\xe5\x38\xcc\xe2\xfe\xbe\xb7\x5a\x6d\xcf\x93\x82\x08\x89\xde\xfc\xc9\xaa\xd2\x46\x9f\xb1\xd1\x68\x04\x77\x77\x0d\x82\xcb\xab\x93\xc6\xea\x0a\x81\x05\x11\x08\x0c\x6f\x15\x48\xb2\x41\x28\x38\x65\xb5\x5f\x6c\x91\x5c\xe8\x21\xd7\xa2\x24\xd9\x8e\x41\x60\x91\x6f\x5d\x9d\xa4\x4b\x88\x9f\x04\x57\x30\xaa\x24\xd4\x4a\x4d\xdb\x3c\x5d\x53\xe5\x78\x72\xcd\x80\xf4\x90\x05\x36\x6b\x92\x23\x5b\xa9\x0c\x4e\x03\xe2\x19\x79\xbb\x14\x02\x96\x19\x5d\xaa\x78\x54\xfb\x5e\xbe\xd1\xa1\x2f\x43\xe0\x79\xaa\x7d\x42\xbc\xa4\x42\xaa\x91\xdd\x8a\x26\x24\x16\x24\xc1\x3d\x97\xf1\x94\x1b\x4b\x32\xee\x96\x7e\x43\x58\x94\x4a\x71\x26\x4f\xac\xe7\x2c\x25\x8a\x2a\x12\xac\x50\xc9\x66\x5a\x3a\xf3\xd6\xe6\xf1\x3a\x9d\xc2\x6f\x95\xde\x58\x64\xbf\x41\x92\x13\x29\xad\xe7\xd4\x9e\x98\xe4\xd2\x3a\x6b\x63\x97\xa4\xa2\xe9\x21\x91\x64\x3b\xa1\x2c\xc9\xcb\x14\x65\x1c\x79\xe8\xa2\x51\x57\x87\xa6\x53\x13\x5c\x8c\xc3\xd7\x8b\xe1\x4b\x20\x26\x5b\xd1\x6f\x0a\x6f\x55\x49\x72\x2b\xe1\x8e\xa1\x15\xf9\xd6\x26\x5a\xf6\xd5\xfc\x3d\x5a\x0a\xc4\x25\x17\xeb\xe8\xff\x4f\x46\x7b\x80\x66\xab\xec\x26\x26\x19\xa7\x09\x3e\x84\x7e\x41\x93\xeb\xd6\x41\x8c\x5c\x61\x4d\x26\x13\xaa\x40\x66\xc4\xa4\x33\xd8\xd2\x08\x38\x90\x13\x4f\xcc\x8d\x16\x98\xb0\xba\xb7\x4b\xaf\x8a\x52\x66\xf1\xef\x92\x6c\x67\xd0\x9a\xc6\xcc\xfe\xbb\x1f\xd5\x7e\xa8\x36\xbb\x84\xaf\xb5\xe6\xd7\xf8\x5d\x8b\x09\x9a\xde\xb9\x9d\x3f\x14\xd2\x1e\x63\x76\xc1\xe8\xe5\xfb\x8b\xb1\xf5\x2f\x52\xe9\x00\x41\x97\xdb\x90\x15\x8e\x7a\x6b\xd2\x79\x6f\x59\x18\x63\x92\xaa\x5e\x48\x93\xd8\x4f\x8c\x36\xbf\xa7\x52\x4d\x48\x9a\x36\xca\xd9\x8c\x47\xa3\xd6\x43\xda\xb1\x2f\x82\x14\x30\x87\x94\x27\xa5\x71\xff\x89\x40\xa2\xf0\x75\x8e\xfa\x57\x1c\xa5\x74\x53\xc3\xb4\xf3\x2d\x95\x0f\xc4\x24\xf6\x35\x8d\x1b\x41\x8a\xca\xff\xb6\xec\x90\xa2\x40\x96\x9e\x67\x34\x4f\xe3\x16\x7e\xd4\x2e\x87\x32\xa9\x48\x9e\xbb\xc6\xac\xb5\xd5\xa4\x76\x36\xb1\xcf\xa8\xcd\xe4\xdf\x99\x41\x57\x36\x49\x93\x32\x76\x13\x0f\x83\xe8\xe1\x2b\x83\x3a\x3d\xed\x2f\xce\x7c\x76\xd7\xe6\x91\xf8\xac\xfd\xf9\x30\x09\xbd\x12\x22\x90\xf4\xe8\x68\x38\xad\x15\x2f\x94\x12\x74\x51\x2a\x8c\xa3\x22\x27\x09\x66\xda\x3f\x8a\x68\x0c\xd1\xff\xfe\xf7\xbf\xff\xcf\xbf\xfe\xc7\x1f\xff\xf5\x6f\x7f\xfc\xcb\x7f\x4e\x26\x93\x30\xa7\xee\xee\x36\x78\x83\xb4\x48\x9a\xbe\xde\x20\x53\x5a\x33\x90\xa1\x88\xa3\x6b\xdc\x16\x02\xa5\x8c\xc6\xce\x8e\x62\x37\xf3\x31\xee\x7b\xa5\x81\x84\x23\x22\xdf\xfd\x2f\x21\xc6\xc9\x35\x6e\xcf\x79\x8a\x30\x9f\xc3\xb3\x3f\x77\xb1\xe8\x07\x27\x85\x40\xcd\xc2\x2b\x5c\x92\x32\xd7\xc1\xa1\x37\xa7\x2a\x33\xad\xa2\xfc\x5c\x62\x89\xc6\x03\x99\x6c\x03\xce\x20\xc9\x91\x08\x5d\x36\xf2\x52\xc5\xce\xa4\x11\xcc\x9c\x82\xc1\xe6\x5d\x86\x59\xc5\xdb\xac\xde\x58\xcc\x82\xab\x1e\x51\x93\xb0\x10\xa9\x5e\x1a\x84\xae\x34\xbf\x96\x28\xb6\x17\x98\xa3\xce\xad\x5f\xe4\x79\x1c\x4d\x2c\xd5\x89\x24\xdb\x28\xc0\xbf\x87\xa6\xfd\x71\xd9\xbe\xd6\x51\xf5\x08\x9e\x5d\xed\x84\xdf\xdf\x77\x7e\x39\x56\x5d\x99\x96\x8c\xad\x5b\xee\xc7\x1b\xfd\x3c\x79\x08\x68\x1b\x4b\x46\x3d\x04\xfa\x39\x0b\x33\xe0\xba\x95\x8c\xa6\x29\xb2\x01\xf8\x5a\x26\xbd\x41\x92\xa6\x16\x6b\x3f\x93\xd6\xcf\xc1\x73\x59\x10\x66\xc3\xf2\xdc\x8f\xae\x95\x66\xd8\x28\x74\x7a\x00\x87\xd6\x3b\x6c\x48\x5e\x22\x1c\x42\xf4\x7c\xaa\x41\x4f\xa3\x71\x10\x71\xd8\xa9\xbb\xcf\x7d\x18\x30\x1c\x81\xc3\x73\xbd\xaa\xa2\x7e\x02\xaa\x82\xb7\x0a\x05\x23\xf9\x84\xb2\x0d\xbf\xc6\xf8\x60\x8d\x52\x92\x15\xce\xfc\x55\x05\x20\x75\x5c\xab\x1c\xdf\x90\x01\x25\x4e\x95\xad\x23\x78\xbd\xf2\x68\x40\xd0\xed\xfc\x38\xbc\x31\x50\x3b\x93\x99\xc3\x5b\x78\x03\xc0\xfa\xff\x0d\x9f\x55\xc5\xf3\xf0\x34\xbf\x5e\x9f\x75\x7e\x87\x25\xf4\x48\x5d\x73\xf4\x63\x0e\x91\xbf\xfe\xfb\xb6\x80\x71\xb6\x39\x1c\xbb\x1a\x9f\xeb\xcc\x74\xe2\xa1\x54\xdb\x1c\x27\x55\x3b\xe6\x25\x57\x8a\xaf\x35\xbd\x67\xc7\xc7\xc5\x6d\x14\xf2\xc7\x4b\x9e\x94\x32\x1e\xb9\x85\x53\xb7\x0b\x72\xd6\x09\x7d\x71\x67\x42\xeb\xf5\x9c\x60\x4a\x15\xd8\xa6\x55\xc5\x5d\x27\xd8\x7f\xb6\x63\x8f\x08\xf7\x16\x22\x18\xf0\x2b\x42\x74\x4d\x56\x94\x61\xb5\xcc\x25\x17\x10\xa7\x5c\x49\x98\xc3\xf1\x09\x98\xb7\xe7\xf0\x67\xfb\x76\x78\xd8\x0d\xd4\x29\xdf\x15\x3f\x3b\x21\x3a\xe5\xca\x67\x24\xe5\xea\xd7\x08\x0e\x2d\x95\x43\x88\xf4\x4b\xd4\x15\x50\xb5\x02\x57\x9a\x29\x57\xd5\xd6\x77\x13\x9b\x7e\xbe\x62\xc1\x9d\x8c\x85\x24\x09\x16\xca\xf6\x7a\xf6\xc1\x72\xec\xf5\xda\xac\xd8\x48\xee\xf7\x46\x8c\x2d\xb8\xcb\xb7\xa1\x0a\x96\xa5\x50\x19\x0a\x30\x10\xf1\xc8\xd8\xa2\x84\x1b\x9a\xe7\x60\x99\x31\x81\xcb\x6b\x33\xdd\x50\x95\x69\x2f\x4a\xf5\x2f\x92\xc3\x22\xe7\xc9\xb5\x84\xaa\x1d\x0b\x94\x55\x4d\xad\x82\x08\xb2\x46\xe5\x98\x52\xd3\xce\xfa\x68\x9a\x50\x13\xdb\x7c\x8a\x2b\x43\xb5\x60\xb6\x68\xfb\xe9\xe3\xfb\x5f\xde\xbc\x7b\xff\x1e\x3e\xbd\xfe\xf9\x1f\xdf\x7d\x7a\xfd\x0a\xde\x7c\xfc\x04\x1f\xdf\xbf\x7a\xfd\x09\x5e\x7e\xfa\xf8\xe5\xe2\xf5\xa7\x8b\x36\x89\x37\x4b\x36\xfe\xb1\xc2\x75\x19\x55\xfe\x2c\xba\xf2\x33\x58\x93\x3a\x25\xf8\x52\x20\xb9\x36\x3d\xba\x36\x70\x9b\x6e\xc7\x86\x08\x6a\x4a\x71\x9d\x0a\x49\x20\xb6\x95\xa7\x38\xdc\x64\x68\xf6\xc9\xcc\xa2\xb6\x9a\xd4\xfa\x4e\x14\xdd\x98\xc0\xee\x66\xf3\xba\xc6\x61\x55\xac\x6f\x79\xf3\xe4\xa1\xd3\x75\x57\x1c\x5f\x4b\x94\x36\xf3\xfe\xfc\xd7\xf7\xc6\x59\x9c\xf8\x89\xbf\x05\x08\x94\xd7\x7f\xaf\x19\x48\x32\xce\x9d\x96\x88\x86\xd0\x13\x27\x6d\x1d\xd5\xf4\xcb\xbb\xd1\xa7\x9d\x37\xd1\x19\x92\x90\xd8\xcd\x8d\x8c\x51\x99\x9d\xb3\x56\x45\xe1\xb9\x0b\x65\x73\x8a\x13\xa0\xbe\x7d\xd5\x4f\xec\xe4\x77\xb9\x16\x72\xc9\xd4\x50\x04\xf4\x77\xe1\x70\x3e\xe8\xbb\x77\x44\xe9\x08\x8c\x3b\x9c\x47\x4d\xab\xff\xc8\xb4\xdd\x67\x70\x00\x87\x83\x08\xfd\x73\x81\x29\xfc\x00\x4f\x2d\xab\x3b\x60\x0e\xd6\x32\x02\xce\xce\x73\x9a\x5c\xcf\xa3\x5d\xd8\x25\xe6\xcb\x1d\xc3\xd1\xc4\x2a\x5a\x7c\x10\xed\x98\x85\x79\x35\x6d\x17\xa6\x83\x31\x7c\x0f\x49\xbd\xc9\x3b\xd1\x8c\x4e\x8c\xe2\xfa\x89\xd7\x81\x93\x00\x1d\x8c\xfe\x29\x3a\xfd\x1b\x50\xaa\x93\xa6\xde\x8c\xfb\x91\xa3\xc2\x97\xf4\x6a\x0c\x74\x34\x14\x40\x9b\x37\x2e\x52\x14\x36\xbd\x93\x16\x5c\x92\xad\x1c\xef\x48\xbe\xfa\x71\xc6\x2c\xd7\x76\x8d\xe2\xa8\x0e\x2f\x1d\x93\xf0\x15\xd5\xf4\x29\xfa\x2b\x38\x73\x72\x4d\x0f\x60\x17\x3f\x66\x3d\x63\xa8\xf3\xeb\xde\xd8\x6c\x88\x63\x93\x19\x37\xec\x3a\xc9\x85\x1b\xe5\xa7\x53\x60\x64\x43\x57\x3a\x68\x44\x56\x9b\x64\xd4\xfa\x29\x52\x1f\x5d\xb4\xfc\x5d\xe3\xd6\x78\x66\x85\xbe\xd9\x6a\x7f\xa0\x67\x75\x27\xf7\xdc\x8b\xcd\x09\x6f\x28\x4b\xf9\xcd\xe5\x35\x6e\xaf\x3a\x49\x21\x9c\xb9\x83\x71\x93\x57\xb4\x2b\x38\xe9\x04\x10\x8b\xc5\x73\x6a\x7b\xfe\xbe\xc7\xfd\x88\xa0\x61\x46\x63\x88\x7b\xa7\x34\x9a\x67\x7f\x9b\x67\x66\x45\x66\x31\x7e\xa7\x88\xa4\x29\x08\x3c\x5a\xe9\x12\x96\xa8\xba\x7d\xa5\xad\xc1\x1c\x49\xe9\x98\x90\x55\xbd\x5e\xa9\xdc\x04\x59\xfb\xe3\x41\xde\x75\xb1\x55\x6d\xf0\x6e\x4f\xdd\xed\x17\xf5\x0b\x98\x5d\xc5\x8b\xd3\x41\x33\xb5\x4b\x4d\x72\x67\xe1\xb2\xa3\xfb\x06\x5e\x07\xee\xbe\xdb\x28\x22\x05\xb5\xa7\x45\xbd\x7c\xd0\x26\x28\x19\x65\xc3\xa7\x37\x8f\x33\xca\x4a\x41\x5a\xbc\x17\x8a\x17\x8f\xc5\x1d\x30\x9f\xfb\x66\x31\x91\xd6\xe2\x08\x90\x24\x19\x54\x99\x85\x49\x83\x28\x33\x42\x5f\x09\x5e\x16\x4d\x7e\xe1\xfa\x20\x8f\x8b\xaf\xe3\xa6\xd4\xe9\x5a\x92\x3d\x13\x19\x62\x19\x4c\x08\x51\x75\xef\x61\xb7\xf3\xa8\x49\xc4\x1d\x6f\x39\xf6\x63\xdc\x28\x64\x5f\x9a\x95\x82\x4b\x93\xd9\x99\x58\xdf\x8c\x98\x1c\xc0\x43\xa8\xe7\x32\xbc\x6d\x4a\x02\x98\xb7\xa0\x87\xf0\xd5\x69\x35\xf8\xae\xd3\x83\x39\x6d\x81\x86\x67\x1d\x1d\x35\x43\xdd\x05\xc7\xfd\x76\xdc\x18\x28\x4b\xf1\x36\xb4\x35\xdf\xdf\xe5\xfa\x71\x1c\xf7\xa5\x41\x77\xe5\xc8\xae\xef\xaf\xfd\x9d\x1e\xc5\x86\xd0\xd8\x5b\x44\xd7\x54\x1a\x7e\xe2\x7e\xdf\xd8\x26\xf5\x24\x5c\x44\xd9\xc6\xd4\xdc\x2d\x39\xf5\x60\xc3\x71\xff\x00\xa7\xe0\x52\x61\x5a\x75\xab\xc7\x90\x95\x6b\xc2\xc6\x90\xd3\x8d\xd7\x6f\xab\x73\xd4\xca\x61\x3b\x1d\x78\xef\xaa\x85\xfd\x3e\xf3\x12\x54\x8d\xaa\x85\x34\xbf\xfa\x80\xe6\xf3\xcc\x74\x28\x4c\xbe\xdd\x5c\x49\xc8\x88\x02\x73\x3a\xc5\x15\x44\x7a\x56\xd4\xfc\xac\xee\x5f\xd8\xf3\xfd\xea\x00\x1e\xf5\xdb\x72\x89\x02\x99\x72\x4e\x2d\xa6\x53\xb3\x0b\x9d\x6b\x1f\x86\xe8\x59\x1d\xdc\xdc\xb1\x19\x1c\x7b\x6a\xef\x8d\xfa\xb9\xb7\xd9\xb0\xae\xa6\x74\x29\xb5\xd8\xee\x01\x73\x89\xdf\x99\xfe\x30\xc6\xee\x3d\x16\xdd\x4b\x26\x2e\x7c\xfb\x7d\xe6\x9a\x6c\xab\x49\x76\xab\x01\x6d\x61\xeb\xe1\x5c\xf4\xda\x8e\xbb\x8b\xe0\x16\xe6\xbc\x8a\x1f\xc3\xa0\x3a\xa8\x44\xbd\x8e\x45\xa8\x98\xaf\xab\x78\xd0\xe5\x74\xfc\xa4\x5a\x5c\xd4\x44\xd3\xc8\x28\xdc\x08\x0e\xad\xfa\x75\x51\x56\xac\x04\xdb\x04\x55\x98\xeb\xd5\xe5\x35\x0c\x65\x0c\x45\x55\x74\x49\xd2\x43\x1d\x28\xc9\x2b\xc8\x70\x2b\x86\x32\x89\x42\xbd\xc4\x25\x17\x18\xd7\x37\x86\x3a\xa5\xbc\x23\xa0\x2a\xe1\xd2\x91\x15\x85\x2d\x5c\xfc\xa3\x51\xc7\x0c\xa3\xae\x0a\xb6\xb2\x78\x69\x4f\x26\x61\xde\x59\xdc\x60\x0b\xb9\x39\x2f\x1c\xa8\xf4\xbe\xd9\x4a\xef\x1b\x3c\xf7\x29\x34\xc5\xde\xb7\x07\x14\x7b\x43\xce\xd5\xd4\xb3\x98\x4f\x0a\xa2\x6d\xf8\x03\x4f\xd1\x7d\x7d\x48\xc7\x38\xdc\x86\x03\x5b\x75\xd8\x6e\x98\xbd\xfa\x34\xd7\x1f\xf8\x72\x29\x51\x7d\x31\x1f\x8e\xbc\x0b\x4e\x4f\xe1\x07\x38\x6c\xee\x56\x1d\x42\x54\xdc\x06\x0b\x11\x6f\x0f\x2e\xbf\x5d\x0d\x15\x22\xae\xde\xf4\x8e\x3c\x12\x5d\x27\x46\xdf\xcb\xfd\x43\xc5\x76\x58\x04\x03\xf5\xf6\x83\xc5\x10\xdc\xae\xe3\x6a\x17\xfa\xbd\x4c\x07\x24\x24\x22\xb7\x37\x3e\xd2\x29\x78\x0f\xb5\x56\xe0\xc1\x4e\xa9\x83\xdd\x66\x79\xce\xd9\x14\x67\x76\xeb\xc2\x52\xef\x89\x87\x5e\x7d\x2f\x42\x43\x9d\x2d\x86\xcf\x15\x8c\x35\xa6\x1d\x72\xf7\xa1\x00\x3e\x9d\x82\xa2\x6b\x1c\x83\xa4\xdf\x10\xf6\xeb\x78\xd5\x8c\xdf\x10\x73\x9e\x5b\xf9\x33\xdf\xcf\x6b\xdd\xf3\xfc\xf5\x9a\x32\xeb\x22\xbe\xec\x84\xfa\x8b\x0f\xa5\x8d\x49\x92\xe6\xbe\xc4\x53\x27\x4a\x9c\x76\x60\xcd\x5d\x17\x13\xe4\x43\xde\xc4\xf0\x7a\x38\x77\xe0\x9f\x42\x8b\xf8\xa4\x3f\xf7\x79\x87\xe3\x33\x88\xab\xf5\x7a\xdf\x07\x84\xfe\xe0\x7c\xf6\xa1\xb5\xc0\x43\xb3\x5d\x3f\x6b\xd9\xdf\x7f\x18\x27\x8f\x2c\xba\xc7\x76\x8b\x8e\xfa\x32\xef\x85\x8d\x3a\x9f\x7b\x04\x1b\x0f\xef\x53\x38\x07\xed\xda\x0a\xcf\x49\x9e\xf4\x82\x84\xeb\x1e\xbb\xce\xf0\x24\x44\xdd\x37\x6b\x47\xa5\xe0\xac\x47\xcc\x4f\x14\x07\x91\xf8\x77\x5a\x9e\xd3\xf5\x0a\xa4\x48\xe6\x01\x93\x3f\x83\xe8\xc7\xe3\xbf\xeb\x7b\xe9\x59\x00\x71\x90\x70\x47\x78\x81\xb3\x54\x9b\x2f\xc7\x6e\xf1\x0f\xee\x55\x11\xd3\xfa\x6d\x4a\xf1\xdd\x25\xba\x73\x83\xca\x27\xf2\xc4\x69\x43\xef\xef\x0f\xdc\x05\xa9\xae\x30\x39\x57\xb6\xba\x1c\xd5\x37\x94\x65\x22\xb8\xb9\x46\xe9\x4d\x68\x4e\x93\xde\x22\x5d\x65\xda\x34\xdb\xf3\x25\x2b\x76\x3b\xd0\x31\x4d\x83\xeb\x55\x95\x64\x27\xd8\x68\x8c\x3d\x6c\x32\xa3\x15\xc2\xa3\xfe\xc8\x67\x5e\x84\xd0\xbd\x35\x25\x79\x0f\xf5\xd4\x5c\xc0\xee\xe9\xac\x9d\x17\x2a\xa2\xbf\x1b\x33\x9f\x99\x98\x19\x24\xd5\x70\xf2\xd0\xe0\xb9\x2b\x74\x3e\xd8\x81\xb9\xcf\xe3\xf6\x11\x4e\xbb\x12\xdc\x89\x1c\x6c\xef\x2b\x88\x29\x20\x42\xfd\xf9\xb0\xdd\x93\xe1\x9c\xaa\x7e\xba\x8d\xb9\xee\x73\x3f\x06\x0a\x4f\xe1\xc7\xc1\x38\xfd\xdd\xca\x79\x30\x44\x78\xfa\x30\xee\x77\xeb\xfb\xae\xf7\x30\xe0\x7a\xbb\x85\xb6\xbd\x5a\x08\x85\xc0\x0d\xe5\xa5\xf4\x8c\xd0\xcc\x09\x25\x63\xc3\x97\x15\x7b\x5a\x35\x70\xd5\x20\x80\xe1\x92\x5e\x4d\xb4\xab\xf0\xb3\xf1\x41\xb5\xea\xdc\x14\x18\x40\x68\xeb\x7d\x1f\xa5\x16\x60\xe0\x53\xdb\x60\xa9\x3a\x5f\xf7\x7b\xff\x17\x00\x00\xff\xff\x24\x30\x47\x25\xcd\x31\x00\x00")

func componentHomoJsBytes() ([]byte, error) {
	return bindataRead(
		_componentHomoJs,
		"component/homo.js",
	)
}

func componentHomoJs() (*asset, error) {
	bytes, err := componentHomoJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/homo.js", size: 12749, mode: os.FileMode(420), modTime: time.Unix(1554735280, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _componentStylesInputCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xc1\x6e\x9c\x30\x10\x86\xcf\xf0\x14\xa3\x44\x95\xd2\x68\x4d\xcc\xaa\x7b\xa8\xf7\x05\x7a\xea\x3b\x8c\xf1\x00\x56\x8d\x6d\xd9\x43\x17\x5a\xf5\xdd\x2b\x93\x40\xa2\xe4\xd0\x4b\x35\x12\x32\xa3\x9f\x99\x8f\xff\xf7\xd3\x23\x64\x5e\x1d\xc1\x9c\x29\x81\xf5\x71\x66\xe8\x2d\x39\x03\x8f\x4f\x75\xa3\x67\xad\x1d\x89\x2e\x78\x46\xeb\x29\x41\xb3\x29\xc4\x2d\x61\x84\xdf\x75\x15\x43\xb6\x6c\x83\x57\x80\x3a\x07\x37\x33\x5d\xeb\x4a\x07\xe6\x30\x29\x90\xd7\xba\x72\xd4\xf3\xf3\x29\xd9\x61\x7c\x39\xf6\xc1\xb3\xe8\x71\xb2\x6e\x55\x70\xf7\x8d\xdc\x4f\x62\xdb\x21\x7c\xa7\x99\xee\x4e\x70\x34\x4e\x90\xd1\x67\x91\x29\xd9\xfe\x5a\x57\x5d\x70\x21\x29\xb8\x3f\x77\xa5\xae\xf5\x9f\x7f\xf0\x31\x2d\x8c\x89\xb0\x80\xde\xac\xe1\x51\x41\x87\xae\x7b\x68\xa5\xfc\x04\x02\xce\x32\x2e\x9f\xff\x1f\x4d\xa5\xb1\xfb\x31\xa4\x30\x7b\xa3\x20\x0d\x1a\x1f\xce\x17\x79\x82\xd7\x87\x6c\xbe\x5e\x8e\x7d\xd9\xfe\x22\x05\x2d\x4d\x9b\x47\xcc\x94\x44\x8e\xd8\x59\x3f\x28\x68\x2e\x71\xd9\x75\x37\x7a\xb6\xed\x8b\x2c\xc6\x4d\x98\x06\xeb\x15\xb4\x72\x53\xe8\x90\x0c\x25\x91\xd0\xd8\x39\x2b\x68\x2f\x6f\xba\x0a\x7c\xf0\x25\x8d\x88\xc6\x6c\x63\xcb\x47\xbb\x26\xcc\xec\xac\xa7\x43\xa4\xc3\x22\xf2\x88\x26\xdc\x14\xc8\xad\xda\xb8\xc0\xbd\x91\xa5\xc0\xfa\x4c\x5c\x40\xad\x27\x31\xbe\x10\xb5\xcd\xf9\x52\xf0\x8f\x14\x9a\x44\xd1\xad\xa2\x4f\x44\x7d\x48\x53\x71\x7d\xe7\x95\x1f\x64\xef\xc4\x4d\xc6\x15\xde\xa6\x49\x9e\x8f\x77\x3d\x33\x07\xff\x3a\x4f\x70\x88\xaa\x00\x5e\xeb\xaa\x64\x2c\xd0\xd9\xc1\x2b\x28\x77\xed\xe3\xa2\x8c\xeb\x3e\x68\xb4\xc6\x90\x7f\x07\x56\x71\x42\x9f\x0b\x84\x82\xdc\xa1\xa3\x07\x59\x52\xda\x7f\x73\x43\xff\x1b\x00\x00\xff\xff\xb7\x50\x66\xcd\x25\x03\x00\x00")

func componentStylesInputCssBytes() ([]byte, error) {
	return bindataRead(
		_componentStylesInputCss,
		"component/styles/input.css",
	)
}

func componentStylesInputCss() (*asset, error) {
	bytes, err := componentStylesInputCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/styles/input.css", size: 805, mode: os.FileMode(420), modTime: time.Unix(1554450500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _componentStylesReplyCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xcd\x6e\xe3\x36\x10\x3e\xd7\x4f\x31\xdd\x60\x81\xd8\x35\x6d\xc9\x8e\xd3\x84\xbe\x14\x28\xd0\x6b\x0f\x7d\x02\x4a\x1c\xd9\x84\x29\x8e\x30\xa4\x12\x67\x8b\xbc\x7b\x41\x4a\x8a\xed\xd8\x9b\x1f\x74\x17\xb0\x05\x89\xc3\xf9\xfb\xf8\xcd\x0c\xe7\x13\xf0\xe1\xc9\x22\xb4\x1e\x19\x18\x7d\x43\xce\x23\x30\x36\xf6\x09\x26\xf3\xd1\xac\x68\x8b\xc2\xe2\xac\x5b\xf8\x77\x04\x50\xa8\x72\xb7\x61\x6a\x9d\x96\x10\x58\x39\xdf\x28\x46\x17\xd6\x51\x44\x7b\xe1\xb7\x4a\xd3\xa3\x04\x47\x0e\xe3\x5a\x65\x49\x05\x09\x6c\x36\xdb\xb4\xa7\x21\x6f\x82\x21\x27\x81\xd1\xaa\x60\x1e\xd2\xae\x64\xa8\x22\xae\x05\xb1\xd9\x18\xd7\x2b\x40\xa0\x26\x8a\x6b\xc5\x69\xf1\xae\xd9\x43\x06\x79\xd6\xec\x93\x29\xa5\xb5\x71\x1b\x09\x59\xb7\x67\x2f\x1e\x8d\x0e\x5b\x09\xb7\xab\xaf\xeb\xd1\xf3\x69\xec\xb3\xad\xf1\x81\xb8\xcb\x61\xb0\x97\x41\x06\x8b\x68\x73\x0d\xf3\x09\x30\xd6\x58\x17\xc8\xa8\xa1\xd3\xf4\xa0\x09\x1c\x05\x70\x88\x1a\x02\x81\x0f\xca\x69\xa0\x36\xc0\x83\x51\xbd\x95\x88\xd2\x6b\x5f\x5e\x75\x7e\xe6\x93\xe8\xcc\xb8\x21\xb0\xe5\xaa\x0b\xfd\x5c\x05\xfa\x2f\x51\x92\x0b\xe8\x42\x52\x4f\xa8\xf4\x68\x29\x6b\x61\x91\x65\xb5\x3f\xcb\xec\x4c\x75\xf8\x2e\xda\x10\xc8\x9d\x1d\x1a\x6f\x0a\x75\x7d\x73\x33\x85\xe1\x9f\xcd\x6e\x7f\x1f\x47\x08\x4b\xb2\xc4\x12\xae\xaa\xaa\x3a\xc1\x37\xc2\x9e\xdf\x76\xa0\x17\xc4\x1a\x59\xb0\xd2\xa6\xf5\x12\xf2\x55\x94\xc5\x47\xff\x3f\x1c\x97\xb0\x58\x05\x19\xf1\x4d\x27\x8c\xfb\x20\x94\x35\x1b\x27\xa1\x44\x17\x90\xe3\xaa\x36\xbe\xb1\xea\x49\x82\x71\xd6\x38\x14\x85\xa5\x72\x77\x89\x36\x65\xcb\x3e\xc6\xd6\x90\x19\x74\xbf\x83\x4e\xef\x4a\x63\x49\xac\x3a\xe9\xc0\xc5\x47\x62\x2d\x0a\x46\xb5\x8b\x6b\x5c\x2b\xbb\x1e\xfd\x92\x48\x6b\xbe\xa5\x44\x7b\x04\x45\x41\x29\xe6\xf9\x04\x94\x33\x75\xb2\x22\x74\x3b\x98\xcb\xfd\x3a\x9e\x20\x1c\x09\x9d\xaa\x51\xc2\x97\x6e\x01\x45\x3a\x97\x2f\xeb\x93\x2d\x31\x4f\xe1\x83\x0a\x28\xa1\x51\xad\x47\x7d\x2a\xaf\x8c\xb5\xa2\x26\x8d\x12\x2a\xe2\x47\xc5\xda\xf7\x31\x50\xa3\x4a\x13\x9e\x22\xcf\x3b\xbf\x2f\xd5\xd2\x57\xa0\x55\x01\x97\xfa\x3a\x6b\xf6\x53\x18\x1e\xe3\x53\xeb\x1a\x13\xcc\x62\x99\x8c\x8a\xda\x8b\x0b\xb2\x24\x7a\xc4\x62\x67\xc2\x25\x71\xa2\xde\x1f\x3b\x7c\xaa\x58\xd5\xe8\xe1\x24\xdb\x44\xb3\x8a\xa9\x4e\x2f\x70\x1c\xf4\x08\xe0\x39\x46\x4d\xaf\x45\x79\x27\xba\x54\x3d\x1f\xe0\xf4\x65\x68\xb9\x75\xce\xb8\xcd\x11\x0d\x03\x35\x12\x96\x3d\x33\x8d\x13\x5b\x8c\xa4\x92\xb0\xb8\xe9\xd6\xe8\x01\xb9\xb2\xb1\x67\x6d\x8d\xd6\xe8\x62\x96\x9f\xaa\x30\x59\x19\xf6\x41\x94\x5b\x63\x75\x57\x6d\xdf\xab\x91\xfc\xcd\x22\xf9\x5c\x5d\x4b\xab\x06\xa7\xd3\x4f\x29\x0e\x5f\x8d\x29\x77\x6f\x87\x3b\xbc\x9c\x37\xd3\x23\x13\xa8\x3f\x72\x58\x47\x94\xf5\xa5\xb2\x78\x9d\x8d\xdf\xe1\xee\x71\x77\x7f\xe5\x5e\x3a\x0a\xd7\xa7\x31\x8c\xdf\x43\x6b\x1b\x8f\xf9\xff\x02\xf5\x56\x17\xcd\xc7\xeb\x4b\x79\xe6\x1f\xce\xf3\xb8\xcb\x0e\x1c\x55\x6d\xa0\xc4\xc7\xf9\x04\x52\xe3\x53\x65\x64\x3c\x30\x96\xb1\xe5\xa5\xd1\xed\x8f\xe7\xf4\x30\xe5\xde\xce\x6c\xfa\x7a\xff\x0f\xc3\xf5\xc4\xde\x0f\xc5\xf7\x30\xa5\xfa\x51\xa0\xb1\x52\xad\x0d\x07\x78\x9a\x36\x40\x65\xd0\x6a\x1f\x3b\xe8\xcb\x0c\x3f\xa0\x73\x16\x4f\xa7\xf3\xda\x75\x9c\x44\x8a\xc5\x26\x16\x04\xba\x70\x9d\xdf\x2f\x35\x6e\xa6\x70\x95\x57\xaa\x44\x3d\x85\xab\xd5\xea\x7e\xa1\x4b\xc8\xb3\xec\xeb\x18\x7e\x35\x75\x43\x1c\xd4\xf9\x1d\x28\x8b\x87\x0c\x79\xbc\x60\x34\x7b\xb8\xca\xb2\x6c\x0a\x19\x88\xb4\x90\x96\x52\x92\x8b\xd5\x6a\x0a\x87\x47\x36\x5b\xde\x8d\xc1\x38\x8f\xe1\x65\x98\x1d\x0c\xe6\x83\x56\x34\xd5\xfd\x66\xcb\xd5\xf8\x30\x99\xfb\x86\x4b\x6d\x88\x69\x9c\x56\xcf\xe5\xfc\xa5\x1c\x7a\x7e\xfa\x8c\xdd\xb4\xc4\x2d\x59\x8d\xdc\x5f\x62\xe0\xcf\x2d\x53\x8d\xf3\xbf\x1b\x64\x35\xff\x47\x55\x8a\x4d\x37\x89\xfa\x1b\xc3\xa5\x3c\x66\x5d\x54\x27\xf1\x77\x83\xf8\xfd\x78\x6a\xfa\x76\x29\x8c\xbf\x0c\x63\x45\x7b\xc8\xef\x7f\xfb\x89\xee\x19\x95\x16\xe4\x2e\x5c\x77\xcf\x88\x71\x7b\xdb\x11\xe3\xe6\x2e\x5f\xe4\x91\x18\x99\xbe\xc9\x56\x77\x17\x88\xf1\x3c\xfa\x2f\x00\x00\xff\xff\x70\x53\x5a\x14\x6a\x0b\x00\x00")

func componentStylesReplyCssBytes() ([]byte, error) {
	return bindataRead(
		_componentStylesReplyCss,
		"component/styles/reply.css",
	)
}

func componentStylesReplyCss() (*asset, error) {
	bytes, err := componentStylesReplyCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/styles/reply.css", size: 2922, mode: os.FileMode(420), modTime: time.Unix(1554450500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _componentStylesSaysCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x6d\x6e\xdb\x3a\x10\xfc\xef\x53\xec\xc3\x43\x50\x3b\x35\x15\xda\x8d\x82\x80\x39\x0d\x25\xd2\xd2\x22\x14\x57\x20\xd7\xb5\xd4\x22\x77\x2f\x48\x59\xb2\x9d\xa4\x68\x0b\x03\xb2\x34\xdc\xaf\xd9\x19\x3e\xdc\x43\xe4\xd1\x59\xa8\x8e\x55\xe5\x6c\x84\xfb\x87\x55\x31\xbd\x6f\xe7\x17\xc1\x63\x8f\xbe\x81\x9f\x2b\x80\x9a\x1c\x05\x05\xff\xef\x77\xe9\xf7\xb2\x02\xa8\x74\xfd\xda\x04\x3a\x7a\xa3\x20\x34\x95\x5e\xef\xcb\x72\x0b\x97\x87\x2c\x9e\x1f\x37\x29\xb0\xd7\xc6\xa0\x6f\x14\x3c\xf7\x03\xec\x9e\xfa\x21\x67\x53\x30\x36\x88\xa0\x0d\x1e\xa3\x82\x32\x1d\xdd\x3c\x52\xd0\x81\x3c\x8b\x93\xc5\xa6\x65\x05\x8f\x52\x26\x8c\xed\xc0\x82\x83\xf6\xf1\x40\xa1\x53\xe0\xc9\xdb\x05\xd7\x0e\x1b\xaf\xc0\xd9\x03\x2f\xf9\x11\x7f\x58\xb5\xf4\x75\x96\xd9\x06\x11\x7b\x5d\xe7\x99\x8a\x73\xab\x4e\x87\x06\xbd\x02\x09\x12\xf6\xfd\x00\x72\x02\x07\x71\x42\xc3\xad\x82\xa7\xf2\x2e\x57\x74\xa4\xf9\xd2\xb4\x76\x56\x07\x05\x15\x71\x9b\x8b\xa3\xb7\xa2\x3d\xcf\xbb\x2b\x4a\xdb\x25\xf4\x44\xc1\x88\x2a\x58\xfd\xaa\x20\xff\x89\x84\xe4\x99\x67\x1a\x82\x02\xe6\xf6\x69\x72\x60\xea\x97\x53\x64\x24\xaf\x40\x3b\x07\x7b\x29\xbb\x38\xed\x6e\x48\xac\xf2\xfc\x35\x79\xb6\x9e\x45\x45\xc3\xcb\xea\x6d\x56\x0e\x66\x05\xcf\xc7\x59\xc2\xeb\x7a\x94\xf8\xf3\x08\xbb\x32\xd7\x5c\x12\x95\x27\x5e\x17\x51\x8f\x9b\x4f\x4b\x9c\xd3\x54\x5a\xcf\xdb\x3b\x9b\x14\xd8\xe9\x06\xfd\xc5\x3e\x33\x70\x69\x3e\x29\x16\x6b\xed\xec\x5a\x6e\x7e\xcb\x71\x0b\xd3\x0e\xa7\x2f\xd8\xc5\xed\xec\xa1\x05\xb9\x6a\x7f\xd3\x66\x91\x31\x15\x9f\x95\x90\x37\x2e\xcc\xa3\xaf\x16\xff\x77\xd6\xa0\x06\x6e\x35\x7f\x89\x80\x3e\xa2\xf9\xec\x4e\x7c\xd8\x06\x76\xd3\xbd\x38\xdb\xa3\xd6\xae\x5e\xef\xa4\xbc\x83\xaf\xf0\x6d\xdf\x0f\x9b\x6b\x4f\x89\x64\x7c\x31\x3b\x90\xbe\xdb\x70\x70\x74\x52\xd0\xa2\x31\xd6\x27\xcc\x60\xec\x9d\x1e\x15\x54\x8e\xea\xd7\xbf\xbc\x1f\x13\x0b\xf4\x6c\x83\xae\xd3\x0a\x21\xd8\x3a\x2d\x31\x13\xbb\x9e\xbe\x68\x31\x32\x85\x71\xfb\x1e\x78\x4f\xeb\x8f\x01\xcb\x77\x75\x64\x26\xff\xaf\xf1\xaa\x4d\xec\x3f\xb8\x31\x0d\x9d\x74\xfd\x0f\xbb\x9e\x02\x6b\xcf\xd7\xfa\xce\xa5\x6f\x0c\x58\xec\xcb\x14\xf3\x2b\x00\x00\xff\xff\xb8\xb0\xfc\xbb\xc7\x04\x00\x00")

func componentStylesSaysCssBytes() ([]byte, error) {
	return bindataRead(
		_componentStylesSaysCss,
		"component/styles/says.css",
	)
}

func componentStylesSaysCss() (*asset, error) {
	bytes, err := componentStylesSaysCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/styles/says.css", size: 1223, mode: os.FileMode(420), modTime: time.Unix(1554450500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _componentStylesSetupCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x51\x6e\xdb\x30\x0c\x86\x9f\xad\x53\x10\x1d\x06\xb4\x59\x9c\xc8\xcb\x82\x02\xf2\x05\xf6\xb4\x3b\xc8\x12\x6d\x0b\x95\x45\x41\xa2\xe3\x64\x43\xef\x3e\xc8\x6e\xd3\x00\x7e\xa0\xff\x9f\x22\xbe\x9f\x3c\xee\x20\x23\xcf\x11\x0c\x05\xd6\x2e\x60\x82\xcc\x37\x8f\x19\x76\x47\x71\xe8\xe6\xae\xf3\x58\x7f\x79\xff\x44\xd5\x69\xf3\x36\x24\x9a\x83\x55\xf0\xcd\x1a\x6b\x51\xb6\xa2\x1a\xd1\x0d\x23\x2b\x38\xff\x94\xf1\xda\x8a\x6a\xd2\xd7\x7a\x71\x96\x47\x05\xaf\xe7\x4d\xfa\xf8\x6d\xa4\xfc\xbe\x36\xa4\xc1\x05\x05\x12\xf4\xcc\xd4\x8a\x8a\x2e\x98\x7a\x4f\x8b\x82\xd1\x59\x8b\xa1\x15\x55\xa4\xec\xd8\x51\x50\x90\xd0\x6b\x76\x17\x6c\xc5\xfb\x9d\x6a\x49\x3a\x16\xa0\xaf\x2e\xdd\x65\xf2\x33\x63\x2b\x2a\xa6\xa8\xa0\x80\x75\xc4\x4c\xd3\x56\x7b\xec\x79\xab\xd2\x46\x5b\x37\xaf\x2b\x5a\xd4\xd6\xba\x30\x14\xb8\x78\x05\xa3\xbd\x79\x2e\x0e\xfc\x58\x85\x17\x38\x15\xb9\xd9\x62\x7c\x72\xd6\x37\x05\xd9\x24\xf2\xbe\x15\x55\xbd\x60\xf7\xe6\xb8\xbe\x9b\x9b\xb3\xce\x64\x9a\xcd\xf8\xd0\xc3\x49\x87\xdc\x53\x9a\x14\xac\xa5\xd7\x8c\x27\xfb\x2c\xf7\x50\xbe\x97\x12\xf1\xb8\x03\x8a\x25\x93\xf6\x10\xf5\x80\x0f\x37\x19\x9b\x92\x99\xf1\xca\xb5\xf6\x6e\x08\x0a\x0c\x06\xc6\xd4\x8a\xaa\xa7\xc0\xf5\xf2\x71\x88\x93\x94\x9f\x52\x76\x7f\x51\xc1\x2f\x9c\x1e\xd6\x7e\x38\xe3\xb4\x6e\x1e\xe4\xa1\x39\x17\xeb\x5d\x74\x64\x6f\x65\xf8\xfa\xa8\xd7\x93\xf3\x37\x05\x4f\xbf\xd1\x5f\x90\x9d\xd1\xf0\x07\x67\x7c\xda\xc3\x5d\xd8\x43\xd6\x21\xd7\x19\x93\xeb\x1f\x2f\x5a\x66\xfd\x0f\x00\x00\xff\xff\xa0\xad\x2a\xf7\x58\x02\x00\x00")

func componentStylesSetupCssBytes() ([]byte, error) {
	return bindataRead(
		_componentStylesSetupCss,
		"component/styles/setup.css",
	)
}

func componentStylesSetupCss() (*asset, error) {
	bytes, err := componentStylesSetupCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/styles/setup.css", size: 600, mode: os.FileMode(420), modTime: time.Unix(1554450500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _componentStylesTypingCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xd1\x8a\xab\x30\x10\x86\xaf\xcd\x53\x0c\x85\x03\xa7\xa5\xda\x68\xeb\x39\x4b\xbc\xd9\x37\x29\x89\x19\x6d\x68\x4c\x24\x8e\x6c\xa5\xf4\xdd\x97\x68\x61\x2d\xdb\x8b\x90\x61\xfe\xef\xff\xc9\x64\x0e\x3b\x18\x68\xb2\x08\x1b\xeb\xa5\x36\xae\xdd\x80\x0f\xb0\xa1\xa9\x9f\xeb\x81\x24\xc2\xee\xc0\x58\xa6\x46\xa5\x2c\xa6\x8b\x00\x77\x96\x7c\x19\x4d\x17\x01\xc7\x8f\xfe\x56\xb1\xa4\x97\x3a\xba\x05\xe4\x45\x7f\x83\xfc\xdf\xdc\xbc\xa0\x69\x2f\x24\x60\x46\x1e\x2c\xd3\x9e\xa2\x53\xc9\xfa\xda\x06\x3f\x3a\x9d\xd6\xde\xfa\x20\x20\xb4\xea\x6f\x51\x96\xfb\xe7\xd9\x56\x2c\x69\xac\x97\x24\xc0\x62\x43\xab\xa4\xff\x73\x6e\x27\x43\x6b\x5c\x1a\x35\x01\xa7\xb9\xf5\x7c\xcd\xa2\x4b\x67\x3a\x49\xc6\xbb\xd4\xc9\x0e\x05\x28\x3f\xba\x1a\xcf\xda\xd3\x8b\xa8\xc7\x30\x17\x02\x8a\xac\x38\x0d\x2f\x9a\x21\x5c\xc4\xb4\xf6\xa3\x23\x01\xc6\x35\xc6\x19\xc2\xd7\x04\x13\xb0\x5e\x22\x9c\x0f\x9d\xb4\x15\x4b\x94\x0f\x1a\x43\x1a\xa4\x36\xe3\x20\xa0\xfc\x99\xfd\x9c\xc3\x1d\x56\x66\xb4\x72\x12\xc0\xb3\x53\x39\x54\xf0\x44\x8a\x77\x48\x9e\xf1\x15\x72\x7c\x8f\x1c\x17\xe4\xf3\x8a\x53\x13\x64\x87\xc3\x6a\xec\xf8\xeb\xfc\x0f\xdc\x1f\x2c\x29\xe3\x0d\xbf\x56\x10\x37\xc0\xf7\x7c\xcf\xb7\x31\x24\xc9\xf9\x82\x3f\xd8\x77\x00\x00\x00\xff\xff\x55\x2a\x91\x28\x21\x02\x00\x00")

func componentStylesTypingCssBytes() ([]byte, error) {
	return bindataRead(
		_componentStylesTypingCss,
		"component/styles/typing.css",
	)
}

func componentStylesTypingCss() (*asset, error) {
	bytes, err := componentStylesTypingCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component/styles/typing.css", size: 545, mode: os.FileMode(420), modTime: time.Unix(1554450500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6f\xd4\x30\x10\xbd\xef\xaf\x98\x1a\x21\xed\x4a\xcd\x47\xc5\x05\x2d\xc9\x1e\x28\xad\x80\x0b\x48\x14\x21\x8e\x5e\x7b\x36\x31\x75\xc6\x91\x3d\xc9\x36\xa0\xfe\x77\xe4\x98\x6e\x43\x4f\x48\x34\x97\xd8\xe3\x99\x37\x6f\x3e\x5e\x75\xf6\xee\xd3\xe5\xcd\xf7\xcf\x57\xd0\x72\x67\x77\xab\x2a\xfe\xc0\x4a\x6a\x6a\x81\x24\xa2\x01\xa5\xde\xad\x00\x00\xaa\x0e\x59\x82\x6a\xa5\x0f\xc8\xb5\xf8\x7a\x73\x9d\xbd\x16\xbb\x55\x7a\x3b\xcb\x32\x38\x38\x0f\x9d\xdb\x1b\x8b\x10\x94\x47\xa4\x00\x59\xb6\x8c\x25\xd9\x61\x2d\x46\x83\xc7\xde\x79\x16\xa0\x1c\x31\x12\xd7\xe2\x68\x34\xb7\xb5\xc6\xd1\x28\xcc\xe6\xcb\x39\x18\x32\x6c\xa4\xcd\x82\x92\x16\xeb\x8b\xbc\xfc\x2b\x57\xe0\xc9\x62\x68\x11\x39\x80\xf4\x18\xa1\x46\x24\x83\xc4\x76\x82\x80\xbd\xf4\x92\x51\x83\x21\x76\xa0\x5c\xd7\x3b\x42\xe2\x05\x1f\x6b\xe8\x16\x3c\xda\x5a\x3c\x22\x09\xe8\x50\x1b\x59\x0b\x69\xad\x80\xd6\xe3\xa1\x16\xa7\xd8\x22\xf9\x15\x01\x79\xe8\x73\x15\x82\xf8\x6f\x24\x39\x85\x67\x01\xf2\xd8\xdb\xe9\x59\x90\x78\xea\x0d\x35\xcf\x02\x65\xa8\x1f\x78\x89\x34\xdb\xd3\x39\x7e\x7b\xa7\x27\xf8\x75\xba\xce\x26\xa9\x6e\x1b\xef\x06\xd2\x5b\x78\xa1\x95\xd6\x58\xbe\x39\x39\xdc\xaf\x4e\xc7\x7c\x3f\xec\xf7\x16\xb3\xb8\x3e\xd2\x10\xfa\x27\x38\x2d\x9a\xa6\xe5\x2d\x5c\x94\xe5\xd8\xfe\x23\x42\x3e\xf3\xcd\x8e\x5e\xf6\xc0\x78\xc7\xd2\xa3\x7c\x02\xdb\x49\xdf\x18\xda\xc2\x82\x54\xfc\xe6\x75\xdd\x82\x92\x56\xad\x2f\xca\xf2\x25\x64\xf0\xaa\xec\xef\x36\xcb\xc4\x73\x03\x8a\x3f\x1d\xa8\x8a\xa4\xa9\x2a\xb6\x60\xb7\x5a\xcd\xfb\xfc\xc8\x04\x2d\x76\x48\x3c\xab\x49\xb5\x92\xe1\x68\x48\xbb\xe3\xbc\xba\x95\x36\x23\x18\x5d\x8b\xf8\x20\x76\x55\xa1\xcd\xf8\x80\x60\xba\x28\x2a\xe0\x16\xe1\xa3\x1c\xe5\x17\xe5\x4d\xcf\x70\x88\x6a\x9c\x43\x43\x32\x04\xaf\x96\xd3\x6a\x5d\xe7\xf2\x1f\x21\x62\x25\x87\x93\x67\x1a\x55\x51\x3c\x08\xd1\xfc\x44\xd8\x4f\x91\x68\x60\x3f\x28\x36\xd4\x40\x92\xb4\x86\xc3\x40\x8a\x8d\xa3\x3c\xcf\xe7\xa8\x51\x26\xee\xdf\x12\xf5\x1a\x08\x8f\xf0\xde\x75\x6e\xad\x9d\x1a\x62\x79\x79\x83\x7c\x95\x2a\x7d\x3b\x7d\xd0\xeb\x54\xd1\xe6\x1c\xc4\x63\x9c\x38\x5f\x4c\x60\x9e\xcf\xa5\xb4\x36\x6e\xc9\x35\x6d\x4f\x39\x61\xed\x36\x0b\xbf\xd4\xec\xfb\xcd\x6a\x51\x50\x91\x3a\xfd\x3b\x00\x00\xff\xff\x11\x0e\xbe\x52\xed\x04\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1261, mode: os.FileMode(420), modTime: time.Unix(1554528069, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"component/homo.js": componentHomoJs,
	"component/styles/input.css": componentStylesInputCss,
	"component/styles/reply.css": componentStylesReplyCss,
	"component/styles/says.css": componentStylesSaysCss,
	"component/styles/setup.css": componentStylesSetupCss,
	"component/styles/typing.css": componentStylesTypingCss,
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"component": &bintree{nil, map[string]*bintree{
		"homo.js": &bintree{componentHomoJs, map[string]*bintree{}},
		"styles": &bintree{nil, map[string]*bintree{
			"input.css": &bintree{componentStylesInputCss, map[string]*bintree{}},
			"reply.css": &bintree{componentStylesReplyCss, map[string]*bintree{}},
			"says.css": &bintree{componentStylesSaysCss, map[string]*bintree{}},
			"setup.css": &bintree{componentStylesSetupCss, map[string]*bintree{}},
			"typing.css": &bintree{componentStylesTypingCss, map[string]*bintree{}},
		}},
	}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

