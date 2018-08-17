package middleware

var excludelist = map[string]struct{}{}

func SaveBlacklist(excluded []string) {
	for _, excl := range excluded {
		excludelist[excl] = struct{}{}
	}
}

func IsExcluded(ns string) bool {
	_, excl := excludelist[ns]
	return excl
}
