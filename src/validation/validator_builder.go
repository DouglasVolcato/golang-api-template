package validation

import (
	"fmt"
	"strings"
)

var ValidatorTypes = struct {
	IsRequired string
	IsEmail    string
	IsString   string
	IsInteger  string
	IsFloat    string
}{
	IsRequired: "IS_REQUIRED",
	IsEmail:    "IS_EMAIL",
	IsString:   "IS_STRING",
	IsInteger:  "IS_INTEGER",
	IsFloat:    "IS_FLOAT",
}

type ValidatorBuilder struct {
	data           interface{}
	property       string
	validatorTypes []string
}

func NewValidatorBuilder() *ValidatorBuilder {
	return &ValidatorBuilder{}
}

func (builder *ValidatorBuilder) Property(property string) *ValidatorBuilder {
	builder.property = property
	return builder
}

func (builder *ValidatorBuilder) Validators(validatorTypes []string) *ValidatorBuilder {
	builder.validatorTypes = validatorTypes
	return builder
}

func (builder *ValidatorBuilder) Data(data interface{}) *ValidatorBuilder {
	builder.data = data
	return builder
}

func (builder *ValidatorBuilder) Validate() error {
	for _, validatorType := range builder.validatorTypes {
		if validatorType == ValidatorTypes.IsRequired && builder.data != nil {
			_, exists := builder.data.(map[string]interface{})[builder.property]
			if !exists {
				return fmt.Errorf("%s é obrigatório", builder.property)
			}
		}

		if validatorType == ValidatorTypes.IsEmail {
			if email, exists := builder.data.(map[string]interface{})[builder.property]; exists {
				if _, ok := email.(string); !ok || !strings.Contains(email.(string), "@") {
					return fmt.Errorf("%s está inválido", builder.property)
				}
			}
		}

		if validatorType == ValidatorTypes.IsString {
			if _, exists := builder.data.(map[string]interface{})[builder.property]; exists {
				if _, ok := builder.data.(map[string]interface{})[builder.property].(string); !ok {
					return fmt.Errorf("%s está inválido", builder.property)
				}
			}
		}

		if validatorType == ValidatorTypes.IsInteger {
			if _, exists := builder.data.(map[string]interface{})[builder.property]; exists {
				if _, ok := builder.data.(map[string]interface{})[builder.property].(int); !ok {
					return fmt.Errorf("%s está inválido", builder.property)
				}
			}
		}

		if validatorType == ValidatorTypes.IsFloat {
			if _, exists := builder.data.(map[string]interface{})[builder.property]; exists {
				if _, ok := builder.data.(map[string]interface{})[builder.property].(float64); !ok {
					return fmt.Errorf("%s está inválido", builder.property)
				}
			}
		}
	}

	return nil
}
