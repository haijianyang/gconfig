package gconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnvObjConfig struct {
	Bval bool   `json:"bval" env:"OBJ_BVAL"`
	Ival int    `json:"ival" env:"OBJ_IVAL"`
	I32  int32  `json:"i32" env:"OBJ_I32"`
	I64  int64  `json:"i64" env:"OBJ_I64"`
	Str  string `json:"str" env:"OBJ_STR"`
}

type EnvConfig struct {
	Bval bool         `json:"bval" env:"BVAL"`
	Ival int          `json:"ival" env:"IVAL"`
	I32  int32        `json:"i32" env:"I32"`
	I64  int64        `json:"i64" env:"I64"`
	Str  string       `json:"str" env:"STR"`
	Obj  EnvObjConfig `json:"obj"`
}

type FileObjConfig struct {
	Bval bool   `json:"bval"`
	Ival int    `json:"ival"`
	I32  int32  `json:"i32"`
	I64  int64  `json:"i64"`
	Str  string `json:"str"`
}

type FileCover struct {
	Str string `json:"str"`
}

type FileDefault struct {
	Str string `json:"str"`
}

type FileConfig struct {
	Bval  bool          `json:"bval"`
	Ival  int           `json:"ival"`
	I32   int32         `json:"i32"`
	I64   int64         `json:"i64"`
	Str   string        `json:"str"`
	Obj   FileObjConfig `json:"obj"`
	Cover FileCover     `json:"cover"`
}

type CoverConfig struct {
	Str string `json:"str" env:"COVER_STR"`
}

type Config struct {
	Env         EnvConfig   `json:"env"`
	File        FileConfig  `json:"file"`
	FileCover   FileCover   `json:"fileCover"`
	FileDefault FileDefault `json:"fileDefault"`
	Cover       CoverConfig `json:"cover"`
}

func TestUnmarshal(t *testing.T) {
	t.Run("Parse file", func(t *testing.T) {
		config := Config{}
		Unmarshal(&config)

		assert.Equal(t, config.File.Str, "str")
	})

	t.Run("Parse env", func(t *testing.T) {
		os.Setenv("STR", "str")

		config := Config{}
		Unmarshal(&config)

		os.Unsetenv("STR")

		assert.Equal(t, config.Env.Str, "str")
	})
}

func TestFile(t *testing.T) {
	t.Run("Parse file vars", func(t *testing.T) {
		config := Config{}
		gc := New()
		gc.ParseFile(&config)

		assert.Equal(t, config.File.Bval, true)
		assert.Equal(t, config.File.Ival, 1)
		assert.Equal(t, config.File.I32, int32(32))
		assert.Equal(t, config.File.I64, int64(64))
		assert.Equal(t, config.File.Str, "str")
		assert.Equal(t, config.File.Obj.Bval, true)
		assert.Equal(t, config.File.Obj.Ival, 1)
		assert.Equal(t, config.File.Obj.I32, int32(32))
		assert.Equal(t, config.File.Obj.I64, int64(64))
		assert.Equal(t, config.File.Obj.Str, "str")
	})

	t.Run("Parse default file", func(t *testing.T) {
		config := Config{}
		gc := New()
		gc.ParseFile(&config)

		assert.Equal(t, config.FileDefault.Str, "str")
	})

	t.Run("Cover file vars", func(t *testing.T) {
		config := Config{}
		gc := New()
		gc.ParseFile(&config)

		assert.Equal(t, config.FileCover.Str, "cover_str")
	})
}

func TestEnv(t *testing.T) {
	t.Run("Parse env vars to config", func(t *testing.T) {
		os.Setenv("BVAL", "true")
		os.Setenv("IVAL", "1")
		os.Setenv("I32", "32")
		os.Setenv("I64", "64")
		os.Setenv("STR", "str")
		os.Setenv("OBJ_BVAL", "true")
		os.Setenv("OBJ_IVAL", "1")
		os.Setenv("OBJ_I32", "32")
		os.Setenv("OBJ_I64", "64")
		os.Setenv("OBJ_STR", "str")

		config := Config{}
		gc := New()
		gc.ParseEnv(&config)

		os.Unsetenv("BVAL")
		os.Unsetenv("IVAL")
		os.Unsetenv("I32")
		os.Unsetenv("I64")
		os.Unsetenv("STR")
		os.Unsetenv("OBJ_BVAL")
		os.Unsetenv("OBJ_IVAL")
		os.Unsetenv("OBJ_I32")
		os.Unsetenv("OBJ_I64")
		os.Unsetenv("OBJ_STR")

		assert.Equal(t, config.Env.Bval, true)
		assert.Equal(t, config.Env.Ival, 1)
		assert.Equal(t, config.Env.I32, int32(32))
		assert.Equal(t, config.Env.I64, int64(64))
		assert.Equal(t, config.Env.Str, "str")
		assert.Equal(t, config.Env.Obj.Bval, true)
		assert.Equal(t, config.Env.Obj.Ival, 1)
		assert.Equal(t, config.Env.Obj.I32, int32(32))
		assert.Equal(t, config.Env.Obj.I64, int64(64))
		assert.Equal(t, config.Env.Obj.Str, "str")
	})

	t.Run("Cover vars", func(t *testing.T) {
		os.Setenv("COVER_STR", "cover_str")

		config := Config{}
		config.Cover.Str = "str"
		gc := New()
		gc.ParseEnv(&config)

		os.Unsetenv("COVER_STR")

		assert.Equal(t, config.Cover.Str, "cover_str")
	})
}
