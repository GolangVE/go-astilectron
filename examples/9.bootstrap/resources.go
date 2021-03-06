// Code generated by go-bindata.
// sources:
// resources/static/css/base.css
// resources/templates/index.html
// DO NOT EDIT!

package main

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

var _resourcesStaticCssBaseCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x31\x6e\xc3\x30\x0c\x45\x77\x9f\x82\x40\xd1\xad\x2e\x5c\x78\x53\x4f\x43\x8b\xb4\x44\x94\x96\x04\x95\xad\x6b\x14\xbe\x7b\x90\x28\xce\x92\x64\xfc\xef\x93\xff\x45\x5b\xf4\x0d\xa6\x4c\x1b\xfc\x77\x00\x00\x13\xfa\xaf\x50\xf3\x4f\xa2\xde\x67\xcd\xd5\xc1\xcb\x38\x8e\x9f\x97\xee\x00\xf3\x3c\x37\x10\x59\x42\x34\x07\x1f\xc3\xf0\xda\xc8\x82\x35\x48\x72\x30\xb4\x58\x90\x48\x52\xb8\xe5\x55\xc8\xe2\x71\xbf\x77\xdd\xfb\x5a\xb1\x14\xae\x57\x39\xc9\x77\x51\xdc\x1c\x18\x4e\xca\xcf\x1c\x77\x23\x3e\x27\xe3\x64\x0f\x47\x7a\xcf\xaa\xed\xcf\xf8\xcf\x7a\x54\x09\xc9\x81\xe7\x64\x5c\x1b\xff\xe5\x6a\xe2\x51\x8f\x6e\x11\xa2\xb3\x7d\x3f\x05\x00\x00\xff\xff\x50\x5c\x96\x33\x1d\x01\x00\x00")

func resourcesStaticCssBaseCssBytes() ([]byte, error) {
	return bindataRead(
		_resourcesStaticCssBaseCss,
		"resources/static/css/base.css",
	)
}

func resourcesStaticCssBaseCss() (*asset, error) {
	bytes, err := resourcesStaticCssBaseCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/static/css/base.css", size: 285, mode: os.FileMode(420), modTime: time.Unix(1494924234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x50\xbd\x4e\x03\x31\x0c\xde\xfb\x14\xc6\xcc\x34\x2b\x42\xbe\x2c\x05\x26\x24\x18\xca\xc0\xe8\xcb\xf9\xc8\xa9\xb9\x4b\x15\xbb\xa0\xaa\xea\xbb\xa3\xfb\xa9\x54\x95\x29\x89\xbf\xdf\x98\xee\x9e\xdf\x37\xdb\xaf\x8f\x17\x88\xd6\x27\xbf\xa2\xf1\x80\xc4\xc3\x77\x85\x32\xe0\x38\x10\x6e\xfc\x0a\x00\x80\x7a\x31\x86\x10\xb9\xa8\x58\x85\x9f\xdb\xd7\x87\x47\x5c\xa0\xd4\x0d\x3b\x28\x92\x2a\x54\x3b\x26\xd1\x28\x62\x08\xb1\x48\x5b\xa1\x53\x63\xeb\x82\x0b\xaa\xae\x66\x95\x75\x50\x45\xe7\x57\xe4\x66\x6f\xaa\x73\x73\x5c\x7c\x9a\xee\x07\x42\x62\xd5\x0a\x7f\x0b\xef\xf7\x52\x96\x84\x5b\x34\xe4\xc1\x64\xb0\x2b\xf4\xc2\xf0\xa7\x13\xac\xdf\xb8\x96\x04\xe7\x33\xb9\x71\xf2\x9f\x43\x7c\x29\x17\x0e\x6a\xb9\x77\x25\x1f\x4c\x10\xa6\xf6\xa3\x7b\xca\xe5\xe9\xbe\x6d\x5b\xf4\x9b\xd4\x85\x1d\x44\x29\x02\x96\x41\x45\xc0\xa2\xc0\x2c\x83\x49\x46\x8e\xfd\x4d\xd0\xd5\x73\xb9\x92\x9b\xbf\x49\x6e\xda\xf4\x5f\x00\x00\x00\xff\xff\xca\x6d\x64\x81\x79\x01\x00\x00")

func resourcesTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesTemplatesIndexHtml,
		"resources/templates/index.html",
	)
}

func resourcesTemplatesIndexHtml() (*asset, error) {
	bytes, err := resourcesTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/templates/index.html", size: 377, mode: os.FileMode(420), modTime: time.Unix(1494924234, 0)}
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
	"resources/static/css/base.css":  resourcesStaticCssBaseCss,
	"resources/templates/index.html": resourcesTemplatesIndexHtml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"css": &bintree{nil, map[string]*bintree{
				"base.css": &bintree{resourcesStaticCssBaseCss, map[string]*bintree{}},
			}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{resourcesTemplatesIndexHtml, map[string]*bintree{}},
		}},
	}},
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
