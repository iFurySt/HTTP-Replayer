/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "tp [target] [target]",
		Short: "Traffic Replayer is a tool for capturing and replaying network traffic non-intrusively. Supports replaying 7th layer HTTP traffic currently.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, targetArgs []string) error {
			targets := filterTargets(targetArgs)
			if len(targets) == 0 {
				return fmt.Errorf("no valid target URLs")
			}
			httpMethods := filterHttpMethods(argHttpMethods)
			uris := filterUris(argUris)
			headers := filterHeaders(argHeaders)
			rate := filterRate(argRate)
			nic := filterNic(argNic)
			ports := filterPorts(argPorts)

			args := Argument{
				Targets:     targets,
				HttpMethods: httpMethods,
				Uris:        uris,
				Headers:     headers,
				Rate:        rate,
				Nic:         nic,
				Ports:       ports,
			}

			display(args)

			Capture(args)

			return nil
		},
	}

	rootCmd.PersistentFlags().StringSliceVarP(&argHttpMethods, "methods", "M", []string{}, "HTTP methods")
	rootCmd.PersistentFlags().StringSliceVarP(&argUris, "uris", "U", []string{}, "URIs to filter")
	rootCmd.PersistentFlags().StringSliceVarP(&argHeaders, "headers", "H", []string{}, "Headers to filter, format: key=value, key=, =value")
	rootCmd.PersistentFlags().StringVarP(&argRate, "rate", "R", "", "Rate control, format: number/s|number/min|number/h, such as 100/s, 1000/min, 10000/h")
	rootCmd.PersistentFlags().StringVarP(&argNic, "nic", "N", "", "Network interface to capture")
	rootCmd.PersistentFlags().StringSliceVarP(&argPorts, "ports", "P", []string{}, "Ports to filter")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
