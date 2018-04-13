package downloader

import (
	"bytes"
	"github.com/rs/xid"
	"io"
	"os"
	"os/exec"
)

type Fragmentor struct {
	downloadId string
	outputPath string
	tempPath   string
	inputPath  string
}

func (f *Fragmentor) GetDownloadId() string {
	return f.downloadId
}

func (f *Fragmentor) GetOutputPath() string {
	return f.outputPath
}

func (f *Fragmentor) GetInputPath() string {
	return f.inputPath
}

func (f *Fragmentor) GetTempPath() string {
	return f.tempPath
}

func (f *Fragmentor) Fragment(file io.Reader, from string, to string) error {
	if err := f.setDownloadId().setPaths().writeFile(file); err != nil {
		return err
	}

	if err := f.ffmpegFragment(from, to); err != nil {
		return err
	}

	if err := f.ffmpegNormalize(); err != nil {
		return err
	}

	return nil
}

func (f *Fragmentor) setDownloadId() *Fragmentor {
	f.downloadId = xid.New().String()

	return f
}

func (f *Fragmentor) setPaths() *Fragmentor {
	f.outputPath = f.downloadId + ".mp3"
	f.tempPath = "temp" + f.downloadId + ".mp3"
	f.inputPath = "in" + f.downloadId + ".mp3"

	return f
}

func (f *Fragmentor) writeFile(file io.Reader) error {
	inputFile, err := os.Create(f.inputPath)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(file)

	inputFile.Write(buf.Bytes())

	return nil
}

func (f *Fragmentor) ffmpegFragment(from string, to string) error {
	cmd := exec.Command(FFMPEG, "-loglevel", "quiet", "-i", f.inputPath, "-ss", from, "-to", to, "-acodec", "copy", f.tempPath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (f *Fragmentor) ffmpegNormalize() error {
	cmd := exec.Command(FFMPEG, "-loglevel", "quiet", "-i", f.tempPath, "-af", "volume=5dB", f.outputPath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
