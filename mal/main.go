package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func dirList(path string) []string {
	fileList := make([]string, 0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

func main() {
	var rootCmd = &cobra.Command{
		Use:           "mal",
		Short:         "test",
		Long:          `wip test :)`,
		SilenceErrors: true,
	}

	var eFunc = &cobra.Command{
		Use:   "e",
		Short: "Encrypt Dir",
		Args:  cobra.RangeArgs(1, 2),
		Run:   demoEncrypt,
	}
	var cFunc = &cobra.Command{
		Use:   "c",
		Short: "Create Dir and test files",
		Args:  cobra.RangeArgs(1, 2),
		Run:   demoCreate,
	}
	rootCmd.AddCommand(eFunc)
	rootCmd.AddCommand(cFunc)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func demoEncrypt(cmd *cobra.Command, args []string) {
	for _, i := range dirList(args[0]) {
		src, err := os.Open(i)
		baseDir := path.Dir(i)
		if err != nil {
			continue
		}
		defer src.Close()
		fileName := strings.TrimSuffix(path.Base(i), filepath.Ext(i))
		dst, err := os.Create(fmt.Sprintf("%s/%s.%s", baseDir, fileName, args[1]))
		if err != nil {
			continue
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			continue
		}
		os.Remove(i)
	}
}

func demoCreate(cmd *cobra.Command, args []string) {
	copies := 10
	var err error
	if len(args) > 1 {
		copies, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Second args was not a number!")
		}
	}
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		os.Mkdir(args[0], os.ModePerm)
	}
	for i := 0; i < copies; i++ {
		os.Create(fmt.Sprintf("%s/test%d.txt", args[0], i))
	}
}
