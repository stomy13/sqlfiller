package sqlform

import (
	"fmt"
)

type SqlForm interface {
	Run() (string, error)
}

func buildSqlForm(sqlType string) (SqlForm, error) {
	switch sqlType {
	case "company":
		return &restoreCompanyForm{}, nil
	case "user":
		return &restoreUserForm{}, nil
	}

	return nil, fmt.Errorf("sql type %s not found", sqlType)
}
