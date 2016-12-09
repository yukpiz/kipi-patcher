package args

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Option struct {
	To      int  `args:"to,t"`
	From    int  `args:"from,f"`
	Version bool `args:"version,v"`
	Patch   bool `args:"patch,p"`
	Info    bool `args:"info,i"`
	Server  bool `args:"server,s"`
	Help    bool `args:"help,h"`
}

const tag string = "args"

func (o *Option) ParseArgs() {
	if len(os.Args) < 2 {
		o.wrapError(errors.New("Argument is missing."))
		return
	}

	options := os.Args[1:]
	ovalue := reflect.ValueOf(o).Elem() //値セット用
	ftype := ovalue.Type()              //フィールド名取得用
	ttype := reflect.TypeOf(*o)         //タグ取得用
	for i := 0; i < ftype.NumField(); i++ {
		fname := ftype.Field(i).Name
		ofield, _ := ttype.FieldByName(fname)
		ops := strings.SplitN(ofield.Tag.Get(tag), ",", 2)

		for j, option := range options {
			replaced := strings.Replace(option, "-", "", -1)
			for _, op := range ops {
				if replaced == op {
					ofield := ovalue.FieldByName(fname)
					switch ofield.Type().Name() {
					case "int":
						ival, err := strconv.ParseInt(options[j+1], 10, 64)
						if err != nil {
							//Failed to parse int.
							o.wrapError(err)
							return
						}
						ofield.SetInt(ival)
					case "bool":
						ofield.SetBool(true)
					case "string":
						ofield.SetString(string(options[j+1]))
					}
				}
			}
		}
	}
	return
}

func (o *Option) wrapError(err error) {
	fmt.Println(err)
	o.initialize()
	o.Help = true
	return
}

func (o *Option) initialize() {
	*o = Option{}
}
