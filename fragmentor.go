package downloader

import (
	"bytes"
	"github.com/rs/xid"
	"io"
	"os"
	"os/exec"
)

type Fragmentor struct {
	filePath   string
	downloadId string
}

func (f *Fragmentor) GetDownloadId() string {
	return f.downloadId
}

func (f *Fragmentor) Fragment(file io.Reader, from string, to string) error {
	if err := f.setDownloadId().setFilePath().writeFile(file); err != nil {
		return err
	}

	if err := f.ffmpegFragment("in"+f.downloadId+".mp3", from, to); err != nil {
		return err
	}

	if err := f.ffmpegNormalize("temp" + f.downloadId + ".mp3"); err != nil {
		return err
	}

	return nil
}

func (f *Fragmentor) setDownloadId() *Fragmentor {
	f.downloadId = xid.New().String()

	return f
}

func (f *Fragmentor) setFilePath() *Fragmentor {
	f.filePath = f.downloadId + ".mp3"

	return f
}

func (f *Fragmentor) writeFile(file io.Reader) error {
	inputFile, err := os.Create("in" + f.downloadId + ".mp3")

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(file)

	inputFile.Write(buf.Bytes())

	return nil
}

func (f *Fragmentor) ffmpegFragment(source string, from string, to string) error {
	cmd := exec.Command(FFMPEG, "-i", "-loglevel", "quiet", source, "-ss", from, "-t", to, "-acodec", "copy", "temp"+f.downloadId+".mp3")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (f *Fragmentor) ffmpegNormalize(source string) error {
	cmd := exec.Command(FFMPEG, "-i", "-loglevel", "quiet", source, "-af", "volume=5dB", f.filePath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
