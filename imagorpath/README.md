# imagorpath

Parse and generate imagor endpoint using Go struct

```go
import "github.com/xudaolong/imagor/imagorpath"

...

func Test(t *testing.T) {
	params := imagorpath.Params{
		Image:    "raw.githubusercontent.com/cshum/imagor/master/testdata/gopher.png",
		FitIn:    true,
		Width:    500,
		Height:   400,
		PaddingTop: 20,
		PaddingBottom: 20,
		Filters: imagorpath.Filters{
			{
				Name: "fill",
				Args: "white",
			},
		},
	}

	// generate signed imagor endpoint from Params struct with secret
	path := imagorpath.Generate(params, imagorpath.NewDefaultSigner("mysecret"))

	assert.Equal(t, path, "OyGJyvfYJw8xNkYDmXU-4NPA2U0=/fit-in/500x400/0x20/filters:fill(white)/raw.githubusercontent.com/cshum/imagor/master/testdata/gopher.png")

	assert.Equal(t,
		// parse Params struct from signed imagor endpoint
		imagorpath.Parse(path),

		// Params include endpoint attributes with path and signed hash
		imagorpath.Params{
			Path:     "fit-in/500x400/0x20/filters:fill(white)/raw.githubusercontent.com/cshum/imagor/master/testdata/gopher.png",
			Hash:     "OyGJyvfYJw8xNkYDmXU-4NPA2U0=",
			Image:    "raw.githubusercontent.com/cshum/imagor/master/testdata/gopher.png",
			FitIn:    true,
			Width:    500,
			Height:   400,
			PaddingTop: 20,
			PaddingBottom: 20,
			Filters: imagorpath.Filters{
				{
					Name: "fill",
					Args: "white",
				},
			},
		},
	)
}

```