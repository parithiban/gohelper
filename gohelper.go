package gohelper

import (
	"reflect"
	"regexp"
)

// FindInSlice find the given data in slice
// Returns a boolean if data is present
func FindInSlice(sl interface{}, v interface{}) bool {
	switch reflect.TypeOf(sl).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(sl)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == v {
				return true
			}
		}

		return false
	default:
		return false
	}
}

// FindInStruct find the value of struct/[]struct is present based on the key
// Returns a boolean if data is present
func FindInStruct(st interface{}, v interface{}, key string) bool {

	s := reflect.ValueOf(st)

	switch s.Kind() {

	case reflect.Struct, reflect.Ptr:
		value := ValueIsAvailableInStruct(s, key)

		if value.IsValid() && value.Interface() == v {
			return true
		}

		return false

	case reflect.Slice:

		for i := 0; i < s.Len(); i++ {
			str := s.Index(i)
			value := ValueIsAvailableInStruct(str, key)

			if value.IsValid() && value.Interface() == v {
				return true
			}
		}

		return false
	default:
		return false
	}
}

// ValueIsAvailableInStruct find the value is present in struct based on key
func ValueIsAvailableInStruct(st reflect.Value, key string) reflect.Value {
	if st.Kind() == reflect.Ptr {
		if st.Elem().FieldByName(key).Kind() == reflect.Ptr {
			return st.Elem().FieldByName(key).Elem()
		}

		return st.Elem().FieldByName(key)
	}

	if st.FieldByName(key).Kind() == reflect.Ptr {
		return st.FieldByName(key).Elem()
	}

	return st.FieldByName(key)
}

// EmailValidation check the given input is a valid email address
// Returns a boolean if email is valid
func EmailValidation(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(s)
}
