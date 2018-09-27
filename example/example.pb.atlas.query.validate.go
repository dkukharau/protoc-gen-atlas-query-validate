// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/example.proto

package example

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.

var ExampleMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{
	"/example.TestService/List": map[string]options.FilteringOption{
		"first_name":           options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE}, ValueType: options.QueryValidate_STRING},
		"weight":               options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_LE}, ValueType: options.QueryValidate_NUMBER},
		"speciality":           options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_EQ, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE}, ValueType: options.QueryValidate_STRING},
		"comment":              options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"last_name":            options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"id":                   options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
		"custom_type_string":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.city":    options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE}, ValueType: options.QueryValidate_STRING},
		"home_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
}
var ExampleMethodsRequireSortingValidation = map[string][]string{
	"/example.TestService/List": []string{
		"first_name",
		"weight",
		"comment",
		"last_name",
		"id",
		"custom_type_string",
		"home_address.country",
	},
	"/example.TestService/Read": []string{
		"first_name",
		"weight",
		"comment",
		"last_name",
		"id",
		"custom_type_string",
		"home_address.country",
	},
}
var ExampleMethodsRequireFieldSelectionValidation = map[string][]string{
	"/example.TestService/List": {
		"first_name",
		"weight",
		"on_vacation",
		"speciality",
		"comment",
		"last_name",
		"id",
		"array",
		"custom_type",
		"custom_type_string",
		"home_address.city",
		"home_address.country",
		"home_address",
		"work_address.city",
		"work_address.country",
		"work_address",
	},
	"/example.TestService/Read": {
		"first_name",
		"weight",
		"on_vacation",
		"speciality",
		"comment",
		"last_name",
		"id",
		"array",
		"custom_type",
		"custom_type_string",
		"home_address.city",
		"home_address.country",
		"home_address",
		"work_address.city",
		"work_address.country",
		"work_address",
	},
}

func ExampleValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := ExampleMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func ExampleValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := ExampleMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func ExampleValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := ExampleMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}
