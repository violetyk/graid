package main

import . "regexp"

type RegexpUtil struct {
	*Regexp
}

func NewRegexpUtil(r *Regexp) *RegexpUtil {
	return &RegexpUtil{r}
}

func (r *RegexpUtil) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		if i == 0 || len(name) == 0 {
			continue
		}
		captures[name] = match[i]
	}
	return captures
}
