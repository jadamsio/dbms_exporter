// +build tds

package db

import _ "github.com/thda/tds"

func init() {
	name := "tds"
	drv := dsqlDrv(name)
	Register(name, &drv)
}
