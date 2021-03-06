// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 20210421131113_initial_database_structure.down.sql (20B)
// 20210421131113_initial_database_structure.up.sql (126B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __20210421131113_initial_database_structureDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xc9\x4f\xc9\x8f\xcf\xc9\x2c\x2e\x01\x04\x00\x00\xff\xff\x7e\x27\x5d\x6a\x14\x00\x00\x00")

func _20210421131113_initial_database_structureDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210421131113_initial_database_structureDownSql,
		"20210421131113_initial_database_structure.down.sql",
	)
}

func _20210421131113_initial_database_structureDownSql() (*asset, error) {
	bytes, err := _20210421131113_initial_database_structureDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210421131113_initial_database_structure.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0x83, 0x30, 0x5d, 0x51, 0xf7, 0x78, 0xe7, 0xd3, 0x8f, 0x23, 0xa5, 0x9c, 0x16, 0x7b, 0x8d, 0x25, 0x1b, 0x20, 0x6a, 0xde, 0x95, 0xce, 0xf2, 0x71, 0xc8, 0xbf, 0x52, 0xc4, 0x2d, 0xb2, 0xd}}
	return a, nil
}

var __20210421131113_initial_database_structureUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xc9\x31\x0a\xc2\x30\x14\x06\xe0\xbd\xa7\xf8\x47\x05\x2f\xf1\xa2\x7f\x25\xf8\x4c\x24\x7d\x05\x3b\x39\x98\x0e\x01\x69\xc4\xe6\xfe\x08\x22\x9d\xbf\x63\xa2\x18\x61\xe2\x94\x68\x35\xd7\xc7\xab\xac\x6d\xd7\x01\x40\xc9\x70\xfe\x3c\x30\x79\x51\xdc\x92\xbf\x4a\x9a\x70\xe1\x74\xf8\x69\x9e\xd7\xe7\xa7\xbc\x5b\xa9\x0b\x8c\x77\x43\x88\x86\x30\xaa\xfe\xb9\x2e\x33\x5c\x8c\x4a\x09\x1b\xe1\xc4\x5e\x46\x35\xf4\xa2\x03\xbb\xfd\x37\x00\x00\xff\xff\x89\xac\x7f\x86\x7e\x00\x00\x00")

func _20210421131113_initial_database_structureUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210421131113_initial_database_structureUpSql,
		"20210421131113_initial_database_structure.up.sql",
	)
}

func _20210421131113_initial_database_structureUpSql() (*asset, error) {
	bytes, err := _20210421131113_initial_database_structureUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210421131113_initial_database_structure.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x47, 0x5, 0xf8, 0x85, 0x42, 0xa2, 0xf8, 0xdf, 0x2, 0x6c, 0x44, 0x21, 0x1b, 0x47, 0xfc, 0x39, 0xf4, 0x81, 0x52, 0xd0, 0xa9, 0x46, 0xc6, 0xe0, 0x75, 0x52, 0xc7, 0xd4, 0xdb, 0x5b, 0x9b, 0xbe}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"20210421131113_initial_database_structure.down.sql": _20210421131113_initial_database_structureDownSql,
	"20210421131113_initial_database_structure.up.sql":   _20210421131113_initial_database_structureUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"20210421131113_initial_database_structure.down.sql": {_20210421131113_initial_database_structureDownSql, map[string]*bintree{}},
	"20210421131113_initial_database_structure.up.sql":   {_20210421131113_initial_database_structureUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
