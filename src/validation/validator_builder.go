package validation

import (
	"app/src/domain/abstract"
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
	data           abstract.DtoType
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

func (builder *ValidatorBuilder) Data(data abstract.DtoType) *ValidatorBuilder {
	builder.data = data
	return builder
}

func (builder *ValidatorBuilder) Validate() error {
	if builder.data == nil {
		return fmt.Errorf("dados não fornecidos")
	}

	value, exists := builder.data[builder.property]
	if !exists {
		for _, validatorType := range builder.validatorTypes {
			if validatorType == ValidatorTypes.IsRequired {
				return fmt.Errorf("%s é obrigatório", builder.property)
			}
		}
		return nil
	}

	for _, validatorType := range builder.validatorTypes {
		switch validatorType {
		case ValidatorTypes.IsEmail:
			if str, ok := value.(string); !ok || !strings.Contains(str, "@") {
				return fmt.Errorf("%s está inválido", builder.property)
			}

		case ValidatorTypes.IsString:
			if _, ok := value.(string); !ok {
				return fmt.Errorf("%s deve ser uma string", builder.property)
			}

		case ValidatorTypes.IsInteger:
			if _, ok := value.(int); !ok {
				return fmt.Errorf("%s deve ser um número inteiro", builder.property)
			}

		case ValidatorTypes.IsFloat:
			if _, ok := value.(float64); !ok {
				return fmt.Errorf("%s deve ser um número decimal", builder.property)
			}
		}
	}

	return nil
}
