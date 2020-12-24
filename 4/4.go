package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	countValidPassports := 0

	// Read passport block
	fileEnded := false

	for {
		passport := make(map[string]string)

		// Read passport block lines
		for {
			var kvPair string
			// I tried _, err := fmt.Fscanln(input, &kvPair);
			// but it returned string with missing 1st char
			// (the one that should be at index 0)
			_, err := fmt.Fscanf(input, "%s", &kvPair)

			if err != io.EOF {
				// Get all k:v pairs from line
				if kvPair != "" {
					// I tried fmt.Sscanf(kvPair, "%s:%s", &k, &v);
					// but it just returned k = kvPair and v = "" :(
					kv := strings.Split(kvPair, ":")
					passport[kv[0]] = kv[1]
				} else {
					break
				}
			} else {
				fileEnded = true
				break
			}
		}

		// Checks definition
		requiredKeys := map[string]func(string) bool{
			"byr": func(v string) bool {
				r := regexp.MustCompile(`^\d{4}$`)
				if r.MatchString(v) {
					return v >= "1920" && v <= "2002"
				}
				return false
			},
			"iyr": func(v string) bool {
				r := regexp.MustCompile(`^\d{4}$`)
				if r.MatchString(v) {
					return v >= "2010" && v <= "2020"
				}
				return false
			},
			"eyr": func(v string) bool {
				r := regexp.MustCompile(`^\d{4}$`)
				if r.MatchString(v) {
					return v >= "2020" && v <= "2030"
				}
				return false
			},
			"hgt": func(v string) bool {
				r := regexp.MustCompile(`^(?P<v>\d+)(?P<type>cm|in)$`)
				if r.MatchString(v) {
					res := r.FindStringSubmatch(v)
					if res[2] == "cm" {
						return res[1] >= "150" && res[1] <= "193"
					}
					return res[1] >= "59" && res[1] <= "76"
				}
				return false
			},
			"hcl": func(v string) bool {
				r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
				return r.MatchString(v)
			},
			"ecl": func(v string) bool {
				r := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
				return r.MatchString(v)
			},
			"pid": func(v string) bool {
				r := regexp.MustCompile(`^\d{9}$`)
				return r.MatchString(v)
			},
		}

		valid := true
		for k, checkFunction := range requiredKeys {
			e, keyInMap := passport[k]
			if !keyInMap || !checkFunction(e) {
				valid = false
				break
			}
		}

		if valid {
			countValidPassports++
		}

		if fileEnded {
			break
		}
	}

	fmt.Println(countValidPassports)
}
