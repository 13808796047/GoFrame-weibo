package helpers

import "github.com/gogf/gf/text/gstr"

func Route2Class(routeName string) string {
	return gstr.Replace(routeName,".","-",-1)
}