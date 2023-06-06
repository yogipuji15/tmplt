package tmplt

import "text/template"

func GetTemplate() (*template.Template,error) {
	return template.New("sql.tmpl").ParseFiles("tmplt/sql.tmpl")
}