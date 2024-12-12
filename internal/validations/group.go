package validations

func ValidateCreateGroup(name string, description string) error {
	fields := map[string]struct {
		Value interface{}
		Tag   string
	}{
		"name":        {Value: name, Tag: "required,min=3,max=100"},
		"description": {Value: description, Tag: "required,min=3,max=100"},
	}

	return validateFields(fields)
}
