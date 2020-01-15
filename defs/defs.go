// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package defs

// To support additional specification versions, I need to update the following code
// libstix2/t_collectiondataManifest.go sqlCollectionDataWhereSpecVersion()
// server/server.go processURLParameters()

// These are the STIX timestamp formats.
// TIME_RFC_3339 is the general purpose timestamp.
// TIME_RFC_3339_MICRO is the timestamp for the created and modified properties.
const (
	STIX_VERSION       = "2.1"
	TAXII_VERSION      = "2.1"
	MEDIA_TYPE_STIX    = "application/stix+json"
	MEDIA_TYPE_TAXII   = "application/taxii+json"
	MEDIA_TYPE_STIX20  = "application/vnd.oasis.stix+json;version=2.0"
	MEDIA_TYPE_TAXII20 = "application/vnd.oasis.taxii+json;version=2.0"
	MEDIA_TYPE_STIX21  = "application/stix+json;version=2.1"
	MEDIA_TYPE_TAXII21 = "application/taxii+json;version=2.1"
	MEDIA_TYPE_JSON    = "application/json"
	MEDIA_TYPE_HTML    = "text/html; charset=utf-8"

	TIME_RFC_3339       = "2006-01-02T15:04:05Z07:00"
	TIME_RFC_3339_MILLI = "2006-01-02T15:04:05.999Z07:00"
	TIME_RFC_3339_MICRO = "2006-01-02T15:04:05.999999Z07:00"

	STRICT_TYPES  = true
	KEEP_RAW_DATA = false
)
