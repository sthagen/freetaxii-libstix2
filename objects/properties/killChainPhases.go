// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
KillChainPhasesProperty - A property used by one or more STIX objects that
captures a list of kll chain phases as defined by STIX.
*/
type KillChainPhasesProperty struct {
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}

/*
KillChainPhase - This type defines all of the properties associated with the
STIX Kill Chain Phase type.
*/
type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name,omitempty"`
	PhaseName     string `json:"phase_name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesProperty - Setters
// ----------------------------------------------------------------------

/*
CreateKillChainPhase - This method takes in two parameters and creates and
adds a new kill chain phase to the list. The first value is a string value
representing the name of the kill chain being used. The second value is a string
value representing the phase name from that kill chain.
*/
func (o *KillChainPhasesProperty) CreateKillChainPhase(name, phase string) error {
	k, _ := o.newKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
	return nil
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhasesProperty - Setters
// ----------------------------------------------------------------------

/*
newKillChainPhase - This method returns a reference to a slice location. This
will enable the code to update an object located at that slice location.
*/
func (o *KillChainPhasesProperty) newKillChainPhase() (*KillChainPhase, error) {
	var s KillChainPhase

	// if o.KillChainPhases == nil {
	// 	a := make([]KillChainPhase, 0)
	// 	o.KillChainPhases = a
	// }

	positionThatAppendWillUse := len(o.KillChainPhases)
	o.KillChainPhases = append(o.KillChainPhases, s)
	return &o.KillChainPhases[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhase - Setters
// ----------------------------------------------------------------------

/*
SetName - This method takes in a string value representing the name of a kill
chain and updates the kill chain name property.
*/
func (o *KillChainPhase) SetName(s string) error {
	o.KillChainName = s
	return nil
}

/*
SetPhase - This method takes in a string value representing the phase of a
kill chain and updates the phase name property.
*/
func (o *KillChainPhase) SetPhase(s string) error {
	o.PhaseName = s
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesProperty - Checks
// ----------------------------------------------------------------------

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *KillChainPhasesProperty) Compare(obj2 *KillChainPhasesProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Kill Chain Phases Property Length
	if len(o.KillChainPhases) != len(obj2.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in kill chain phases do not match: %d | %d", len(o.KillChainPhases), len(obj2.KillChainPhases))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in kill chain phases match: %d | %d", len(o.KillChainPhases), len(obj2.KillChainPhases))
		resultDetails = append(resultDetails, str)
		for index := range o.KillChainPhases {
			// Check Kill Chain Phases values
			if o.KillChainPhases[index].KillChainName != obj2.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- The kill chain name values do not match: %s | %s", o.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The kill chain name values match: %s | %s", o.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				resultDetails = append(resultDetails, str)
			}

			// Check Kill Chain Phases values
			if o.KillChainPhases[index].PhaseName != obj2.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- The kill chain phase values do not match: %s | %s", o.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The kill chain phase values match: %s | %s", o.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
