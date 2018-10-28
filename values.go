package mikrotik

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Values struct {
	values map[string]string
}

func ValuesFrom(m map[string]string) Values {
	return Values{m}
}

func (v Values) Set(key, val string) {
	if v.values == nil {
		v.values = make(map[string]string)
	}

	v.values[key] = val
}

func (v Values) Lookup(key string) (string, bool) {
	if v.values == nil {
		return "", false
	}

	val, ok := v.values[key]
	return val, ok
}

func (v Values) Get(key string) string {
	if v.values == nil {
		return ""
	}

	if val, ok := v.values[key]; ok {
		return val
	}

	return ""
}

func (v Values) Del(key string) {
	if v.values == nil {
		return
	}

	delete(v.values, key)
}

func (v Values) String() string {
	if v.values == nil {
		return ""
	}

	var vals []string
	for key, val := range v.values {
		vals = append(vals, fmt.Sprintf("%s=%s", key, val))
	}

	return strings.Join(vals, " ")
}

func (v Values) To(i interface{}) (err error) {
	rv := reflect.ValueOf(i).Elem()

	//if it`s i not a slice
	if rv.Kind() != reflect.Slice {

		itemType := reflect.TypeOf(i)
		for itemType.Kind() == reflect.Ptr {
			itemType = itemType.Elem()
		}
		item := reflect.New(itemType)

		if err := v.setStruct(item.Elem(), itemType); err != nil {
			return err
		}

		if rv.Kind() != reflect.Ptr && item.Kind() == reflect.Ptr {
			rv.Set(item.Elem())
		} else {
			rv.Set(item)
		}
		return
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	itemType := rv.Type().Elem()

	//if slice element is pointer on struct
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
		item := reflect.New(itemType)
		if err := v.setStruct(item.Elem(), itemType); err != nil {
			return err
		}

		rv.Set(reflect.Append(rv, item))
	} else {
		//if slice element is struct
		item := reflect.New(itemType)
		if err := v.setStruct(item.Elem(), itemType); err != nil {
			return err
		}

		rv.Set(reflect.Append(rv, item.Elem()))
	}

	return
}

func (v Values) setStruct(rv reflect.Value, rt reflect.Type) (err error) {
	for i := 0; i < rv.NumField(); i++ {
		vfield := rv.Field(i)
		tfield := rt.Field(i)
		if !vfield.IsValid() && !vfield.CanSet() {
			continue
		}

		var fieldname string

		if tag, ok := tfield.Tag.Lookup("mikrotik"); ok {
			fieldname = tag
		} else {
			fieldname = ToMikrotikName(tfield.Name)
		}

		// log.Println(fieldname)
		val, ok := v.Lookup(fieldname)
		if !ok {
			continue
		}

		if tag, ok := tfield.Tag.Lookup("trim"); ok {
			val = strings.Trim(val, tag)
		}

		switch vfield.Interface().(type) {
		case bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			vfield.SetBool(b)

		case int:
			n, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			vfield.SetInt(int64(n))

		case net.IP:
			ip := net.ParseIP(val)
			vfield.Set(reflect.ValueOf(ip))

		case time.Duration:
			dur, err := time.ParseDuration(val)
			if err != nil {
				return err
			}
			vfield.Set(reflect.ValueOf(dur))

		default:
			fieldVal := reflect.ValueOf(val).Convert(vfield.Type())
			vfield.Set(fieldVal)
		}
	}

	return nil
}

//ToFieldName convert incoming fieldname to struct fieldname, example:
// address -> Address
// actual-interface -> ActualInterface
func ToFieldName(fieldname string) string {
	var newname strings.Builder
	newname.Grow(len(fieldname))

	for _, s := range strings.Split(fieldname, "-") {
		newname.WriteString(strings.Title(s))
	}

	return newname.String()
}

//ToMikrotikName convert struct fieldname to mikrotik like:
// Address -> address
// ActualInterface -> actual-interface
func ToMikrotikName(fieldname string) string {
	var newname strings.Builder
	newname.Grow(len(fieldname) + 2)

	for i, c := range fieldname {
		if i > 0 && unicode.IsUpper(c) {
			newname.WriteString("-")
		}
		newname.WriteRune(unicode.ToLower(c))

	}

	return newname.String()
}

func ToArgs(i interface{}) (args []string) {
	rv := reflect.ValueOf(i)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rt := rv.Type()

	if rv.Kind() == reflect.String {
		return strings.Split(rv.String(), " ")
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		structField := rt.Field(i)

		// if !field.CanSet() {
		// 	continue
		// }

		if IsEmpty(field) {
			continue
		}

		var name string
		tag, ok := structField.Tag.Lookup("mikrotik")
		if ok && tag != "" {
			if tag == "-" {
				continue
			}

			tags := strings.Split(tag, ",")
			if len(tags) > 1 && tags[1] == "ro" {
				continue
			}
			name = tags[0]
		} else {
			name = ToMikrotikName(structField.Name)
		}

		args = append(args, fmt.Sprintf("=%s=%v", name, field.Interface()))
	}

	return
}

func IsEmpty(v reflect.Value) bool {
	switch v.Interface().(type) {
	case string:
		return v.String() == ""
	case bool:
		return v.Bool() == false
	case int:
		return v.Int() == 0
	case net.IP:
		return v.Len() == 0
	}

	return false
}

func SetID(i interface{}, id string) {
	rv := reflect.ValueOf(i)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if !rv.CanSet() {
		return
	}

	rv.FieldByName("ID").SetString(id)
}
