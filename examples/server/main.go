package main

import (
	"github.com/xudaolong/imagor"
	"github.com/xudaolong/imagor/imagorpath"
	"github.com/xudaolong/imagor/loader/httploader"
	"github.com/xudaolong/imagor/server"
	"github.com/xudaolong/imagor/storage/filestorage"
	"github.com/xudaolong/imagor/vips"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())

	// create and run imagor server programmatically
	server.New(
		imagor.New(
			imagor.WithLogger(logger),
			imagor.WithUnsafe(true),
			imagor.WithProcessors(vips.NewProcessor()),
			imagor.WithLoaders(httploader.New()),
			imagor.WithStorages(filestorage.New("./")),
			imagor.WithResultStorages(filestorage.New("./")),
			imagor.WithResultStoragePathStyle(imagorpath.SuffixResultStorageHasher),
		),
		server.WithPort(8000),
		server.WithLogger(logger),
	).Run()
}
