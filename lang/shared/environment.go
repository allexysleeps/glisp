package shared

type Scope struct {
	Vars   map[string]*Variable
	Parent *Scope
}

func (s *Scope) Set(v Variable) {
	s.Vars[v.Name] = &v
}

func (s *Scope) Get(name string) *Variable {
	if v, ok := s.Vars[name]; ok {
		return v
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func CreateScope(parent *Scope) *Scope {
	return &Scope{
		Vars:   make(map[string]*Variable),
		Parent: parent,
	}
}
