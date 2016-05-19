package main

import (
	"os"

	"github.com/elastic/beats/metricbeat/beater"

	// Uncomment the following line to include all official metricbeat module and metricsets
	//_ "github.com/elastic/beats/metricbeat/include"

	// Make sure all your modules and metricsets are linked here
	_ "github.com/ruflin/countbeat/module/count/increment"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "countbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
