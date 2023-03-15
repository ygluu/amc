package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

type newCsvObjHandler func() interface{}
type onReadObjHandler func(cur interface{}, prev interface{})
type parseHandler func(fv reflect.Value, svalue string)

var parsers map[string]parseHandler = make(map[string]parseHandler)

func RegParser(parseType interface{}, handler parseHandler) {
	t := reflect.TypeOf(parseType)
	tn := strings.ToLower(t.String())
	h := parsers[tn]
	if h != nil {
		Panic("[amc]RegParser => 类型<%s>已经注册解析器，不能重复注册", tn)
	}
	parsers[tn] = handler
}

func fileExist(file string) bool {
	_, err := os.Lstat(file)
	return !os.IsNotExist(err)
}

func isNotes(str string) bool {
	bytes := []byte(str)
	return (len(bytes) >= 1) && (bytes[0] == '/') && (len(bytes) >= 2) && (bytes[1] == '/')
}

func CsvRead(file string, newCsvObj newCsvObjHandler, onCsvObj onReadObjHandler) (errs []string) {
	errs = []string{}

	if !fileExist(file) {
		errs = append(errs, fmt.Sprintf("csv.Read：文件不存在，%s", file))
		return
	}

	f, e := os.OpenFile(file, os.O_RDONLY, 0644)
	if e != nil {
		errs = append(errs, fmt.Sprintf("csv.Read：打开文件失败，%v，%s", e, file))
		return
	}

	defer func() {
		f.Close()
	}()

	r := csv.NewReader(f)
	if r == nil {
		errs = append(errs, fmt.Sprintf("csv.Read：创建Reader失败，%s", file))
		return
	}

	records, e := r.ReadAll()
	if e != nil {
		errs = append(errs, fmt.Sprintf("csv.Read：读取文件失败，%s", file))
		return
	}
	if len(records) == 0 {
		errs = append(errs, fmt.Sprintf("csv.Read：空文件，%s", file))
		return
	}

	record := records[0]
	if len(record) == 0 {
		errs = append(errs, fmt.Sprintf("csv.Read：文件空行，%s", file))
		return
	}

	if !utf8.Valid([]byte(record[0])) {
		errs = append(errs, fmt.Sprintf("csv.Read：文件不是UTF8格式，%s", file))
		return
	}

	var columns map[string]int = nil

	for i, record := range records {
		if len(record) == 0 {
			errs = append(errs, fmt.Sprintf("csv.Read：第d行是空行，%s", i+1, file))
			continue
		}

		if isNotes(record[0]) {
			continue
		}

		if columns == nil {
			columns = make(map[string]int)
			for i, fn := range record {
				columns[strings.ToLower(fn)] = i + 1
			}
			continue
		}

		if len(record) != len(columns) {
			errs = append(errs, fmt.Sprintf("csv.Read：第d行字段数量和列数量不一致，%s", i+1, file))
			continue
		}

		obj := newCsvObj()
		fvs := refFiledInfo(obj)
		for fn, fv := range fvs {
			index := columns[fn]
			if index == 0 {
				errs = append(errs, fmt.Sprintf("csv.Read：第%d行字段名<%s>缺失，%s", i+1, fn, file))
				return
			}

			sv := record[index-1]

			istrue, err := myTypeParse(fv, fn, sv)
			if istrue {
				continue
			}
			if err != nil {
				errs = append(errs, fmt.Sprintf("csv.Read：第%d行字段<%s>，%s，%s", i+1, fn, err, file))
				return
			}

			err = setFiledValue(fv, sv)
			if err != nil {
				errs = append(errs, fmt.Sprintf("csv.Read：第%d行字段<%s>，%s，%s", i+1, fn, err, file))
				return
			}
		}
		fmt.Println(fmt.Sprintf("%+v", obj))
	}

	return
}

func myTypeParse(fv reflect.Value, fn string, sv string) (ret bool, err error) {
	defer func() {
		parseErr := recover()
		if parseErr != nil {
			err = errors.New(fmt.Sprintf("%v", parseErr))
		}
	}()
	ret = false
	err = nil
	parser := parsers[strings.ToLower(fv.Type().String())]
	if parser != nil {
		parser(fv, sv)
		ret = true
	}
	return
}

func refFiledInfo(obj interface{}) (ret map[string]reflect.Value) {
	ret = make(map[string]reflect.Value)
	vs := reflect.ValueOf(obj).Elem()
	ts := vs.Type()
	for i := 0; i < ts.NumField(); i++ {
		t := ts.Field(i)
		fn := t.Name
		if fn[0] >= 'A' && fn[0] <= 'Z' {
			ret[strings.ToLower(fn)] = vs.Field(i)
		}
	}
	return
}

func setFiledValue(fv reflect.Value, sv string) error {
	switch fv.Kind() {
	case reflect.String:
		fv.Set(reflect.ValueOf(sv))
		return nil
	case reflect.Int:
		v, err := strconv.Atoi(sv)
		if err == nil {
			fv.Set(reflect.ValueOf(v))
		}
		return nil
	case reflect.Uint:
		v, err := strconv.Atoi(sv)
		if err == nil {
			fv.Set(reflect.ValueOf(uint(v)))
		}
		return err
	case reflect.Int64:
		v, err := strconv.ParseInt(sv, 10, 64)
		if err == nil {
			fv.Set(reflect.ValueOf(v))
		}
		return err
	case reflect.Uint64:
		v, err := strconv.ParseInt(sv, 10, 64)
		if err == nil {
			fv.Set(reflect.ValueOf(uint64(v)))
		}
		return err
	case reflect.Float32:
		v, err := strconv.ParseFloat(sv, 32)
		if err == nil {
			fv.Set(reflect.ValueOf(v))
		}
		return err
	case reflect.Float64:
		v, err := strconv.ParseFloat(sv, 64)
		if err == nil {
			fv.Set(reflect.ValueOf(v))
		}
		return err
	case reflect.Bool:
		v, err := strconv.ParseBool(sv)
		if err == nil {
			fv.Set(reflect.ValueOf(v))
		}
		return err
	default:
		return errors.New("[amc]setFiledValue => 不支持的字段类型<" + fv.Type().String() + ">")
	}
}
