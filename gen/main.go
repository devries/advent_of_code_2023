package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	"aoc/utils"
)

type DayDirectory struct {
	Name string
	Day  int
	Part string
}

func newDayDirectory(name string) DayDirectory {
	res := DayDirectory{Name: name}

	fmt.Sscanf(name, "day%dp%s", &(res.Day), &(res.Part))

	return res
}

type TemplateData struct {
	DayDirectories   []DayDirectory
	CurrentDirectory DayDirectory
}

type ByName []DayDirectory

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func main() {
	files, err := os.ReadDir(".")
	utils.Check(err, "unable to read directory")

	td := TemplateData{}

	latest := time.Time{}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), "day") && f.IsDir() {
			found := false
			sub, err := os.ReadDir(f.Name())
			utils.Check(err, "unable to read subdirectory %s", f.Name())

			for _, sf := range sub {
				switch sf.Name() {
				case "solution.go":
					t, err := getModTime(sf)
					utils.Check(err, "unable to get mod time for %s/%s", f.Name(), sf.Name())
					if t.After(latest) {
						latest = t
						td.CurrentDirectory = newDayDirectory(f.Name())
					}
					found = true
				}
			}

			if found {
				td.DayDirectories = append(td.DayDirectories, newDayDirectory(f.Name()))
			}
		}
	}

	sort.Sort(ByName(td.DayDirectories))

	tpls := template.Must(template.ParseFS(os.DirFS("templates"), "*.tmpl"))

	fout, err := os.Create("run.go")
	utils.Check(err, "unable to create run.go file")
	defer fout.Close()

	err = tpls.ExecuteTemplate(fout, "run.tmpl", td)
	utils.Check(err, "unable to execute template")
}

func getModTime(de os.DirEntry) (time.Time, error) {
	fi, err := de.Info()
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to get file information: %w", err)
	}

	return fi.ModTime(), nil
}
