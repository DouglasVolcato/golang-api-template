package utils

import (
	"app/src/domain/abstract/dtos"
	"errors"
	"fmt"
	"strconv"
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
	data           dtos.DtoType
	property       string
	label          string
	validatorTypes []string
}

func NewValidatorBuilder() *ValidatorBuilder {
	return &ValidatorBuilder{}
}

func (builder *ValidatorBuilder) Property(property string, label string) *ValidatorBuilder {
	builder.property = property
	builder.label = label
	return builder
}

func (builder *ValidatorBuilder) Validators(validatorTypes []string) *ValidatorBuilder {
	builder.validatorTypes = validatorTypes
	return builder
}

func (builder *ValidatorBuilder) Data(data dtos.DtoType) *ValidatorBuilder {
	builder.data = data
	return builder
}

func (builder *ValidatorBuilder) Validate() error {
	if builder.data == nil {
		return errors.New("Data is required")
	}

	value, exists := builder.data[builder.property]
	if !exists {
		for _, validatorType := range builder.validatorTypes {
			if validatorType == ValidatorTypes.IsRequired {
				return fmt.Errorf("%s is required", builder.label)
			}
		}
		return nil
	}

	for _, validatorType := range builder.validatorTypes {
		switch validatorType {
		case ValidatorTypes.IsEmail:
			if str, ok := value.(string); !ok || !strings.Contains(str, "@") {
				return fmt.Errorf("%s is invalid", builder.label)
			}

		case ValidatorTypes.IsString:
			if _, ok := value.(string); !ok {
				return fmt.Errorf("%s must be a string", builder.label)
			}

		case ValidatorTypes.IsInteger:
			if str, ok := value.(string); ok {
				if intValue, err := strconv.Atoi(str); err == nil {
					value = intValue
				} else {
					return fmt.Errorf("%s must be an integer number", builder.label)
				}
			}
			if floatValue, ok := value.(float64); ok {
				if floatValue == float64(int(floatValue)) {
					value = int(floatValue)
				} else {
					return fmt.Errorf("%s must be an integer number", builder.label)
				}
			}
			if _, ok := value.(int); !ok {
				return fmt.Errorf("%s must be an integer number", builder.label)
			}

		case ValidatorTypes.IsFloat:
			if str, ok := value.(string); ok {
				if floatValue, err := strconv.ParseFloat(str, 64); err == nil {
					value = floatValue
				} else {
					return fmt.Errorf("%s must be a decimal number", builder.label)
				}
			}
			if _, ok := value.(float64); !ok {
				return fmt.Errorf("%s must be a decimal number", builder.label)
			}
		}
	}

	return nil
}
