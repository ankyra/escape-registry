// Code generated by go-bindata.
// sources:
// dao/sqlite/schemas/1_initial_schema.down.sql
// dao/sqlite/schemas/1_initial_schema.up.sql
// dao/sqlite/schemas/2_project_metadata.down.sql
// dao/sqlite/schemas/2_project_metadata.up.sql
// dao/sqlite/schemas/3_migrate_existing_projects.up.sql
// dao/sqlite/schemas/4_application_metadata.down.sql
// dao/sqlite/schemas/4_application_metadata.up.sql
// dao/sqlite/schemas/5_build_dependencies.down.sql
// dao/sqlite/schemas/5_build_dependencies.up.sql
// dao/sqlite/schemas/6_release_changes.up.sql
// dao/sqlite/schemas/7_reset_dependencies.up.sql
// dao/sqlite/schemas/8_is_extension_on_dependency_table.up.sql
// DO NOT EDIT!

package sqlite

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

var __1_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4a\xcd\x49\x4d\x2c\x4e\xb5\xe6\x42\x12\x2b\x48\x4c\xce\x4e\x4c\x47\x15\x4b\x4c\xce\xb1\xe6\x02\x04\x00\x00\xff\xff\xd7\x9e\x4f\xa4\x38\x00\x00\x00")

func _1_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initial_schemaDownSql,
		"1_initial_schema.down.sql",
	)
}

func _1_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initial_schema.down.sql", size: 56, mode: os.FileMode(436), modTime: time.Unix(1502613569, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x4d\x4f\x84\x30\x10\x86\xef\xfd\x15\x73\x84\xa4\x17\xcf\x9e\xd0\xd4\x84\xf8\x19\xb6\x07\xf7\xb4\x99\x74\x27\x58\x5d\xa0\x99\x16\x7e\xbf\xa1\x52\xa5\xc1\xb0\xd7\xa7\x9d\x77\x9e\x79\xef\x1b\x55\x69\x05\xba\xba\x7b\x52\x50\x3f\xc0\xcb\xab\x06\xf5\x5e\x1f\xf4\x01\x98\x2e\x84\x9e\xa0\x10\x00\x00\x3d\x76\x04\x3e\xb0\xed\x5b\x09\x91\x2c\xef\x27\x7b\x4e\x3c\xe2\x89\xd8\xdb\xa1\xcf\x58\x47\x01\xcf\x18\x30\x83\x8e\x87\x4f\x32\x21\x63\x6f\x4d\xfd\x5c\x35\x47\x78\x54\xc7\x62\xde\x28\x53\x9c\x4c\xdf\x4b\x51\xde\x0a\xb1\xa3\xed\xd0\x7c\x61\x9b\xb4\xff\x5b\xb2\x15\xff\x39\x68\x64\x9b\x83\xb5\xcd\x12\x24\x57\xd3\x72\x9e\xb8\xea\x83\xe6\xb2\xe3\xd2\xf2\x30\xba\xd3\xb6\x5c\x47\xdc\x59\x1f\x8b\x9c\x90\xcd\x07\x72\x71\x53\x6e\x3b\xfa\xb5\xfa\xcb\x89\x42\xdf\x01\x00\x00\xff\xff\x97\x6a\x54\x4e\xd7\x01\x00\x00")

func _1_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_initial_schemaUpSql,
		"1_initial_schema.up.sql",
	)
}

func _1_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_initial_schema.up.sql", size: 471, mode: os.FileMode(436), modTime: time.Unix(1502815314, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_project_metadataDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\xca\xcf\x4a\x4d\x2e\xb1\xe6\x02\x04\x00\x00\xff\xff\xa5\x8e\xd4\xaa\x14\x00\x00\x00")

func _2_project_metadataDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_project_metadataDownSql,
		"2_project_metadata.down.sql",
	)
}

func _2_project_metadataDownSql() (*asset, error) {
	bytes, err := _2_project_metadataDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_project_metadata.down.sql", size: 20, mode: os.FileMode(436), modTime: time.Unix(1502795863, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_project_metadataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x28\xca\xcf\x4a\x4d\x2e\x51\xd0\xe0\x52\x50\x50\x50\xc8\x4b\xcc\x4d\x55\x28\x2e\x29\xca\xcc\x4b\xd7\x01\x0b\xa4\xa4\x16\x27\x17\x65\x16\x94\x64\xe6\xe7\xa1\x88\xe7\x17\xa5\x87\x06\xf9\xa0\x08\xe5\xe4\xa7\xe7\xa3\x08\x04\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x2a\x68\x80\x4c\xd6\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\xb1\x38\xae\x06\x7e\x00\x00\x00")

func _2_project_metadataUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_project_metadataUpSql,
		"2_project_metadata.up.sql",
	)
}

func _2_project_metadataUpSql() (*asset, error) {
	bytes, err := _2_project_metadataUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_project_metadata.up.sql", size: 126, mode: os.FileMode(436), modTime: time.Unix(1502807067, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_migrate_existing_projectsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x28\x28\xca\xcf\x4a\x4d\x2e\xd1\xc8\x4b\xcc\x4d\xd5\x51\x48\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\xd3\x51\xc8\x2f\x4a\x0f\x0d\xf2\xd1\x51\xc8\xc9\x4f\xcf\xd7\xe4\x0a\x76\xf5\x71\x75\x0e\x51\x48\xc9\x2c\x2e\xc9\xcc\x4b\x2e\xd1\x80\x6a\xd4\xd4\x51\x50\x57\x87\x61\x2e\xb7\x20\x7f\x5f\x85\xa2\xd4\x9c\xd4\xc4\xe2\x54\x6b\x2e\x40\x00\x00\x00\xff\xff\x5b\xed\x91\x00\x68\x00\x00\x00")

func _3_migrate_existing_projectsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_migrate_existing_projectsUpSql,
		"3_migrate_existing_projects.up.sql",
	)
}

func _3_migrate_existing_projectsUpSql() (*asset, error) {
	bytes, err := _3_migrate_existing_projectsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_migrate_existing_projects.up.sql", size: 104, mode: os.FileMode(436), modTime: time.Unix(1502806779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_application_metadataDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\xb3\xe6\x02\x04\x00\x00\xff\xff\xc6\x19\x92\xd8\x18\x00\x00\x00")

func _4_application_metadataDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_application_metadataDownSql,
		"4_application_metadata.down.sql",
	)
}

func _4_application_metadataDownSql() (*asset, error) {
	bytes, err := _4_application_metadataDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_application_metadata.down.sql", size: 24, mode: os.FileMode(436), modTime: time.Unix(1502815392, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_application_metadataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x4a\xc4\x30\x18\x84\xef\x79\x8a\xb9\x6d\x0b\x7d\x83\x3d\xd5\xfa\x2f\x04\x77\x5b\x49\xfe\x83\x7b\x92\xd0\xfd\x2d\x95\xda\x84\x24\xfa\xfc\x62\x6a\xc1\x8b\x7b\xfd\xe6\x1b\x98\xe9\x0c\xb5\x4c\xe0\xf6\xe1\x4c\xd0\x27\xf4\x03\x83\x5e\xb4\x65\x0b\x17\xc2\x32\x8f\x2e\xcf\x7e\x45\xa5\x00\x60\x75\x1f\x82\x94\xe3\xbc\x4e\x0d\x0a\x09\xd1\xbf\xcb\x98\x77\x58\xd8\x4d\xd2\x18\xe7\x50\x7a\x1b\xc7\x4d\xde\xdc\xe7\x92\x71\x38\x6c\xca\xe2\xb2\xa4\xfc\xfa\x25\x31\xdd\xb3\xfc\xe4\xff\xcb\x9e\x8d\xbe\xb4\xe6\x8a\x27\xba\x56\x3f\xb3\x9a\x7d\x4a\xad\xea\xa3\x52\xba\xb7\x64\x18\xba\xe7\xe1\xef\x8f\xea\x57\x6a\xca\x95\x5a\x59\x3a\x53\xc7\x78\xd4\x96\x75\xdf\x31\xf6\xbc\xde\x04\x75\x32\xc3\x05\x51\x16\x71\x49\x8e\xea\x3b\x00\x00\xff\xff\x83\x99\xe1\x3f\x2c\x01\x00\x00")

func _4_application_metadataUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_application_metadataUpSql,
		"4_application_metadata.up.sql",
	)
}

func _4_application_metadataUpSql() (*asset, error) {
	bytes, err := _4_application_metadataUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_application_metadata.up.sql", size: 300, mode: os.FileMode(436), modTime: time.Unix(1502825588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_build_dependenciesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\x49\x2d\x48\xcd\x4b\x49\xcd\x4b\xae\xb4\xe6\x02\x04\x00\x00\xff\xff\xf4\xa7\x3f\x11\x1f\x00\x00\x00")

func _5_build_dependenciesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_build_dependenciesDownSql,
		"5_build_dependencies.down.sql",
	)
}

func _5_build_dependenciesDownSql() (*asset, error) {
	bytes, err := _5_build_dependenciesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_build_dependencies.down.sql", size: 31, mode: os.FileMode(436), modTime: time.Unix(1503525858, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_build_dependenciesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xbd\x0a\x83\x30\x14\xc5\xf1\xdd\xa7\x38\xa3\x42\xde\xa0\x93\x2d\x29\x48\x3f\xd1\x0c\x75\x12\x35\x97\x62\x49\x93\x90\xd8\x82\x6f\x5f\x94\x56\x94\x66\xfd\x5f\xee\xe1\xb7\xcb\x79\x2a\x38\x44\xba\x3d\x72\x64\x7b\x9c\x2f\x02\xfc\x96\x15\xa2\x80\x23\x45\xb5\xa7\x4a\x92\x25\x2d\x49\xb7\x03\xe2\x08\x00\xac\x33\x0f\x6a\x7b\xf8\xde\x75\xfa\xce\xa6\xa6\xeb\x27\xad\xc2\x9b\x9c\xef\x8c\x5e\x35\x49\xb6\x0a\x3d\x8f\xfd\x6f\x60\x8c\xa1\x91\xe6\xd5\x29\x59\xf9\xd6\x58\x42\x63\x8c\xa2\x5a\xcf\x0f\xca\x0c\xa1\xcb\x35\xcf\x4e\x69\x5e\xe2\xc0\x4b\xc4\x5f\x00\x9b\xc8\xec\xe7\x64\x4b\x1c\x9b\x45\x6c\xc9\x48\xa2\x64\x13\x7d\x02\x00\x00\xff\xff\x78\x04\xf8\x44\x31\x01\x00\x00")

func _5_build_dependenciesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_build_dependenciesUpSql,
		"5_build_dependencies.up.sql",
	)
}

func _5_build_dependenciesUpSql() (*asset, error) {
	bytes, err := _5_build_dependenciesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_build_dependencies.up.sql", size: 305, mode: os.FileMode(436), modTime: time.Unix(1503525852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_release_changesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xd0\xcf\x6a\x84\x30\x10\xc7\xf1\xbb\x4f\x31\x37\x4f\x85\xf6\xec\x29\xd6\x54\x84\x34\x82\x8d\x67\x99\x26\x63\x09\x0d\x49\x30\x96\xd2\x3e\xfd\xb2\x27\x59\xf6\x4f\x76\x7d\x80\xcf\x77\x7e\x0c\x13\x8a\x0f\xa0\x58\x2d\x38\x2c\xe4\x08\x13\x01\x6b\x1a\x78\xed\xc5\xf8\x2e\x21\x2e\x41\x53\x4a\x64\x26\x43\x91\xbc\x21\xaf\x2d\x25\xa8\xfb\x5e\x70\x26\xa1\xe1\x6f\x6c\x14\x0a\xca\x19\x5d\xa2\xb2\x2a\x32\x39\x13\x7e\xbd\x0b\x68\x12\x74\x52\xf1\x96\x0f\x5b\xe1\x39\xaf\x7f\xe2\xd1\x92\x99\x3e\xff\xe0\x43\x0d\x9d\x6c\x37\xfe\x80\xc6\xf5\xf2\xf5\x93\x00\xc6\xe8\xac\xc6\xd5\x06\xbf\x73\x42\xae\x70\xd7\x8c\x88\xfa\x1b\xbf\xf6\x7e\xe1\x96\xbe\x76\x3e\xe3\x67\xeb\x28\xd9\x7f\x3a\xc7\x4f\x2f\x65\x55\x1c\x02\x00\x00\xff\xff\xdf\x38\xf6\x44\x4f\x02\x00\x00")

func _6_release_changesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_release_changesUpSql,
		"6_release_changes.up.sql",
	)
}

func _6_release_changesUpSql() (*asset, error) {
	bytes, err := _6_release_changesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_release_changes.up.sql", size: 591, mode: os.FileMode(436), modTime: time.Unix(1504013091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_reset_dependenciesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x71\xf5\x71\x0d\x71\x55\x70\x0b\xf2\xf7\x55\x28\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\x49\x2d\x48\xcd\x4b\x49\xcd\x4b\xae\xb4\xe6\x0a\x0d\x70\x71\x0c\x71\x85\x49\x29\x04\xbb\x86\x28\x14\x14\xe5\x27\xa7\x16\x17\xa7\xa6\x20\x14\x66\xa6\x16\x2b\xd8\x2a\xa8\xa7\x25\xe6\x14\xa7\xaa\x5b\x73\x01\x02\x00\x00\xff\xff\xb0\x76\xaa\xa7\x55\x00\x00\x00")

func _7_reset_dependenciesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_reset_dependenciesUpSql,
		"7_reset_dependencies.up.sql",
	)
}

func _7_reset_dependenciesUpSql() (*asset, error) {
	bytes, err := _7_reset_dependenciesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_reset_dependencies.up.sql", size: 85, mode: os.FileMode(436), modTime: time.Unix(1504045627, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_is_extension_on_dependency_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xb1\xaa\x83\x40\x10\x85\xe1\xde\xa7\x38\x9d\x0f\x21\xb7\x58\xef\x8e\xd5\xe8\x06\x33\xd6\x22\x7a\x02\x82\xac\xe2\xa4\x48\xde\x3e\x55\xb0\x49\xfd\xff\x7c\x51\x54\x4c\xd0\xf4\xa9\xc5\xc9\x8d\x93\x73\x5c\x78\x30\x2f\xcc\xf3\xbb\x2a\x86\x5b\x0c\x26\xdf\x84\xbb\x18\x8e\x73\x9f\xe9\xce\xe5\x1a\x57\x3a\xfe\x50\x3e\xa6\xcd\x59\x56\x45\x50\x93\x1e\x16\x6a\x95\x1f\x28\x42\x8c\xf8\x4f\x3a\xb4\x1d\x56\x1f\xf9\x7a\x32\xfb\xba\x67\xd4\x29\xa9\x84\x0e\x51\x9a\x30\xa8\x5d\xde\x27\x00\x00\xff\xff\x15\x0c\x98\xe0\xa5\x00\x00\x00")

func _8_is_extension_on_dependency_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_is_extension_on_dependency_tableUpSql,
		"8_is_extension_on_dependency_table.up.sql",
	)
}

func _8_is_extension_on_dependency_tableUpSql() (*asset, error) {
	bytes, err := _8_is_extension_on_dependency_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_is_extension_on_dependency_table.up.sql", size: 165, mode: os.FileMode(436), modTime: time.Unix(1504046711, 0)}
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
	"1_initial_schema.down.sql": _1_initial_schemaDownSql,
	"1_initial_schema.up.sql": _1_initial_schemaUpSql,
	"2_project_metadata.down.sql": _2_project_metadataDownSql,
	"2_project_metadata.up.sql": _2_project_metadataUpSql,
	"3_migrate_existing_projects.up.sql": _3_migrate_existing_projectsUpSql,
	"4_application_metadata.down.sql": _4_application_metadataDownSql,
	"4_application_metadata.up.sql": _4_application_metadataUpSql,
	"5_build_dependencies.down.sql": _5_build_dependenciesDownSql,
	"5_build_dependencies.up.sql": _5_build_dependenciesUpSql,
	"6_release_changes.up.sql": _6_release_changesUpSql,
	"7_reset_dependencies.up.sql": _7_reset_dependenciesUpSql,
	"8_is_extension_on_dependency_table.up.sql": _8_is_extension_on_dependency_tableUpSql,
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
	"1_initial_schema.down.sql": &bintree{_1_initial_schemaDownSql, map[string]*bintree{}},
	"1_initial_schema.up.sql": &bintree{_1_initial_schemaUpSql, map[string]*bintree{}},
	"2_project_metadata.down.sql": &bintree{_2_project_metadataDownSql, map[string]*bintree{}},
	"2_project_metadata.up.sql": &bintree{_2_project_metadataUpSql, map[string]*bintree{}},
	"3_migrate_existing_projects.up.sql": &bintree{_3_migrate_existing_projectsUpSql, map[string]*bintree{}},
	"4_application_metadata.down.sql": &bintree{_4_application_metadataDownSql, map[string]*bintree{}},
	"4_application_metadata.up.sql": &bintree{_4_application_metadataUpSql, map[string]*bintree{}},
	"5_build_dependencies.down.sql": &bintree{_5_build_dependenciesDownSql, map[string]*bintree{}},
	"5_build_dependencies.up.sql": &bintree{_5_build_dependenciesUpSql, map[string]*bintree{}},
	"6_release_changes.up.sql": &bintree{_6_release_changesUpSql, map[string]*bintree{}},
	"7_reset_dependencies.up.sql": &bintree{_7_reset_dependenciesUpSql, map[string]*bintree{}},
	"8_is_extension_on_dependency_table.up.sql": &bintree{_8_is_extension_on_dependency_tableUpSql, map[string]*bintree{}},
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
