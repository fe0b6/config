package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	confData interface{}
)

// Читаем конфиг
func Read(file string) (err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("[fatal]", err)
		return
	}

	err = json.Unmarshal(b, &confData)
	if err != nil {
		log.Fatalln("[fatal]", err)
		return
	}

	return
}

// Получаем значение строку
func GetStr(k ...string) (str string) {
	return GetStrSilent(false, k...)
}

// Получаем значение строку
func GetStrSilent(silent bool, k ...string) (str string) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	str, ok := a.(string)
	if !ok {
		log.Println("[error]", "value not string", a)
		return
	}

	return
}

// Получаем значение среза строк
func GetStrArr(k ...string) (str []string) {
	return GetStrArrSilent(false, k...)
}

// Получаем значение среза строк
func GetStrArrSilent(silent bool, k ...string) (str []string) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	t := a.([]interface{})

	str = make([]string, len(t))
	for i, v := range t {
		s, ok := v.(string)
		if !ok {
			log.Println("[error]", "value not []string", a)
			return []string{}
		}
		str[i] = s
	}

	return
}

// Получаем значение среза строк
func GetMap(k ...string) (str map[string]interface{}) {
	return GetMapSilent(false, k...)
}

// Получаем значение среза строк
func GetMapSilent(silent bool, k ...string) (str map[string]interface{}) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	str, ok := a.(map[string]interface{})
	if !ok {
		log.Println("[error]", "value not map[string]interface{}", a)
		return
	}

	return
}

// Получаем значение среза строк
func GetMapStr(k ...string) (str map[string]string) {
	return GetMapStrSilent(false, k...)
}

// Получаем значение среза строк
func GetMapStrSilent(silent bool, k ...string) (str map[string]string) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	arr, ok := a.(map[string]interface{})
	if !ok {
		log.Println("[error]", "value not map[string]interface{}", a)
		return
	}

	str = make(map[string]string)
	for k, v := range arr {
		str[k], ok = v.(string)
		if !ok {
			log.Println("[error]", "value not map[string]string", a)
			return
		}
	}

	return
}

// Получаем значение int
func GetInt(k ...string) (i int) {
	return GetIntSilent(false, k...)
}

// Получаем значение строку
func GetIntSilent(silent bool, k ...string) (i int) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	f, ok := a.(float64)
	if !ok {
		log.Println("[error]", "value not float64", a)
		return
	}

	i = int(f)

	return
}

// Получаем значение int
func GetInt64(k ...string) (i int64) {
	return GetInt64Silent(false, k...)
}

// Получаем значение строку
func GetInt64Silent(silent bool, k ...string) (i int64) {

	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	f, ok := a.(float64)
	if !ok {
		log.Println("[error]", "value not float64", a)
		return
	}

	i = int64(f)

	return
}

// Получаем значение bool
func GetBool(k ...string) (str bool) {
	return GetBoolSilent(false, k...)
}

// Получаем значение bool
func GetBoolSilent(silent bool, k ...string) (str bool) {
	a := getSilent(silent, confData, k...)
	if a == nil {
		return
	}

	str, ok := a.(bool)
	if !ok {
		log.Println("[error]", "value not bool", a)
		return
	}

	return
}

func getSilent(silent bool, i interface{}, k ...string) (intr interface{}) {
	intr = get(i, k...)
	if intr == nil {
		if !silent {
			log.Println("[error]", "param not found", k)
		}
	}
	return
}

func get(i interface{}, k ...string) (intr interface{}) {
	d, ok := i.(map[string]interface{})
	if !ok {
		log.Println("[error]", "interface not map", i)
		return
	}

	for key, val := range d {
		if key == k[0] {
			if len(k) == 1 {
				return val
			}

			l := k[1:]
			return get(val, l...)
		}
	}

	return
}
