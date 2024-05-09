/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func filterTargets(ss []string) []*url.URL {
	res := make([]*url.URL, 0, len(ss))
	for _, s := range ss {
		u, err := url.Parse(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		res = append(res, u)
	}
	return res
}

func filterHttpMethods(ss []string) []string {
	res := make([]string, 0, len(ss))
	for _, v := range ss {
		s := strings.TrimSpace(v)
		switch s {
		case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
			http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
			res = append(res, s)
		}
	}
	return res
}

func filterUris(ss []string) []string {
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		u, err := url.Parse(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		res = append(res, u.String())
	}
	return res
}

func filterHeaders(ss []string) []Header {
	res := make([]Header, 0, len(ss))
	for _, s := range ss {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) != 2 {
			continue
		}

		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		if k == "" && v == "" {
			continue
		}

		t := HeaderTypeBoth
		if k == "" {
			t = HeaderTypeValue
		}
		if v == "" {
			t = HeaderTypeKey
		}

		res = append(res, Header{
			Key:   k,
			Value: v,
			Type:  t,
		})
	}
	return res
}

// rootCmd.PersistentFlags().StringVarP(&rate, "rate", "R", "", "Rate control, format: number/s|number/min|number/h, such as 100/s, 1000/min, 10000/h")
func filterRate(s string) *Rate {
	ss := strings.Split(strings.TrimSpace(s), "/")
	if len(ss) != 2 {
		return nil
	}

	num, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil
	}

	var unit RateUnit
	switch strings.TrimSpace(ss[1]) {
	case string(RateUnitSecond):
		unit = RateUnitSecond
	case string(RateUnitMinute):
		unit = RateUnitMinute
	case string(RateUnitHour):
		unit = RateUnitHour
	default:
		return nil
	}

	return &Rate{
		Number: num,
		Unit:   unit,
	}
}

func isNicExists(nicName string) bool {
	interfaces, err := net.Interfaces()
	if err != nil {
		return false
	}

	for _, i := range interfaces {
		if i.Name == nicName {
			return true
		}
	}

	return false
}

func filterNic(s string) string {
	s = strings.TrimSpace(s)
	if isNicExists(s) {
		return s
	}
	return ""
}

func filterPorts(ss []string) []int {
	res := make([]int, 0, len(ss))
	for _, v := range ss {
		port, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		if port < 0 || port > 65535 {
			continue
		}
		res = append(res, port)
	}
	return res
}
