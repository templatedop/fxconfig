package fxconfig

import (
	//"gotemplate/logger"
)

type options struct {
	FileName string
	FileExt  string
	FilePath string
	//log      *logger.Logger
}

var DefaultConfigOptions = options{
	//log:nil,

	FileName: "config",
	FileExt:  "yaml",
	FilePath: "./configs",
}

type ConfigOption func(o *options)

func WithFileName(n string) ConfigOption {
	return func(o *options) {
		o.FileName = n
	}
}

func WithType(n string) ConfigOption {
	return func(o *options) {
		o.FileExt = n
	}
}

func WithPath(n string) ConfigOption {
	return func(o *options) {
		o.FilePath = n
	}
}
