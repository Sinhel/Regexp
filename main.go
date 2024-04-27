package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

type flags struct {
	appendpath    *bool
	txtinput      *string
	omitmatch     *bool
	printpath     *bool
	reinput       *string
	seperatorchar *string
	wildcard      *string
	charmap       *string
	listmaps      *bool
	decode        *bool
}

func main() {
	args := flags{
		appendpath:    flag.Bool("a", false, "Append file path as seperate column in csv"),
		txtinput:      flag.String("i", "./input.txt", "Specify input file here\n If file is specified, file is parsed\n If directory is specified, directory is parsed recursively\n"),
		omitmatch:     flag.Bool("o", false, "Omit regular expression's full match from output, so only capture groups will be printed"),
		printpath:     flag.Bool("p", false, "Print paths that are read from input flag"),
		reinput:       flag.String("r", "./regex.txt", "Specify file regular expression is read from\n"),
		seperatorchar: flag.String("s", ",", "Specify which character to use as seperator between matches\n"),
		wildcard:      flag.String("w", "*", "Limit what kind of files recursive parsing should go through\n"),
		charmap:       flag.String("c", "Windows 1252", "Charmap to decode from, if program detects non UTF-8 characters\n"),
		listmaps:      flag.Bool("lc", false, "lists all available charmaps"),
		decode:        flag.Bool("d", false, "decode string in format chosen by argument c"),
	}
	flag.Parse()

	if *args.printpath {
		printpaths(recursivepathsearch(args))
		return
	}

	if *args.listmaps {
		listcharmaps()
		return
	}

	//Read regular expression into memory, and store SubexpNames into keys slice
	reinput, err := os.ReadFile(*args.reinput)
	if err != nil {
		log.Fatal(err)
	}

	if *args.decode {
		reinput = decodestring(reinput, args)
	}

	validcheck := utf8.ValidString(string(reinput))
	if !validcheck {
		log.Fatal("regex not UTF-8 compliant, decode with -d")
	}

	keys := printReSubexpNames(string(reinput), args)

	for _, v := range recursivepathsearch(args) {
		if !isDir(v) {
			txtinput, err := os.ReadFile(v)
			if err != nil {
				log.Fatal(err)
			}

			if *args.decode {
				txtinput = decodestring(txtinput, args)
			}

			validcheck := utf8.ValidString(string(txtinput))
			if !validcheck {
				log.Fatal("Input not UTF-8 compliant, decode with -d")
			}

			totres := runRegexAllStringSubmatch(string(txtinput), string(reinput), keys)
			printReSubexContents(totres, keys, v, args)
		}
	}
}

func runRegexAllStringSubmatch(txtinput string, reinput string, keys []string) (totres []map[string]string) {
	re := regexp.MustCompile(reinput)
	for _, match := range re.FindAllStringSubmatch(txtinput, -1) {
		matchres := map[string]string{}
		totres = append(totres, matchres)
		for i, k := range keys {
			matchres[k] = match[i]
		}
	}
	return
}

// print SubexNames, typically at top of file to get columns for csv
func printReSubexpNames(reinput string, args flags) (keys []string) {
	re := regexp.MustCompile(string(reinput))
	names := re.SubexpNames()
	for i, k := range names {
		if *args.omitmatch && i == 0 {
			keys = append(keys, strconv.Itoa(i))
		} else {
			if k != "" {
				fmt.Printf("%s%s", k, *args.seperatorchar)
				keys = append(keys, k)
			} else {
				fmt.Printf("%v%s", i, *args.seperatorchar)
				keys = append(keys, strconv.Itoa(i))
			}
		}
	}
	if *args.appendpath {
		fmt.Printf("filepath")
	}
	fmt.Printf("\n")
	return
}

// print contents of Subexpressions, typically done after printing SubexNames
func printReSubexContents(totres []map[string]string, keys []string, path string, args flags) {
	for _, m := range totres {
		for _, v := range keys {
			if *args.omitmatch && v == "0" {
				// do nothing
			} else {
				fmt.Printf("%s%s", m[v], *args.seperatorchar)
			}
		}
		if *args.appendpath {
			fmt.Printf("%s", path)
		}
		fmt.Printf("\n")
	}
}

func printpaths(paths []string) {
	for i := range paths {
		fmt.Printf("%s",
			paths[i])
		fmt.Printf("\n")
	}
}

func recursivepathsearch(args flags) (paths []string) {
	filepath.WalkDir(*args.txtinput, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %v: %v\n",
				s,
				err)
			return err
		}
		if match(args, s) {
			paths = append(paths, s)
		}
		return nil
	})
	return paths
}

// wildcardToRegexp converts a wildcard pattern to a regular expression pattern.
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

func match(args flags, value string) bool {
	result, _ := regexp.MatchString(wildcardToRegexp(*args.wildcard), value)
	return result
}

// check if path is dir or file
func isDir(path string) (isDir bool) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	isDir = fileinfo.IsDir()
	return
}

func listcharmaps() {
	for _, enc := range charmap.All {
		cmap, ok := enc.(*charmap.Charmap)
		if ok {
			fmt.Printf("%s\n", cmap.String())
		}
	}
}

func decodestring(s []byte, args flags) (decoded []byte) {
	mapping, err := mappingcheck(*args.charmap)
	if err != nil {
		log.Fatal(err)
	}
	decoded, err = mapping.NewDecoder().Bytes(s)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%s", string(decoded))
	return decoded
}

func mappingcheck(s string) (enc charmap.Charmap, err error) {
	for _, enc := range charmap.All {
		cmap, ok := enc.(*charmap.Charmap)
		if ok && cmap.String() == s {
			return *cmap, nil
		}
	}
	return enc, errors.New("encoder not found in charmap list")
}
