package middleware

var blacklist = map[string]struct{}{}

func SaveBlacklist(excluded []string) {
	for _, excl := range excluded {
		blacklist[excl] = struct{}{}
	}
}

func IsExcluded(ns string) bool {
	_, excl := blacklist[ns]
	return excl
}
