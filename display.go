/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"fmt"
	"strings"
)

func getValue(s string) string {
	if s == "" {
		return "No limit"
	}
	return s
}

func display(args Argument) {
	fmt.Println("Traffic Replayer is running :)")

	fmt.Println("\tTarget URLs:", args.Targets)
	fmt.Println("\tHTTP Methods:", getValue(strings.Join(args.HttpMethods, ", ")))
	fmt.Println("\tURIs:", getValue(strings.Join(args.Uris, ", ")))
	headerStr := ""
	for _, header := range args.Headers {
		headerStr += header.Key + "=" + header.Value + ", "
	}
	if headerStr != "" {
		headerStr = headerStr[:len(headerStr)-2]
	}
	fmt.Println("\tHeaders:", getValue(headerStr))
	fmt.Println("\tRate:", getValue(fmt.Sprintf("%d/%s", args.Rate.Number, args.Rate.Unit)))
	fmt.Println("\tNetwork interface:", getValue(args.Nic))
	fmt.Println("\tPorts:", getValue(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args.Ports)), ", "), "[]")))
	fmt.Println()
}
