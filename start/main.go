package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
	"time"

	"aoc/utils"
	"github.com/spf13/pflag"
)

func main() {
	now := time.Now()
	day := pflag.IntP("day", "d", now.Day(), "Advent of Code Day")

	pflag.Parse()

	dirname := fmt.Sprintf("day%02dp1", *day)

	if err := os.MkdirAll(dirname, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating destination directory %s: %s\n", dirname, err)
		os.Exit(1)
	}

	tpls := template.Must(template.ParseFS(os.DirFS("templates"), "*.tmpl"))

	partfilename := "solution.go"
	f, err := os.Create(path.Join(dirname, partfilename))
	utils.Check(err, "unable to create file %s", partfilename)
	defer f.Close()

	err = tpls.ExecuteTemplate(f, "part.tmpl", *day)
	utils.Check(err, "unable to execute template part.tmpl")

	testfilename := "solution_test.go"
	ft, err := os.Create(path.Join(dirname, testfilename))
	utils.Check(err, "unable to create file %s", testfilename)
	defer ft.Close()

	err = tpls.ExecuteTemplate(ft, "part_test.tmpl", *day)
	utils.Check(err, "unable to execute template part_test.tmpl")
}
