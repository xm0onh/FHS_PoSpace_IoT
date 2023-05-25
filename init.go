package FHS_PoSpace_IoT

import (
	"flag"
	"net/http"

	"github.com/xm0onh/FHS_PoSpace_IoT/config"
	"github.com/xm0onh/FHS_PoSpace_IoT/log"
)

func Init() {
	flag.Parse()
	log.Setup()
	config.Configuration.Load()
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
}
