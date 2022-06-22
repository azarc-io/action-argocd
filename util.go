package main

import "io/ioutil"

func readFile(p string) (string, error) {
	dat, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func boolRef(b bool) *bool {
	return &b
}

func stringRef(s string) *string {
	return &s
}

func intRef(s int) *int {
	return &s
}

func int64Ref(s int64) *int64 {
	return &s
}
