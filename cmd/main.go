package main

import (
	"log"

	"github.com/shravaniphadolibm/my-csi-driver/driver"
)

func main() {
	endpoint := "unix:///tmp/csi.sock"
	d := driver.NewMyCSIDriver("my.csi.driver", "0.1.0", endpoint)
	if err := d.Run(); err != nil {
		log.Fatalf("driver failed to start: %v", err)
	}
}
