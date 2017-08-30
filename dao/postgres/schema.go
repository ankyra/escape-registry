// Code generated by go-bindata.
// sources:
// dao/postgres/schemas/1_initial_schema.down.sql
// dao/postgres/schemas/1_initial_schema.up.sql
// dao/postgres/schemas/2_project_metadata.down.sql
// dao/postgres/schemas/2_project_metadata.up.sql
// dao/postgres/schemas/3_migrate_existing_projects.up.sql
// dao/postgres/schemas/4_application_metadata.down.sql
// dao/postgres/schemas/4_application_metadata.up.sql
// dao/postgres/schemas/5_build_dependencies.down.sql
// dao/postgres/schemas/5_build_dependencies.up.sql
// dao/postgres/schemas/6_release_changes.up.sql
// dao/postgres/schemas/7_reset_dependencies.up.sql
// dao/postgres/schemas/8_is_extension_on_dependencies.up.sql
// dao/postgres/schemas/9_change_integer_precision.up.sql
// DO NOT EDIT!

package postgres

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

	info := bindataFileInfo{name: "1_initial_schema.down.sql", size: 56, mode: os.FileMode(436), modTime: time.Unix(1502613619, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x4e\xc3\x30\x10\x84\xef\x7e\x8a\x3d\xc6\x92\x2f\x04\x81\x90\x38\x05\x64\xa4\x88\x5f\xa5\x3e\xd0\x53\xb5\x72\x57\xc5\xd0\x24\xd6\xc6\xad\x78\x7c\x54\xe7\x17\x91\x00\xd7\x99\xc9\xe4\x9b\xf5\x6d\xa1\x33\xa3\xc1\x64\x37\x0f\x1a\xf2\x3b\x78\x7a\x36\xa0\x5f\xf3\x95\x59\x01\xd3\x9e\xb0\x21\x48\x04\x00\x40\x85\x25\xc1\x11\xd9\xbe\x21\x27\x67\xe9\x95\x54\x10\xf5\x2e\xb5\x71\xdb\xc1\x4d\x2f\x2e\xa5\x8a\xe6\x91\xb8\x71\x75\x35\x38\xe7\x69\x67\x94\x14\x70\x8b\x01\x21\xd0\x67\x68\x25\xcf\xf5\x3b\xd9\xf0\x33\xfb\x52\xe4\x8f\x59\xb1\x86\x7b\xbd\x4e\x4e\x14\xaa\xaf\x55\xfd\x37\x52\xc8\x6b\x21\x7e\x99\xe2\xd1\x7e\xe0\xae\x9f\xb2\xf8\xa7\xa5\x2d\xed\xd2\x03\xbb\x39\x79\x8a\x37\x16\xa8\x53\xfc\xff\x80\x68\xf7\x7f\xc1\xed\xb8\x3e\xf8\xcd\xb7\x67\x18\x0f\xed\x89\x4b\xd7\xc4\x5b\xbb\x2a\xcc\x90\x75\xb5\x6a\x52\x13\xa1\xbe\x02\x00\x00\xff\xff\x53\x97\x8f\x29\x00\x02\x00\x00")

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

	info := bindataFileInfo{name: "1_initial_schema.up.sql", size: 512, mode: os.FileMode(436), modTime: time.Unix(1502613606, 0)}
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

	info := bindataFileInfo{name: "2_project_metadata.down.sql", size: 20, mode: os.FileMode(436), modTime: time.Unix(1502797856, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_project_metadataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x28\xca\xcf\x4a\x4d\x2e\x51\xd0\xe0\x52\x50\x50\x50\xc8\x4b\xcc\x4d\x55\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd2\xd4\x01\x8b\xa6\xa4\x16\x27\x17\x65\x16\x94\x64\xe6\xe7\x29\x94\xa4\x56\x94\x40\x44\xf3\x8b\xd2\x43\x83\x7c\xe0\xaa\x8d\x4c\xcd\xa0\xca\x73\xf2\xd3\xf3\x91\xd4\x05\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x6a\x80\xcc\xd7\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\x46\xec\xb6\x00\x84\x00\x00\x00")

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

	info := bindataFileInfo{name: "2_project_metadata.up.sql", size: 132, mode: os.FileMode(436), modTime: time.Unix(1502815779, 0)}
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

	info := bindataFileInfo{name: "3_migrate_existing_projects.up.sql", size: 104, mode: os.FileMode(436), modTime: time.Unix(1502806903, 0)}
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

	info := bindataFileInfo{name: "4_application_metadata.down.sql", size: 24, mode: os.FileMode(436), modTime: time.Unix(1502815842, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_application_metadataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4b\xc3\x40\x14\x84\xef\xfb\x2b\xe6\xd6\x2c\xe4\x62\x44\x11\x7a\x5a\xd3\x17\x5c\x4c\x13\xd9\x7d\x4a\x7b\x92\x25\x3e\x24\x12\x9b\x90\x2c\xe2\xcf\x97\xa6\x46\x84\x5c\xe7\xfb\x18\x66\x72\x47\x86\x09\x6c\xee\x4b\x82\x2d\x50\xd5\x0c\x3a\x58\xcf\x1e\x61\x18\xba\xb6\x09\xb1\xed\x4f\x48\x14\x00\x9c\xc2\xa7\xe0\xc5\xb8\xfc\xc1\xb8\xe4\x2a\xbb\xd3\x29\xe6\x7c\x18\xfb\x0f\x69\xe2\x1f\xba\xce\x74\x3a\x83\x37\x99\x9a\xb1\x1d\xe6\x0a\xa6\x03\x63\x47\x85\x79\x2e\x19\x9b\xcd\x45\xe8\x42\x94\x29\xbe\x7e\xc9\x38\x9d\x9d\xa5\x20\xbb\xb9\xd5\x6b\xb7\x7f\xef\x11\xe5\x3b\xae\xc8\x93\xb3\x7b\xe3\x8e\x78\xa4\x63\x72\xde\x98\x2e\x8b\xb4\xd2\x5b\xa5\x6c\xe5\xc9\x31\x6c\xc5\xf5\xff\x53\xc9\xaf\x94\xce\xbf\xb4\xf2\x54\x52\xce\xd8\x59\xcf\xb6\xca\x19\x0b\xd7\x17\x41\x15\xae\xde\x63\x94\x4e\xc2\x24\x5b\xf5\x13\x00\x00\xff\xff\xeb\x2f\x52\x2f\x39\x01\x00\x00")

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

	info := bindataFileInfo{name: "4_application_metadata.up.sql", size: 313, mode: os.FileMode(436), modTime: time.Unix(1502825603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_build_dependenciesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\x49\x2d\x48\xcd\x4b\x49\xcd\x4b\xae\xb4\xe6\x02\x04\x00\x00\xff\xff\xbd\x8f\xb0\xd2\x29\x00\x00\x00")

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

	info := bindataFileInfo{name: "5_build_dependencies.down.sql", size: 41, mode: os.FileMode(436), modTime: time.Unix(1503531283, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_build_dependenciesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x6b\x83\x40\x14\x84\xef\xfe\x8a\x39\x2a\xec\xa5\xf6\x52\xe8\x69\x95\x2d\x95\x5a\x2d\xeb\x52\xe2\x49\x8c\xfb\x0e\x06\xb3\xbb\x68\x12\xf0\xdf\x07\xc5\x48\x48\xe2\x75\xbe\x79\xc3\xf7\x62\x29\xb8\x12\x50\x3c\x4a\x05\x92\x2f\x64\xb9\x82\xd8\x25\x85\x2a\xd0\x53\x47\xf5\x40\x95\x26\x47\x46\x93\x69\x46\xf8\x1e\x00\xb8\xde\x1e\xa8\x39\xe1\x9f\xcb\xf8\x9b\x4b\xff\x3d\x0c\xd8\x0c\x4c\x7d\xa4\x35\x7d\x0b\x3f\x96\xf8\x42\xfd\xd0\x5a\xf3\xdc\xd7\xe4\xaa\xcd\xb1\x09\x6e\x0c\x4e\x68\x73\x74\x7f\x6e\x3b\x5d\x0d\x8d\x75\x84\x28\xcf\x53\xc1\xb3\xf5\xaa\xb3\xe3\x2b\xf2\x27\x93\x5f\x2e\x4b\xfc\x88\x12\xfe\x22\xc4\xe6\x6f\xd8\x4d\x9e\x61\x6e\x3e\x58\xb3\xd5\x92\xdd\x4b\x05\x5e\xf0\xe9\x5d\x03\x00\x00\xff\xff\x7e\x8c\x8a\x02\x5a\x01\x00\x00")

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

	info := bindataFileInfo{name: "5_build_dependencies.up.sql", size: 346, mode: os.FileMode(436), modTime: time.Unix(1503531268, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_release_changesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd0\xcb\x4a\xc4\x30\x14\xc6\xf1\x7d\x9f\xe2\xec\xaa\x0b\x41\x41\xdc\x74\x95\xb6\xf1\x02\xb1\x85\x90\xba\x2d\xc7\xe4\x54\x82\x21\x09\x4d\x45\xf4\xe9\xc5\x55\x19\xe6\x92\x19\xfa\x00\xdf\x8f\x3f\x1f\x13\x8a\x4b\x50\xac\x16\x1c\x66\x72\x84\x89\x80\xb5\x2d\x34\xbd\x18\x5e\x3b\x88\x73\xd0\x94\x12\x99\xd1\x50\x24\x6f\xc8\x6b\x4b\x09\xea\xbe\x17\x9c\x75\xd0\xf2\x47\x36\x08\x05\xe5\x84\x2e\x51\x59\x15\x19\xce\x84\x6f\xef\x02\x9a\x04\x2f\x9d\xe2\x4f\x5c\xae\xc2\x6d\x7e\xfd\x15\xff\xb7\x64\xc6\xf7\x1f\x78\x63\xb2\x79\x66\xf2\xea\xe1\xfe\x7a\x35\x2e\x20\x70\x39\x9c\xb0\x03\x60\x8c\xce\x6a\x5c\x6c\xf0\x5b\x3a\x72\xcc\x59\x2d\x11\xf5\x27\x7e\x6c\xfa\xe3\x14\x71\xac\x21\xb3\x9f\xac\xa3\x64\x7f\x69\x7f\x7c\x73\x57\x56\xc5\x5f\x00\x00\x00\xff\xff\xf1\x23\x83\x34\x5e\x02\x00\x00")

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

	info := bindataFileInfo{name: "6_release_changes.up.sql", size: 606, mode: os.FileMode(436), modTime: time.Unix(1504062920, 0)}
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

	info := bindataFileInfo{name: "7_reset_dependencies.up.sql", size: 85, mode: os.FileMode(436), modTime: time.Unix(1504045907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_is_extension_on_dependenciesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xb1\xaa\x83\x40\x10\x85\xe1\xde\xa7\x38\x9d\x0f\x21\xb7\x58\xef\x8e\xd5\xe8\x06\x33\xd6\x22\x7a\x02\x82\xac\xe2\xa4\x48\xde\x3e\x55\xb0\x49\xfd\xff\x7c\x51\x54\x4c\xd0\xf4\xa9\xc5\xc9\x8d\x93\x73\x5c\x78\x30\x2f\xcc\xf3\xbb\x2a\x86\x5b\x0c\x26\xdf\x84\xbb\x18\x8e\x73\x9f\xe9\xce\xe5\x1a\x57\x3a\xfe\x50\x3e\xa6\xcd\x59\x56\x45\x50\x93\x1e\x16\x6a\x95\x1f\x28\x42\x8c\xf8\x4f\x3a\xb4\x1d\x56\x1f\xf9\x7a\x32\xfb\xba\x67\xd4\x29\xa9\x84\x0e\x51\x9a\x30\xa8\x5d\xde\x27\x00\x00\xff\xff\x15\x0c\x98\xe0\xa5\x00\x00\x00")

func _8_is_extension_on_dependenciesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_is_extension_on_dependenciesUpSql,
		"8_is_extension_on_dependencies.up.sql",
	)
}

func _8_is_extension_on_dependenciesUpSql() (*asset, error) {
	bytes, err := _8_is_extension_on_dependenciesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_is_extension_on_dependencies.up.sql", size: 165, mode: os.FileMode(436), modTime: time.Unix(1504046895, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_change_integer_precisionUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\xcd\x49\x4d\x2c\x4e\x55\x80\x88\x39\xfb\xfb\x84\xfa\xfa\x29\xa4\xe4\x97\xe7\xe5\xe4\x27\xa6\x14\x2b\x84\x44\x06\xb8\x2a\x38\x79\xba\x7b\xfa\x85\x58\x13\xd6\x57\x5a\x00\xd2\x95\x9a\x12\x9f\x58\x82\x5b\x67\x62\x41\x41\x4e\x66\x72\x62\x49\x66\x7e\x1e\x19\xba\x0b\x12\x93\xb3\x13\xd3\xc9\xb1\x17\xab\xce\xb4\xcc\x9c\xd4\xe2\xcc\xaa\x54\x54\x6d\x80\x00\x00\x00\xff\xff\xe0\x40\xf0\xda\x22\x01\x00\x00")

func _9_change_integer_precisionUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__9_change_integer_precisionUpSql,
		"9_change_integer_precision.up.sql",
	)
}

func _9_change_integer_precisionUpSql() (*asset, error) {
	bytes, err := _9_change_integer_precisionUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "9_change_integer_precision.up.sql", size: 290, mode: os.FileMode(436), modTime: time.Unix(1504063180, 0)}
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
	"8_is_extension_on_dependencies.up.sql": _8_is_extension_on_dependenciesUpSql,
	"9_change_integer_precision.up.sql": _9_change_integer_precisionUpSql,
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
	"8_is_extension_on_dependencies.up.sql": &bintree{_8_is_extension_on_dependenciesUpSql, map[string]*bintree{}},
	"9_change_integer_precision.up.sql": &bintree{_9_change_integer_precisionUpSql, map[string]*bintree{}},
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
