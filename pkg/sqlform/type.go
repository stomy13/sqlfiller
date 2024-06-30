package sqlform

import "github.com/charmbracelet/huh"

func inputSqlType() (string, error) {
	var sqlType string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose sql type").
				Options(
					huh.NewOption("User", "user"),
					huh.NewOption("Company", "company"),
				).
				Value(&sqlType),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return sqlType, nil
}
