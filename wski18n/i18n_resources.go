/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by go-bindata.
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package wski18n

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

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesDe_deAllJson,
        "wski18n/resources/de_DE.all.json",
    )
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
    bytes, err := wski18nResourcesDe_deAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x6f\xdb\x38\x13\xbe\xfb\x57\x0c\x72\xf1\x25\xd0\x7b\xef\x2d\xe8\x9b\x6d\x83\x6e\x93\x20\xc9\x6e\x11\x74\x0b\x84\x11\x47\x16\x1b\x8a\x14\x48\x2a\x85\x57\xf0\x7f\x5f\x50\x1f\xb6\x93\x90\x14\x25\x7f\xb4\x0b\xec\xcd\xb6\x38\xcf\xf3\x0c\x87\x9c\x0f\xf9\xeb\x0c\xa0\x9e\x01\x00\x9c\x30\x7a\xf2\x0e\x4e\x3e\x22\xe7\xf2\xe4\xb4\xfd\xc9\x28\x22\x34\x27\x86\x49\x61\x9f\x9d\x09\x38\xbb\xbe\x80\x5c\x6a\x03\x45\xa5\x0d\x3c\x22\x94\x4a\x3e\x33\x8a\x34\x39\x99\x01\xac\x4e\x5f\xc3\x7d\x66\x5a\x33\xb1\x80\xb4\xa0\xf0\x84\x4b\x0f\x70\xbf\x6a\x9e\x16\x74\x0e\x4c\x94\x95\x69\x56\x3b\x21\x8b\x6e\x71\x41\x04\xcb\x50\x9b\x64\x49\x0a\x0e\x19\xe3\x38\x80\xee\x30\x70\x12\x90\xca\xe4\x52\xb1\xbf\x1b\x00\x78\xf8\x74\x7e\xff\xe0\x41\x76\xad\x74\x42\xfe\xc8\x99\x7e\x6a\x36\xef\xe1\xe3\xd5\xed\x9d\x0f\xef\xcd\x32\x27\x98\x20\x05\xea\x92\xa4\x3e\x7f\x37\xcf\x87\xb4\xfc\x79\x7e\x73\x7b\x71\x75\x19\x21\x67\xbd\xd2\x1d\xe5\x6e\x67\x9b\x4d\x05\x21\x0d\x64\xb2\x12\x14\x88\x81\x92\x98\x1c\xea\x3a\x29\x95\xfc\x8e\xa9\xb9\x26\x26\x5f\xad\x92\xbf\x84\x2f\x56\x13\x90\x82\x07\xaf\xae\x9b\x70\xaf\x56\xff\xb3\x9f\xec\x87\x06\x3a\x81\xe1\x05\xfb\x73\xea\xa7\x68\x89\x88\x14\xd3\x2f\xb1\x4c\x8e\x2d\xde\xd7\xba\x4e\xec\x8a\x16\xed\x5b\x6c\xbc\xc6\xe0\x39\xe5\xfd\xd1\x6f\x43\x7f\x57\x5b\x03\xc8\xa4\x02\x8a\x25\x97\xcb\x02\x85\xf1\xcb\x89\xb7\x1f\x4d\x5f\x89\x5d\x05\xbc\x46\x70\x4a\xb0\x5b\xa6\x2a\x61\x58\xb1\xde\x4e\x5d\x95\xa5\x54\x06\x29\x3c\x2e\xe1\xaa\x44\xd1\xde\xcb\x92\x13\x93\x49\x55\xf8\xc5\x4c\xc3\x72\x27\x0d\xfd\xd4\x8a\x87\x9c\x68\x48\x73\xa9\x51\x00\x81\x92\x28\xc3\xd2\x8a\x13\xb5\x26\xb2\x9e\x5a\x62\x92\x5a\x19\x7e\x71\xbb\x20\xba\x83\x27\x36\xce\xf5\xa6\x66\x59\xe2\x29\x68\x34\x60\x24\x08\x49\xf1\xbb\xf6\x05\x2e\xd2\xda\x49\x7d\x67\xd5\x55\x26\x47\x61\x58\xda\x16\x83\x27\x5c\xf6\x7b\x9e\x4a\x91\xb1\x45\xa5\x90\xfa\x77\x63\x0c\x82\x57\xc2\xba\x38\x8f\x24\x0e\xdb\x79\xe9\xd6\x75\x66\x2c\xdf\x80\x61\x94\x7f\x75\x9d\x90\x92\xd9\x6f\xab\xd5\x29\x64\x4a\x16\xdd\x4f\x5a\x56\x2a\xc5\x50\x3a\x9e\x04\x15\x8c\x7b\x1f\x2b\x8d\x66\x0b\xa0\x32\x79\x9c\x98\x68\x88\xb8\x50\xd4\x75\xb2\xfe\xbe\xed\xd1\xfa\xc7\x38\x55\xd3\x31\x9d\x32\x7f\x23\x8c\x23\xb5\x37\x69\x81\x6d\x65\x78\x73\xe1\x74\x0b\x6b\xd3\xd2\x97\x26\x2d\x69\x54\xcf\x2c\xc5\x77\x96\x09\x95\x0a\x29\xde\x1b\xbc\x53\xfc\xad\x21\xaa\xc9\x02\x95\x28\x88\xd2\x39\xe1\x5b\xc9\x93\x89\x4c\xb6\xd0\x5c\xa6\x84\xc3\x33\xe1\x15\x6a\xbf\xd4\x89\x60\x9e\xa4\x17\x82\x60\xc2\xa0\x12\x18\xaa\x57\xd1\xf6\x4e\xfa\xff\xaf\x0b\x1a\xa4\xb2\x28\x39\xda\xed\xd6\x55\x9a\xa2\xd6\x59\xc5\xf9\xd2\xcf\x1c\x65\xea\x24\xfd\x20\x0d\xa0\x52\x52\x41\xaa\x90\x18\x5b\x70\x4b\x92\x3e\x91\x05\xc2\x0f\x66\xf2\xee\x59\x81\x5a\x93\xc5\x56\x70\x81\x08\xda\xdb\x49\xda\x3e\xb0\x1f\x42\xa7\xea\x20\x54\xb1\x4e\xb5\xf5\xee\x5f\xec\xd3\xe6\xa6\xbd\xe7\xcc\x06\xfa\xdc\x9a\x7b\x64\x79\x16\x3b\x81\x2f\xc4\x33\xe1\x8c\x76\x93\xa2\xcc\xe0\x7e\x60\x04\x0c\x18\x44\x47\xa3\x64\xc7\x08\xc5\x6e\x34\xb1\xce\xa8\x8a\x1f\xe5\xb2\xec\xc8\x33\xe0\x8e\x46\xd3\xb0\x34\xd9\xde\x10\x53\x69\x1b\xda\x03\xfb\x76\x10\xd2\xd8\xb8\x19\xc5\x16\x0b\x54\xc7\x08\xdd\xee\x54\x63\x9d\xca\x10\xe9\x31\x3d\xdb\x91\x6f\x6c\x6d\x7a\x64\x82\xda\xef\x47\xcc\xe7\xbb\x53\x0e\x55\x7d\x99\xd9\xb9\x1a\x05\x45\x91\x2e\xad\x29\xc5\xf2\x92\x14\xb8\x5a\x01\x65\xb4\x6b\xee\xdb\xea\x6e\x8b\xfb\xba\xb6\xc3\x4d\x25\xe0\x61\x33\x03\xf6\xb3\xf1\x83\xed\x8b\x14\x16\xf2\x19\xdb\x51\x90\x70\xbe\xec\x46\x77\xa4\x40\xb4\x46\x13\x68\xad\x7e\x05\x65\x81\x2d\xdb\xaa\xea\x75\x9d\xc8\xca\x94\x95\x59\xad\x20\x49\x92\xa0\x3f\x01\xb3\x01\xb2\x26\x2d\x8d\xa5\x72\x1a\x0d\x10\xbd\xb8\x53\x63\x09\x83\xc6\x03\xc4\xfd\x51\x1f\xcb\xe9\xb3\x8b\xa4\xeb\x6f\xd6\x54\x5a\x9f\xfd\x00\xfd\xcb\x13\x3d\x8a\x39\x60\xea\x6e\xde\x3e\x25\xf0\x9e\x88\x14\x39\xef\xcc\x07\x5f\x7e\x05\x4d\x06\x48\xac\x41\xdc\x2b\xb6\xb0\x8d\x67\x4a\xda\x2c\x0a\xdf\xfe\xc0\xa0\x34\x02\x62\x28\x6b\xb6\x69\x64\xc2\xa4\xe4\x33\xfc\x75\xbd\xee\x11\xde\x1e\xc1\x4d\x3e\xf6\x1f\xdf\x58\xeb\x61\xff\xa7\x96\x83\xb8\xad\xd9\x01\x7d\xa0\x8d\xa0\xc8\xf1\x38\xd3\xe0\xfe\x98\x62\x5d\x3a\x70\xa7\xbe\x2f\x9e\xff\xa6\x90\x57\xfb\x79\xf8\x29\x64\x8f\x54\x63\x9d\x3a\xf0\x14\xb2\x6f\xbe\x58\xf7\x0e\xff\x86\x6c\x8f\x54\x4e\xa7\xbe\x9c\xdd\x5c\x5e\x5c\x7e\x78\x07\x77\x39\xc2\xbc\x7d\xdf\x3c\x87\xfb\xb3\xcf\xbf\xb7\xef\xce\xc5\x7a\x4b\x51\x18\x66\x9a\xb7\xe9\x14\x4b\x85\x29\x31\x48\x13\xb8\xe6\x48\x34\x42\xa5\x11\xe6\x76\xd3\xe7\xc0\x84\x36\x48\x6c\xef\x0e\x14\x75\xaa\xd8\x23\x52\x8b\xa3\x4b\x4c\x59\xd6\xfd\x0f\x13\x98\x36\x7e\xa6\xa2\xd8\xb8\xf7\x8d\xe6\x11\xe3\xbf\x3f\x4a\x5f\x1f\xe5\x6a\x64\x73\xa2\xe1\x11\x51\xbc\xe8\x8d\xd6\xf3\x59\xb0\xbb\x9a\x06\xe7\x14\x77\x1d\xe8\xef\xa7\x2a\xdc\x0d\x33\x28\x73\x9f\xf2\xf6\x20\xeb\xae\xbb\x2e\xfb\x90\x35\x0d\x2b\x28\xeb\xcd\x7c\xba\xab\xb6\x09\x80\x4e\x81\x37\xaf\x27\xf6\xa9\xc2\x26\x00\x39\x05\x9d\xbd\x7d\x5b\x31\x55\xd2\x24\xa8\x88\xc4\xb1\xe9\xd0\xdd\x70\xfd\x60\x19\x9f\x3b\xc6\x21\x0e\x5d\x80\xae\x74\x6c\xa3\x75\xd3\x5f\xd4\x99\x0f\x9b\x07\x4f\x91\xed\x22\x47\xd1\x46\x18\x0e\x9c\x92\x76\xe4\x18\x45\x1a\x69\x6c\x89\x67\xb3\xd9\xb7\xd9\x3f\x01\x00\x00\xff\xff\x8d\x08\xa1\xf4\xea\x28\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesEn_usAllJson,
        "wski18n/resources/en_US.all.json",
    )
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
    bytes, err := wski18nResourcesEn_usAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 10474, mode: os.FileMode(420), modTime: time.Unix(1505853001, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesEs_esAllJson,
        "wski18n/resources/es_ES.all.json",
    )
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
    bytes, err := wski18nResourcesEs_esAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x4a\xaa\x2c\x48\x2c\x2e\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x55\x48\xce\x48\x4d\xce\xce\xcc\x4b\x57\xd2\x81\x28\x2c\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\x03\xe9\x08\xce\xcf\x4d\x55\x40\x12\x53\xc8\xcc\x53\x70\x2b\x4a\xcd\x4b\xce\x50\xe2\x52\x50\xa8\xe5\x8a\xe5\x02\x04\x00\x00\xff\xff\x45\xa4\xe9\x62\x65\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesFr_frAllJson,
        "wski18n/resources/fr_FR.all.json",
    )
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
    bytes, err := wski18nResourcesFr_frAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 101, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesIt_itAllJson,
        "wski18n/resources/it_IT.all.json",
    )
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
    bytes, err := wski18nResourcesIt_itAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesJa_jaAllJson,
        "wski18n/resources/ja_JA.all.json",
    )
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
    bytes, err := wski18nResourcesJa_jaAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesKo_krAllJson,
        "wski18n/resources/ko_KR.all.json",
    )
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
    bytes, err := wski18nResourcesKo_krAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesPt_brAllJson,
        "wski18n/resources/pt_BR.all.json",
    )
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
    bytes, err := wski18nResourcesPt_brAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesZh_hansAllJson,
        "wski18n/resources/zh_Hans.all.json",
    )
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
    bytes, err := wski18nResourcesZh_hansAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesZh_hantAllJson,
        "wski18n/resources/zh_Hant.all.json",
    )
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
    bytes, err := wski18nResourcesZh_hantAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1500653295, 0)}
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
    "wski18n/resources/de_DE.all.json": wski18nResourcesDe_deAllJson,
    "wski18n/resources/en_US.all.json": wski18nResourcesEn_usAllJson,
    "wski18n/resources/es_ES.all.json": wski18nResourcesEs_esAllJson,
    "wski18n/resources/fr_FR.all.json": wski18nResourcesFr_frAllJson,
    "wski18n/resources/it_IT.all.json": wski18nResourcesIt_itAllJson,
    "wski18n/resources/ja_JA.all.json": wski18nResourcesJa_jaAllJson,
    "wski18n/resources/ko_KR.all.json": wski18nResourcesKo_krAllJson,
    "wski18n/resources/pt_BR.all.json": wski18nResourcesPt_brAllJson,
    "wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
    "wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
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
    "wski18n": &bintree{nil, map[string]*bintree{
        "resources": &bintree{nil, map[string]*bintree{
            "de_DE.all.json": &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
            "en_US.all.json": &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
            "es_ES.all.json": &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
            "fr_FR.all.json": &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
            "it_IT.all.json": &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
            "ja_JA.all.json": &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
            "ko_KR.all.json": &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
            "pt_BR.all.json": &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
            "zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
            "zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
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

