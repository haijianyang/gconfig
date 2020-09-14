package gconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
)

const DEFAULT_ENV = "development"
const DEFAULT_FILE = "default"
const DEFAULT_FILE_TYPE = ".json"

type Gconfig struct {
	Env         string
	Folder      string
	FileType    string
	DefaultFile string
}

var gconfig *Gconfig

func init() {
	gconfig = New()
}

func New() *Gconfig {
	folder, err := os.Getwd()
	if err != nil {
		log.Fatalln("[gconfig]", err)
	}

	gconfig := Gconfig{
		Env:         getEnv(),
		Folder:      folder,
		FileType:    DEFAULT_FILE_TYPE,
		DefaultFile: DEFAULT_FILE,
	}

	return &gconfig
}

func SetDefault(key string, value interface{}) {
	gconfig.SetDefault(key, value)
}

func (this *Gconfig) SetDefault(key string, value interface{}) {
	rvalue := reflect.ValueOf(this).Elem()
	rv := rvalue.FieldByName(strings.Title(key))
	if rv.IsValid() == true {
		rv.Set(reflect.ValueOf(value))
	}
}

func GetDefault(key string) string {
	return gconfig.GetDefault(key)
}

func (this *Gconfig) GetDefault(key string) string {
	rvalue := reflect.ValueOf(this).Elem()
	rv := rvalue.FieldByName(strings.Title(key))
	if rv.IsValid() == false {
		return ""
	} else {
		return rv.String()
	}
}

func Unmarshal(config interface{}) error {
	return gconfig.Unmarshal(config)
}

func (this *Gconfig) Unmarshal(config interface{}) error {
	if err := gconfig.ParseFile(config); err != nil {
		return err
	}

	if err := gconfig.ParseEnv(config); err != nil {
		return err
	}

	return nil
}

/* file */

func (this *Gconfig) ParseFile(config interface{}) error {
	if err := this.loadFile(config, this.DefaultFile); err != nil {
		return err
	}

	if err := this.loadFile(config, this.Env); err != nil {
		return err
	}

	return nil
}

func (this *Gconfig) loadFile(config interface{}, name string) error {
	filePath := path.Join(this.Folder, strings.Join([]string{name, this.FileType}, ""))
	_, err := os.Stat(filePath)

	if err != nil && err.Error() == os.ErrNotExist.Error() {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	json.Unmarshal(data, &config)
	if err != nil {
		return nil
	}

	return nil
}

/* env */

func ParseEnv(config interface{}) error {
	return gconfig.ParseEnv(config)
}

func (this *Gconfig) ParseEnv(config interface{}) error {
	rvalue := reflect.ValueOf(config)

	scanEnv(rvalue)

	return nil
}

func scanEnv(rvalue reflect.Value) {
	if rvalue.Kind() == reflect.Ptr {
		rvalue = rvalue.Elem()
	}

	for i := 0; i < rvalue.NumField(); i++ {
		sf := rvalue.Type().Field(i)
		rv := rvalue.Field(i)

		switch rv.Kind() {
		case reflect.Bool:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					bval, err := strconv.ParseBool(ev)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(bval))
					}
				}
			}
		case reflect.Int:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					ival, err := strconv.Atoi(ev)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(ival))
					}
				}
			}
		case reflect.Int32:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					ival, err := strconv.ParseInt(ev, 10, 32)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(int32(ival)))
					}
				}
			}
		case reflect.Int64:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					ival, err := strconv.ParseInt(ev, 10, 64)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(ival))
					}
				}
			}

		case reflect.Float32:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					fval, err := strconv.ParseFloat(ev, 32)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(fval))
					}
				}
			}
		case reflect.Float64:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					fval, err := strconv.ParseFloat(ev, 64)
					if err != nil {
						log.Println("[gconfig]", err)
					} else {
						rv.Set(reflect.ValueOf(fval))
					}
				}
			}

		case reflect.String:
			tag := sf.Tag.Get("env")
			if len(tag) > 0 {
				ev := os.Getenv(tag)
				if len(ev) > 0 {
					rv.Set(reflect.ValueOf(ev))
				}
			}
		case reflect.Struct:
			scanEnv(rv)
		}
	}
}

/* common */

func getEnv() string {
	env := os.Getenv("ENV")
	if os.Getenv("GO_ENV") != "" {
		env = os.Getenv("GO_ENV")
	}

	if len(env) == 0 {
		env = DEFAULT_ENV
	}

	return env
}
