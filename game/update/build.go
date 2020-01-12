package update

import "ztaylor.me/cast"

// Build returns a new JSON object with "uri", "data"
func Build(uri string, json cast.JSON) cast.JSON {
	return cast.JSON{
		"uri":  uri,
		"data": json,
	}
}
