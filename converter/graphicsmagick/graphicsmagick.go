package graphicsmagick

import (
	"fmt"
	"github.com/pierrre/imageproxy"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type GraphicsMagickConverter struct {
	Executable string
	TempDir    string
}

func (converter *GraphicsMagickConverter) Convert(sourceImage *imageproxy.Image, parameters imageproxy.Parameters) (image *imageproxy.Image, err error) {
	tempDir, err := ioutil.TempDir(converter.TempDir, "imageproxy_")
	if err != nil {
		return
	}
	defer os.RemoveAll(tempDir)

	inFile := filepath.Join(tempDir, "image")
	outFile := inFile
	err = ioutil.WriteFile(inFile, sourceImage.Data, os.FileMode(0600))
	if err != nil {
		return
	}

	var arguments []string
	arguments = append(arguments, "mogrify")
	width, _ := parameters.GetInt("width")
	height, _ := parameters.GetInt("height")
	if width != 0 && height != 0 {
		if width <= 0 {
			err = fmt.Errorf("Invalid width")
		}
		if height <= 0 {
			err = fmt.Errorf("Invalid height")
		}
		arguments = append(arguments, "-resize", fmt.Sprintf("%dx%d", width, height))
	}
	format, _ := parameters.GetString("format")
	if len(format) > 0 {
		arguments = append(arguments, "-format", format)
		outFile = fmt.Sprintf("%s.%s", outFile, format)
	} else {
		format = sourceImage.Type
	}
	quality, _ := parameters.GetString("quality")
	if len(quality) > 0 {
		arguments = append(arguments, "-quality", quality)
	}
	arguments = append(arguments, inFile)

	cmd := exec.Command(converter.Executable, arguments...)
	err = cmd.Run()
	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(outFile)
	if err != nil {
		return
	}

	image = &imageproxy.Image{}
	image.Data = data
	image.Type = format

	return image, nil
}
