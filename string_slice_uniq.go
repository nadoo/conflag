package conflag

type stringSliceUniqValue struct {
	*stringSliceValue
}

func newStringSliceUniqValue(val []string, p *[]string) *stringSliceUniqValue {
	return &stringSliceUniqValue{stringSliceValue: newStringSliceValue(val, p)}
}

func (s *stringSliceUniqValue) Set(val string) error {
	if !s.changed {
		*s.value = []string{val}
		s.changed = true
	}

	dup := false
	for _, v := range *s.value {
		if v == val {
			dup = true
		}
	}

	if !dup {
		*s.value = append(*s.value, val)
	}

	return nil
}

func (s *stringSliceUniqValue) Type() string {
	return "stringSliceUniq"
}

func (s *stringSliceUniqValue) String() string {
	return ""
}

// StringSliceUniqVar works like StringSliceVar but the items in slice are unique.
func (c *Conflag) StringSliceUniqVar(p *[]string, name string, value []string, usage string) {
	c.Var(newStringSliceUniqValue(value, p), name, usage)
}

// StringSliceUniq works like StringSlice but the items in slice are unique.
func (c *Conflag) StringSliceUniq(name string, value []string, usage string) *[]string {
	p := []string{}
	c.StringSliceUniqVar(&p, name, value, usage)
	return &p
}
