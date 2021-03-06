// Copyright (c) 2015 Andrea Masi. All rights reserved.
// Use of this source code is governed by MIT license
// which that can be found in the LICENSE.txt file.

package conf

import (
	"fmt"
	"reflect"
	"testing"
)

type exampleStruct struct {
	Name    string
	Surname string
}
type iniParseCase struct {
	Path           string
	ExpectedResult interface{}
}

var iniParseTestCases []iniParseCase

func TestParseINI(t *testing.T) {
	for _, testCase := range iniParseTestCases {
		s := exampleStruct{}
		v, _ := getStructValue(&s)
		parseINI(testCase.Path, v)
		isEqual := reflect.DeepEqual(testCase.ExpectedResult, s)
		if !isEqual {
			if testing.Verbose() {
				fmt.Println("Expect:", testCase.ExpectedResult)
				fmt.Println("Got:", s)
			}
			t.Fail()
		}
	}

}

func init() {
	// TestParseINI
	iniTestCase := iniParseCase{
		"test_data/one.ini",
		exampleStruct{"Zaphod", "Beeblebrox"},
	}
	iniParseTestCases = append(iniParseTestCases, iniTestCase)

}
