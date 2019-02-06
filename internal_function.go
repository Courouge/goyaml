package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func add_double_quote_by_arg(t string) string {
	var t1 string
	var flag int = 0
	for i, r := range t {
		if r == '=' && flag == 0 {
			flag = flag + 1
			t1 = t1 + ": \""
		} else {
			t1 = t1 + string(r)
		}
		if i == len(t)-1 {
			t1 = t1 + "\""
		}
	}
	return t1
}

func delete_double_quote(t string) string {
	var t1 string
	for i, r := range t {
		if r == '"' {

    } else {
      t1 = t1 + string(t[i])
    }
	}
	return t1
}

func clean_whitespace(t string) string {
	var t1 string
	for i, r := range t {
		if i != 0  && i != len(t) {
			if t[i] == ' ' && t[i-1] == '{' {
				t1 = t1 + ""
			} else if t[i] == ' ' && t[i+1] == '}'{
				t1 = t1 + ""
			} else {
				t1 = t1 + string(r)
			}
		}
	}
	return t1
}

func Yamlfiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".yml" {
			return nil
		} else {
			files = append(files, path)
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
	return files
}

func add_whitespace(t string) string {
	var t1 string
	for i, r := range t {
    if i != 0 && i != len(t)-1 {
      if t[i] == '{' && t[i-1] == '{' {
        t1 = t1 + string(' ')
      } else if t[i] == '}' && t[i+1] == '}'  {
        t1 = t1 + string(' ')
			} else {
				t1 = t1 + string(r)
			}
		}
	}
	return t1
}

func moduleSearch(s []string) []string {
    var search []string
    for i := range s {
        var tmp =  "  " + s[i] + ": "
        search = append(search, tmp)
        // fmt.Println(search[i])
    }

    return search
}

func testmodule2(s string, modules []string) bool {
var (
ToBe bool = false
)
//   file:
    for i := range modules {
        if (strings.HasPrefix(s, string(modules[i])) == true) {
            ToBe = true
            // println(string(modules[i]))
        }
    }
    return ToBe
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}