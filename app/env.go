package app

import (
	"strconv"
)

var Host string = "localhost"
var Port int = 5000
var BaseUrl string = "http://" + Host + ":" + strconv.Itoa(Port)
