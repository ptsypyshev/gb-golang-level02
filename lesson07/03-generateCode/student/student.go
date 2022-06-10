package student

type Student struct {
	FirstName string
	LastName  string
	Age       int
	IsMale    bool
	Grade     float64
}

func (s *Student) FromMap(m map[string]interface{}) {
	if v, ok := m["first_name"].(string); ok {
		s.FirstName = v
	}
	if v, ok := m["last_name"].(string); ok {
		s.LastName = v
	}
	if v, ok := m["age"].(int); ok {
		s.Age = v
	}
	if v, ok := m["is_male"].(bool); ok {
		s.IsMale = v
	}
	if v, ok := m["grade"].(float64); ok {
		s.Grade = v
	}
}
