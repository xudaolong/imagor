package main

import (
	"github.com/xudaolong/imagor/config"
	"github.com/xudaolong/imagor/config/awsconfig"
	"github.com/xudaolong/imagor/config/gcloudconfig"
	"github.com/xudaolong/imagor/config/vipsconfig"
	"os"
)

func main() {
	var server = config.CreateServer(
		os.Args[1:],
		vipsconfig.WithVips,
		awsconfig.WithAWS,
		gcloudconfig.WithGCloud,
	)
	if server != nil {
		server.Run()
	}
}
