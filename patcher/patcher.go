package patcher

import (
	"github.com/yukpiz/kipi-patcher/args"
	"github.com/yukpiz/kipi-patcher/parser"
	"github.com/yukpiz/kipi-patcher/request"
)

func Execute() error {
	//Parsing command line arguments.
	option := args.Option{}
	option.ParseArgs()
	if option.Help {
		option.PrintHelp()
		return nil
	}

	if option.Version {
		option.PrintVersion()
		return nil
	}

	//Request and get the patch header.
	text, err := request.GetBodyString(PATCH_INFO_URL)
	if err != nil {
		return err
	}

	info := new(patchinfo.PatchInfo)
	if err := info.Unmarshal(text); err != nil {
		return err
	}

	if option.Patch {
		info.Print()
	}

	return nil
}
