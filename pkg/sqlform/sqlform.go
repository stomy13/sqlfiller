package sqlform

func Run() (string, error) {
	sqlType, err := inputSqlType()
	if err != nil {
		return "", err
	}

	form, err := buildSqlForm(sqlType)
	if err != nil {
		return "", err
	}

	sql, err := form.Run()
	if err != nil {
		return "", err
	}

	if err := wait(); err != nil {
		return "", err
	}

	return sql, nil
}
