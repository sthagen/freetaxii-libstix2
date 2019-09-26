// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
NameProperty - A property used by one or more STIX objects that captures a
vanity name for the STIX object in string format.
*/
type NameProperty struct {
	Name string `json:"name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - NameProperty
// ----------------------------------------------------------------------

/*
SetName - This method takes in a string value representing a name of the object
and updates the name property.
*/
func (o *NameProperty) SetName(s string) error {
	o.Name = s
	return nil
}

/*
GetName - This method returns the current name of the object.
*/
func (o *NameProperty) GetName() string {
	return o.Name
}

/*
CompareNameProperties - This function will compare two name properties (object 1
and object 2) to make sure they are the same. This function will return an
integer that tracks the number of problems and a slice of strings that contain
the detailed results, whether good or bad.
*/
func CompareNameProperties(obj1, obj2 *NameProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Name Value
	if obj1.Name != obj2.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", obj1.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", obj1.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
