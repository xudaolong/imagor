package main

import (
	"context"
	"github.com/xudaolong/imagor"
	"github.com/xudaolong/imagor/imagorpath"
	"github.com/xudaolong/imagor/loader/httploader"
	"github.com/xudaolong/imagor/vips"
	"io"
	"net/http"
	"os"
)

func main() {
	app := imagor.New(
		imagor.WithLoaders(httploader.New()),
		imagor.WithProcessors(vips.NewProcessor()),
	)
	ctx := context.Background()
	if err := app.Startup(ctx); err != nil {
		panic(err)
	}
	defer app.Shutdown(ctx)

	buf := downloadBytes("https://raw.githubusercontent.com/cshum/imagor/master/testdata/gopher.png")

	// serve via image buffer
	in := imagor.NewBlobFromBytes(buf)

	out, err := app.ServeBlob(ctx, in, imagorpath.Params{
		Width:  500,
		Height: 500,
		FitIn:  true,
		Filters: []imagorpath.Filter{
			{"fill", "yellow"},
			{"format", "jpg"},
		},
	})
	if err != nil {
		panic(err)
	}
	reader, _, err := out.NewReader()
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	file, err := os.Create("gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, reader); err != nil {
		panic(err)
	}
}

func downloadBytes(urlpath string) []byte {
	resp, err := http.Get(urlpath)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return buf
}
