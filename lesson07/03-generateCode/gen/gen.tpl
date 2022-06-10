package {{.Package}}

type Student struct {
    {{- range .Fields}}
    {{.Name}} {{.Type}}
    {{- end}}
}

func (s *{{.TypeName}}) FromMap(m map[string]interface{}) {
    {{- range .Fields}}
    if v, ok := m["{{.JsonName}}"].({{.Type}}); ok {
        s.{{.Name}} = v
    }
    {{- end}}
}
