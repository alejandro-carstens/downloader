package downloader

import "os"

type FileCleaner struct {
	audioFilePath string
	videoFilePath string
}

func (fc *FileCleaner) SetAudioFilePath(audioFilePath string) *FileCleaner {
	fc.audioFilePath = audioFilePath

	return fc
}

func (fc *FileCleaner) SetVideoFilePath(videoFilePath string) *FileCleaner {
	fc.videoFilePath = videoFilePath

	return fc
}

func (fc *FileCleaner) Clean() error {
	if err := os.Remove(fc.videoFilePath); err != nil {
		return err
	}

	return os.Remove(fc.audioFilePath)
}
