package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Yaml2Go struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

var genconf = &cobra.Command{
	Use: "genconf",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile(*genConfSrc)
		if err != nil {
			panic(err.Error())
		}

		y2g := Yaml2Go{}
		_, err = resty.New().R().SetResult(&y2g).SetFormData(map[string]string{"schema": string(b)}).Post("https://www.printlove.cn/api/yaml2go")
		if err != nil {
			panic(err.Error())
		}

		data := strings.ReplaceAll(y2g.Data, "AutoGenerated", "MyConf")
		data = "package types\n\n" + data

		_, err = os.Stat(*genConfDest)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err.Error())
			}
		} else {
			err = os.Truncate(*genConfDest, 0)
			if err != nil {
				panic(err.Error())
			}
		}
		f, err := os.OpenFile(*genConfDest, os.O_CREATE|os.O_SYNC|os.O_RDWR, 0666)
		if err != nil {
			panic(err.Error())
		}
		_, err = f.Write([]byte(data))
		if err != nil {
			panic(err.Error())
		}
	},
}
var genConfSrc *string
var genConfDest *string

func main() {
	genConfSrc = genconf.Flags().StringP("src", "s", "", "")
	genConfDest = genconf.Flags().StringP("dest", "d", "", "")
	var rootCmd = &cobra.Command{Use: "gen"}
	rootCmd.AddCommand(genconf)
	rootCmd.Execute()
}
