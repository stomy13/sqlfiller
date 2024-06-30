package sqlform

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
)

type restoreUserForm struct{}

func (f *restoreUserForm) Run() (string, error) {
	var userId string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's user id to restore?").
				Value(&userId).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("sorry, user id is required")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	sql := strings.ReplaceAll(restoreUserSqlTemplate, "${userId}", userId)
	return sql, nil
}

const restoreUserSqlTemplate = `BEGIN;

SELECT * FROM user WHERE id = ${userId};

UPDATE user SET deleted_at = NULL, updated_at = updated_at WHERE id = ${userId};

--COMMIT;
ROLLBACK;`
