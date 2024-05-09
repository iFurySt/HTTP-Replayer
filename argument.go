/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import "net/url"

var (
	argHttpMethods []string
	argUris        []string
	argHeaders     []string
	argRate        string
	argNic         string
	argPorts       []string
)

type Argument struct {
	Targets     []*url.URL
	HttpMethods []string
	Uris        []string
	Headers     []Header
	Rate        *Rate
	Nic         string
	Ports       []int
}

type HeaderType int

const (
	HeaderTypeKey HeaderType = iota + 1
	HeaderTypeValue
	HeaderTypeBoth
)

type Header struct {
	Key   string
	Value string
	Type  HeaderType
}

type RateUnit string

const (
	RateUnitSecond RateUnit = "s"
	RateUnitMinute RateUnit = "m"
	RateUnitHour   RateUnit = "h"
)

type Rate struct {
	Number int
	Unit   RateUnit
}
