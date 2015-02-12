package regexputil

import "regexp"

type RegexpUtil struct {
	r *regexp.Regexp
}

func Compile(expr string) (ru *RegexpUtil, err error) {
	r, err := regexp.Compile(expr)
	ru = &RegexpUtil{r}
	return
}

func (ru *RegexpUtil) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := ru.r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range ru.r.SubexpNames() {
		if i == 0 || len(name) == 0 {
			continue
		}
		captures[name] = match[i]
	}
	return captures
}
