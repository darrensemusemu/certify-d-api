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

	"H4sIAAAAAAAC/8RWW2/buBL+KwOeA5wWkC05SZMcvaW3TXdbNMgFfcjmgZbGEhuK1JJDJ97A/30xlKzY",
	"TlLvAgX2IYgoUzPf9831QRS2aa1BQ17kD8Khb63xGA8fnLOOHwprCA3xo2xbrQpJypr0u7eG3/mixkby",
	"038dzkQu/pM+Wk27X3165uxUY/MeSSotlstlIkr0hVMtGxN55w5aZwv0XpkKHP4R0JNYJuLMerpQlblq",
	"rzy68x4lu9y0wb+C8uBVZbCE0Ar200Pg66coNdVfbyOXslT8ndRnzrboSDHvmdQeE9GuvXoQniSF+NTI",
	"+89oKqpFvr+XiFYSoWPXX29FImjRosiFJ6dMFX07lOVXoxciJxdwuGGn37HoqG3osgvWJt0T6L+G7nMP",
	"nV14df7xHRwdZ0evRSLwXjat7tXqvIjzTlto0HtZITTBExhLMEXQ0lXogGppYJLBl7ciEXUX2Zqo9Xma",
	"3t3djaeoKzUbTzH1LRZqtkoLh55S2apRFVSJadsh9GkrF9rK8tLaz+xgXFOjRSKU8SRNwaIEZ/IQVJlP",
	"9vbx4M3h0QiP/z8dTfbK/ZE8eHM4Otg7PJwcTI4OsiwTyRCVg8l+IkgRcxRnnRu4tBaio8ewsP0ezohf",
	"5R2FfAsZR2Uz/OUQnG3969BIM+Igy6lGwPtWSxN1gJUqQBaoVh5sUQTn0BQIdgZUI/RgnibOSvAnDg3I",
	"qbc6EMLV+SeOESVwV6OBEh3OMJovE7Y8VyX6bYClLUKDhjqIM+vWcQCjgFc4rsYQYgmeXn75/HosEjGz",
	"rpEUNVTPwX2M4i7IA8oIHlSJhtRMoY9IBtFeFGsMnwgauQDr4j9O2oVCXcIsOKq5/k2Hlhmq2YYwf4fK",
	"Y61vErmsEU4vL8+guwCFLREqNOgkYQnTRURpnaqUAY9ujq4X+EexH68X6Jtsn0+FDl7N8Yu8V01oVp1j",
	"wK0M7e+JhFtR9/thliWiUaY7HfCpZ6UMYYWOafX18TSFfW0dgQ9NI91iC13MiDF8c4oIDSgDH0ylla9B",
	"mhKGrGKaaCplEJ2HV8EHqXUXGh8Uq8M3jDVAWNRGFVKzirdYW13yJ2yNb2tbSK3+xPL1hi7iAt1cFQhX",
	"Rs6l0uz0udB1L5jiTAbNWsmpDZRPtTTcnncX03Y+ruuwM3eWT9o7pxMWwSlaXPAQ6trJb06S9W9ROnQn",
	"gUfJg5jG08eV/V+/XT7B233WI6QFeJ6T1oh+wLHfzsojMu7X3azlomA/qzZZcHebLUYlzCciEXN0vvMy",
	"GWcspW3RyFaJXOyPszGnWyupjvDTOg7RVGo1j3JXSE8T612NxS23pliTa3sDz2eOtTIgDQTT7xxqjrGy",
	"WGbuvfHup1Lk4lSaUuMvSBexqE6i22RzV9nLsp+2qQw7wjNLyskmDxeM6XNvyLnnjQ9o026tiqnR1dyW",
	"WaZn0PuYezHNZeVFfi062cUNf7qKAZfgYkcM1rXve9fvIcv2DuH92/7F/3aJfh79/Euib6YSl+U/Saef",
	"GRtZFNjS2nLqX4qPH7ZV9tpaT8/1XoN3EPqdtXAYBwn3wmF/BWug1ZK48bwQo8e9+Gl8Jrspv7Bar7cu",
	"kV8/17Sub5Y360KxFc+Y5cBrTZ14vFl2djmlfDS7qcgFyYrrKRHB8Zr60DpLtrB6mafpQ209LXm1TOeT",
	"tDc/l07xOIiE60HmVf8f2ty4lDx/PTaB/8YlznnCGJ6Y1z++dtPtgxHHpvW4Dq+Zie026V9HrtsE3+Mc",
	"tW15BRtI8vU8TePsYwb5cXacbdBknXsdtw1+mKNbUM0pGWddlJ0xGRknQndc3iz/CgAA//+l5FbC6w0A",
	"AA==",
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
