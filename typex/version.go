//
// Warning:
//   This file is generated by go compiler, don't change it!!!
//   Build on: Ubuntu 22.04.1 LTS \n \l
//
package typex

import "fmt"

type Version struct {
	Version     string
	ReleaseTime string
}

func (v Version) String() string {
	return fmt.Sprintf("{\"releaseTime\":\"%s\",\"version\":\"%s\"}", v.ReleaseTime, v.Version)
}

var DefaultVersion = Version{
	Version:   `v0.5.0`,
	ReleaseTime: "2023-06-04 15:22:41",
}
var Banner = `
 **  Welcome to RULEX framework world <'_'>
**   Version: v0.5.0-9ec9c33aaa6401e
 **  Document: https://rulex.pages.dev
`
