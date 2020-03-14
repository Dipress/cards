package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __20200228130253_cards_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4e\x2c\x4a\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x99\x4b\x9f\x4a\x1c\x00\x00\x00")

func _20200228130253_cards_down_sql() ([]byte, error) {
	return bindata_read(
		__20200228130253_cards_down_sql,
		"20200228130253_cards.down.sql",
	)
}

var __20200228130253_cards_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\xc1\x4b\xc3\x30\x14\xc7\xf1\xb3\xf9\x2b\x7e\xc7\x6d\x08\x03\x61\x27\x4f\xcf\xfa\x86\xc1\xb4\x8e\xd7\x57\x71\xa7\x12\xda\x1c\x02\x6e\x2b\x49\xc4\x7f\x5f\xec\xa1\xcc\x83\x08\xe6\x9a\xcf\xef\x0b\xaf\x12\x26\x65\x28\x3d\x38\x86\xdd\xa3\x79\x51\xf0\x9b\x6d\xb5\xc5\xe0\xd3\x98\xb1\x32\x40\x1c\x71\xf5\x5a\x16\x4b\x0e\x07\xb1\x35\xc9\x11\xcf\x7c\xbc\x35\xc0\x47\x0e\xa9\x5f\xa0\x6d\x74\x4e\x35\x9d\x73\xdf\xbf\x9f\x97\x74\xd5\x78\x25\xa9\x9e\x48\x56\x77\xbb\xdd\xfa\x07\x2b\xc9\x9f\xf3\x90\xe2\x54\xe2\xe5\xfc\x07\x7b\xf7\x33\xfa\xb5\x66\x80\xed\x06\x25\x9e\x42\x2e\xfe\x34\x61\xb3\x35\xc0\x90\x82\x2f\x61\xec\x7d\xb9\x01\xd4\xd6\xdc\x2a\xd5\x87\x65\x85\x47\xde\x53\xe7\x14\x55\x27\xc2\x8d\xf6\x0b\x99\x4f\x9c\xc6\xff\x8d\xcd\xfa\xde\x7c\x05\x00\x00\xff\xff\x13\x08\xec\xba\x69\x01\x00\x00")

func _20200228130253_cards_up_sql() ([]byte, error) {
	return bindata_read(
		__20200228130253_cards_up_sql,
		"20200228130253_cards.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"20200228130253_cards.down.sql": _20200228130253_cards_down_sql,
	"20200228130253_cards.up.sql": _20200228130253_cards_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"20200228130253_cards.down.sql": &_bintree_t{_20200228130253_cards_down_sql, map[string]*_bintree_t{
	}},
	"20200228130253_cards.up.sql": &_bintree_t{_20200228130253_cards_up_sql, map[string]*_bintree_t{
	}},
}}
