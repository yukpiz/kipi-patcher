package patcher

import (
	"fmt"
	"github.com/yukpiz/kipi-patcher/args"
	"github.com/yukpiz/kipi-patcher/parser"
	"github.com/yukpiz/kipi-patcher/request"
)

func Execute() error {
	fmt.Println("Execute kipi-patcher ===> (•ө•)♡")
	//Parsing command line arguments.
	option := args.Option{}
	option.ParseArgs()
	fmt.Printf("%+v\n", option)
	//option.OutputHelp()

	//Request and get the patch header.
	text, err := request.GetBodyString(PATCH_INFO_URL)
	if err != nil {
		return err
	}

	info := new(patchinfo.PatchInfo)
	if err := info.Unmarshal(text); err != nil {
		return err
	}

	return nil
}
