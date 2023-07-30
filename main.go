package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	//Set default value for input files as input.txt, regex.txt in same the folder as the program. Files can also be specified as arguments in cli
	/* 	txtinputarg := "./input.txt" */
	reinputarg := "./regex.txt"
	if len(os.Args) >= 2 {
		/* 		txtinputarg = os.Args[1] */
		reinputarg = os.Args[2]
	}

	//Read regular expression into memory, and store SubexpNames into keys slice
	reinput, err := os.ReadFile(reinputarg)
	if err != nil {
		log.Fatal(err)
	}
	keys := printReSubexpNames(string(reinput))

	for _, v := range recursivepathsearch(".", "*.txt") {
		txtinput, err := os.ReadFile(v)
		if err != nil {
			log.Fatal(err)
		}
		totres := runRegexAllStringSubmatch((string(txtinput)), string(reinput))
		printReSubexContents(totres, keys)
		time.Sleep(1 * time.Microsecond)
	}
}

func runRegexAllStringSubmatch(txtinput string, reinput string) (totres []map[string]string) {
	re := regexp.MustCompile(reinput)
	for _, match := range re.FindAllStringSubmatch(txtinput, -1) {
		matchres := map[string]string{}
		totres = append(totres, matchres)
		for i, subexpName := range re.SubexpNames() {
			if subexpName != "" {
				matchres[subexpName] = match[i]
			}
		}
	}
	return
}

// print SubexNames, typically at top of file to get collums for csv
func printReSubexpNames(reinput string) (keys []string) {
	re := regexp.MustCompile(string(reinput))
	keys = re.SubexpNames()
	for _, k := range keys {
		if k != "" {
			fmt.Printf("%s,", k)
		}
	}
	return
}

// print contents of Subexpressions, typically done after printing SubexNames
func printReSubexContents(totres []map[string]string, keys []string) {
	fmt.Printf("\n")
	for _, m := range totres {
		for _, v := range keys {
			if v != "" {
				fmt.Printf("%s,", m[v])
			}
		}
		fmt.Printf("\n")
	}
}

// test funtcion to print paths to console.
// TODO: call with flag
func printpaths(paths []string) {
	for i := range paths {
		fmt.Printf("%s",
			paths[i])
		fmt.Printf("\n")
	}
}

// TODO: call with flag
func recursivepathsearch(ipath string, matchpattern string) (opath []string) {
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

// wildcardToRegexp converts a wildcard pattern to a regular expression pattern.
// TODO: call with flag
func wildcardToRegexp(pattern string) string {
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
	result, _ := regexp.MatchString(wildcardToRegexp(pattern), value)
	return result
}
