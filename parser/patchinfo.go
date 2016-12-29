package patchinfo

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

type PatchInfo struct {
	PatchAccept     int    `patchinfo:"patch_accept"`
	LatestVersion   int    `patchinfo:"main_version"`
	PatchUrl        string `patchinfo:"main_ftp"`
	LauncherVersion int    `patchinfo:"launcherinfo"`
	LoginIp         string `patchinfo:"login"`
	ExecArg         string `patchinfo:"arg"`
	LangPackName    string `patchinfo:"lang"`
	Status          bool
	StatusCode      int
}

const tag string = "patchinfo"

func (p *PatchInfo) Unmarshal(text string) error {
	kvs := parse(text)

	fnames := p.fields()
	fvalue := reflect.ValueOf(p).Elem()
	for _, fname := range fnames {
		t := reflect.TypeOf(*p)
		tfield, _ := t.FieldByName(fname)
		v, find := kvs[tfield.Tag.Get(tag)]
		if !find {
			continue
		}
		field := fvalue.FieldByName(fname)
		if !field.CanSet() {
			return errors.New("Can not set a structure.")
		}
		switch field.Type().Name() {
		case "string":
			//Field type `string`
			field.SetString(string(v))
		case "int":
			//Field type `int64`
			iv, _ := strconv.ParseInt(v, 10, 64)
			field.SetInt(iv)
		case "float":
			//Field type `float64`
			fv, _ := strconv.ParseFloat(v, 64)
			field.SetFloat(fv)
		case "uint":
			//Field type `uint64`
			uv, _ := strconv.ParseUint(v, 10, 64)
			field.SetUint(uv)
		}
	}

	return nil
}

func (p *PatchInfo) fields() []string {
	t := reflect.ValueOf(p).Elem().Type()
	fields := []string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fields = append(fields, f.Name)
	}
	return fields
}

func parse(value string) map[string]string {
	value = strings.Replace(value, "\r", "", -1)
	lines := strings.Split(value, "\n")
	kvs := map[string]string{}
	for _, line := range lines {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		kvs[kv[0]] = kv[1]
	}
	return kvs
}

func (p *PatchInfo) Print() error {
	tpl := template.Must(template.New("patchinfo").Parse(`
[Patch Infomation and Server Status]

 Patch Url          : {{.PatchUrl}}
 Patch Accept       : {{.PatchAccept}}
 Latest Version     : {{.LatestVersion}}
 Launcher Version   : {{.LauncherVersion}}
 Login Ip           : {{.LoginIp}}
 Exec Argument      : {{.ExecArg}}
 Language Pack Name : {{.LangPackName}}
 Status	            : {{.Status}}
 StatusCode         : {{.StatusCode}}
`))

	var o bytes.Buffer
	if err := tpl.Execute(&o, p); err != nil {
		return err
	}

	fmt.Println(o.String())
	return nil
}
