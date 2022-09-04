// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RWW2/buBL+KwOeA5wWkC05SZMcvaW3k55t0SAX9KGbB1oaS2woUksOnXgD//fFkIpj",
	"O8l6FyiwD0FEWZz55ptvLveisl1vDRryorwXDn1vjcd4+OCcdfxQWUNoiB9l32tVSVLW5D+8NfzOVy12",
	"kp/+7XAmSvGv/NFqnn71+ZmzU43deySptFgul5mo0VdO9WxMlMkd9M5W6L0yDTj8LaAnsczEmfV0oRpz",
	"1V95dOcDSna5aYN/BeXBq8ZgDaEX7GeAwJ+fotTUfr2JsdS14ntSnznboyPFcc+k9piJfu3VvfAkKcSn",
	"Tt59RtNQK8r9vUz0kggdu/56IzJBix5FKTw5ZZro26Gsvxq9ECW5gKsv7PQHVim0DV52wdoM9wSG25Cu",
	"e0h24dX5x3dwdFwcvRaZwDvZ9XpgK3kR54lb6NB72SB0wRMYSzBF0NI16IBaaWBSwJe3IhNtymxL1Psy",
	"z29vb8dT1I2ajaeY+x4rNXuQhUNPuezVqAmqxrxPCH3ey4W2sr609jM7GLfUaZEJZTxJUzEpwZkyBFWX",
	"k719PHhzeDTC4/9OR5O9en8kD94cjg72Dg8nB5Ojg6IoRLbKysFkPxOkiGMUZ8kNXFoL0dFjWtj+AGfE",
	"r8oUQrmFjLOymf56lZxt/tvQSTPiJMupRsC7XksTeYAHVoAsUKs82KoKzqGpEOwMqEUYwDwVzgPhTxwa",
	"kFNvdSCEq/NPnCPK4LZFAzU6nGE0X2dsea5q9NsAa1uFDg0liDPr1nEAo4BXOG7GEGIJnl5++fx6LDIx",
	"s66TFDlUz8F9zOIuyCuUETyoGg2pmUIfkaxIe5GsMXwi6OQCrIv/WLQLhbqGWXDUcv2bhJYjVLMNYv5K",
	"KI+1vhnIZYtwenl5BukDqGyN0KBBJwlrmC4iSutUowx4dHN0A8F/lvvxeoG+Kfb5VOng1Ry/yDvVhe6h",
	"c6xwK0P7eyLjVpR+PyyKTHTKpNMBn4aolCFs0HFYQ308lbBvrSPwoeukW2yhi4oYwzeniNCAMvDBNFr5",
	"FqSpYaUqDhNNowyi8/Aq+CC1TqnxQTE7/IWxBgir1qhKambxBlura77C1vhrbSup1e9Yv97gRVygm6sK",
	"4crIuVSanT6XuvSCQ5zJoJkrObWByqmWhtvz7mLa1uM6Dzu1s3zS3llOWAWnaHHBQyi1k1+cJOvfonTo",
	"TgKPknsxjaePD/b//+3yCd50bUBIC/A8J60Rw4Bjv8nKIzLu12nWclGwn4c2WXF3my1GNcwnIhNzdD55",
	"mYwLptL2aGSvRCn2x8WY5dZLaiP8vI1DNJdazSPdDdJTYb1rsbrh1hRrcm1v4PnMuVYGpIFghp1DzTFW",
	"FtPMvTd++6kWpTiVptb4P6SLWFQn0W22uavsFcVP21RWO8IzS8rJZhwuGDNob6W5542v0OZprYrSSDW3",
	"ZZbDM+h91F6UuWy8KL+LRLu45qsPOeASXOzIwTr3Q+/6NRTF3iG8fzu8+M8u0s+jn3+I9E0pcVn+HTn9",
	"zNzIqsKe1pZT/1J+/GpbZa+99fRc7zV4C2HYWSuHcZBwL1ztr2AN9FoSN54XcvS4Fz/Nz2R3yC+s1uut",
	"S5Tfn2ta36+X1+tEsRXPmOUqrjV24vF6meyypHw0u8nIBcmG6ykTwem1fXPVrca15DHqsQv8N65xzrtm",
	"Pp/k0cEy2zb5Hueobc9Lz4bZMs/jtGmtp/K4OC4GO4KDGkBv2/owR7eglvMfB0uMkSVgZGy/6bi8Xv4R",
	"AAD//0Pmak5YDQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
