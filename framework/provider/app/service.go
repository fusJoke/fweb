package app

import (
	"errors"
	"flag"
	"github.com/fusjoke/fweb/framework"
	"github.com/fusjoke/fweb/framework/util"
	"path/filepath"
)

type HadeApp struct {
	container framework.Container
	baseFolder string
}

func (h HadeApp) Version() string {
	return "0.0.3"
}

func (h HadeApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数，默认为当前路径")
	flag.Parse()

	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

func (h HadeApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

func (h HadeApp) LogFolder() string {
	return filepath.Join(h.BaseFolder(), "log")
}

func (h HadeApp) HttpFolder() string {
	return filepath.Join(h.BaseFolder(), "http")
}

func (h HadeApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(), "cosole")
}

func (h HadeApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(), "storage")
}

func (h HadeApp) ProvideFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func (h HadeApp) CommandFolder() string {
	return filepath.Join(h.BaseFolder(), "command")
}

func (h HadeApp) RuntimeFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}


func (h HadeApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &HadeApp{
		baseFolder: baseFolder,
		container: container,
	}, nil
}

