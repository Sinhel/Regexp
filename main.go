package main

import (
	"log"
	"os"
	"regexp"
	/* 	"fmt"
		"io/fs"
	   	"strings"
	   	"path/filepath" */)

func main() {
	//Set default value for input files as input.txt, regex.txt in same the folder as the program. Files can also be specified as arguments in cli
	txtinputarg := "./input.txt"
	reinputarg := "./regex.txt"
	if len(os.Args) >= 2 {
		txtinputarg = os.Args[1]
		reinputarg = os.Args[2]
	}

	//Read files into memory
	reinput, err := os.ReadFile(reinputarg)
	if err != nil {
		log.Fatal(err)
	}
	txtinput, err := os.ReadFile(txtinputarg)
	if err != nil {
		log.Fatal(err)
	}

	regexonfile(string(txtinput), string(reinput))
}

func regexonfile(txtinput string, reinput string) (result []map[string]string, subexpNames []string) {
	re := regexp.MustCompile(string(reinput))
	result = make([]map[string]string, 0)
	for _, match := range re.FindAllStringSubmatch((string(txtinput)), -1) {
		res := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				res[name] = match[i]
			}
		}
		result = append(result, res)
		subexpNames = re.SubexpNames()
	}
	return
}

/* func printSubexpNames(subexpNames []map[string]string) {
	for i, name := range subexpNames {
		if i != 0 && name != "" {
			fmt.Printf("%v,",
				name)
		}
	}
	fmt.Print("\n")
	for i, match := range subexpNames {
		for i, name := range subexpNames {
			if i != 0 && name != "" {
				fmt.Printf("%v,",
					match[name])
			}
		}
		fmt.Print("\n")
		i += i
	}
}*/

// test funtcion to print paths to console.
// TODO: call with flag
/* func printpaths(paths []string) {
	for i := range paths {
		fmt.Printf(paths[i])
		fmt.Printf("\n")
	}
}

func recursivepathsearch(ipath string, matchpattern string) []string {
	opath := []string{}
	filepath.WalkDir(ipath,
		func(s string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %v: %v\n",
					s,
					err)
				return err
			}
			if match(matchpattern, s) {
				opath = append(opath, s)
			}
			return nil
		})
	return opath
}

// wildCardToRegexp converts a wildcard pattern to a regular expression pattern.
func wildCardToRegexp(pattern string) string {
	components := strings.Split(pattern, "*")
	if len(components) == 1 {
		// if len is 1, there are no *'s, return exact match pattern
		return "^" + pattern + "$"
	}
	var result strings.Builder
	for i, literal := range components {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return "^" + result.String() + "$"
}

func match(pattern string, value string) bool {
	result, _ := regexp.MatchString(wildCardToRegexp(pattern), value)
	return result
} */
