package log

var globalentries *Entries

func globalEntries() *Entries {
	if globalentries == nil {
		globalentries = New()
	}
	return globalentries
}
