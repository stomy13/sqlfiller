package sqlform

type restoreCompanyForm struct{}

func (f *restoreCompanyForm) Run() (string, error) {
	return "company sql", nil
}
