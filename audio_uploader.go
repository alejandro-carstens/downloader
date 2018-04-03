package downloader

import (
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	"io"
	"os"
)

const LOCAL string = "local"
const S3 string = "s3"

type AudioUploader struct {
	config stow.ConfigMap
	kind   string
}

func (au *AudioUploader) Init(kind string) *AudioUploader {
	switch kind {
	case S3:
		au.config = stow.ConfigMap{
			s3.ConfigAccessKeyID: os.Getenv("S3_ACCESS_KEY"),
			s3.ConfigSecretKey:   os.Getenv("S3_SECRET_KEY"),
			s3.ConfigRegion:      os.Getenv("S3_REGION"),
		}
		au.kind = S3
		break
	default:
		path := os.Getenv("LOCAL_PATH")

		if path == "" {
			path = LOCAL
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		au.config = stow.ConfigMap{"path": path}
		au.kind = LOCAL
	}

	return au
}

func (au *AudioUploader) Upload(fileName string, filePath string) error {
	location, err := stow.Dial(au.kind, au.config)

	if err != nil {
		return err
	}

	defer location.Close()

	container, err := au.getContainer(location)

	if err != nil {
		return err
	}

	contents, fileSize, err := au.getFileContents(filePath)

	if err != nil {
		return err
	}

	_, err = container.Put(fileName, contents, fileSize, nil)

	if err != nil {
		return err
	}

	return nil
}

func (au *AudioUploader) getContainer(location stow.Location) (stow.Container, error) {
	switch au.kind {
	case S3:
		return location.Container(os.Getenv("S3_BUCKET"))
	}

	return location.Container(LOCAL)
}

func (au *AudioUploader) getFileContents(fileName string) (io.Reader, int64, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return file, 0, err
	}

	stat, err := file.Stat()

	if err != nil {
		return file, 0, err
	}

	return file, stat.Size(), nil
}
