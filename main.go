package main

import (
	"bufio"
	"os"
	"strings"
	"log"
	//"fmt"
	//"io/ioutil"
	//"regexp"
)

func main() {

  var ymlfiles, finalfile []string

  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
/*
    data, err := ioutil.ReadFile(dir + "/modules.txt")
    r := regexp.MustCompile( ".*_module.html")
    fmt.Println(r)
    match := r.FindAllString(string(data), -1)
    fmt.Println(match[0])
    var filter []string
    for i := range match {

        fmt.Println(string(match[i]))
        strings.Replace(match, "=", "", -1)

    }
*/

  // configure your ansible roles path
  ymlfiles = Yamlfiles(dir + "/roles")
  // println(add_double_quote_by_arg(delete_double_quote("pkg=\"{{ foo }}/trololo\"")))
  for _, y  := range ymlfiles {
    finalfile = nil
    fileHandle, _ := os.Open(y)
    fileScanner := bufio.NewScanner(fileHandle)
    for fileScanner.Scan() {

      words :=  strings.Fields( clean_whitespace(fileScanner.Text()))

      if testmodule(fileScanner.Text()) == true {
          for i, v := range words {
            if i == 0 && words[0] != fileScanner.Text() {
                v =   "  "  + v + string('\n')
                println(v)
                finalfile = append(finalfile,v)
              } else {
                // condition sur egal
                //if strings.Count(v, "dest=") == 1 {
                //  v = strings.Replace(v, "=", ": ", -1)
                //}
                v = strings.Replace(v, "=", ": ", -1)
                v = strings.Replace(v, "{{", "{{ ", -1)
                v = strings.Replace(v, "}}", " }}", -1)
                v =  "  " + "  " + v + string('\n')
                if len(words) == i {
                    v = v + "sdsds"

                }
                finalfile = append(finalfile, v)
              }
          }

        } else {
          finalfile = append(finalfile, fileScanner.Text() + string('\n'))
        }

    }

    defer fileHandle.Close()
    // fmt.Printf("%v", finalfile)

    if err := WriteStringToFile(y , strings.Join(finalfile,"")); err != nil {
      panic(err)
    }
  }
}