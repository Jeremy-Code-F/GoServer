package main

import (
	gofarmlib "server/gofarm_mod"
	gofarmhttp "server/gofarm_mod/Http"
)

func main() {
	gofarmlib.Serve()
	gofarmhttp.Test()
}
