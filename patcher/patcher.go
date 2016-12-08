package patcher

import (
	"fmt"
	"github.com/yukpiz/kipi-patcher/request"
	"path/filepath"
)

const ROOT_PATH string = "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-patcher"

func Execute() error {
	fmt.Println("Execute kipi-patcher ===> (•ө•)♡")
	//Load yaml configuration.
	var config Config
	fpath := filepath.Join(ROOT_PATH, "kipi.yml")
	if err := LoadConfig(fpath, &config); err != nil {
		return err
	}

	info, err := request.GetBodyString(config.Url.PatchInfo)
	if err != nil {
		return err
	}
	return nil
}
