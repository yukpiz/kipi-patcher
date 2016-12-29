package args

import (
	"fmt"
)

var VersionTmpl string = `
(•ө•)♡  kipi-patcher v1.0.0
Downloader of the M*binogi patch files.

Usage: patcher --from <FROM_PATCH_NUM> --to <TO_PATCH_NUM>
`

var HelpTmpl string = `
 --to <num>		
 --from <num>	
 --patch		Get patch infomations.
 --version		Show command versions.
 --info			Show md5 file hashes and size.
 --server		Show patch server status.
`

func (o *Option) PrintVersion() {
	fmt.Println(VersionTmpl)
}

func (o *Option) PrintHelp() {
	fmt.Println(VersionTmpl, HelpTmpl)
}
