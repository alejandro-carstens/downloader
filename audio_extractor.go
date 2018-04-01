package downloader

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const FFMPEG string = "ffmpeg"

type AudioExtractor struct {
    filePath string
}

func (ae *AudioExtractor) Extract(source string) (*AudioExtractor, error) {
    ffmpeg, err := exec.LookPath(FFMPEG)        
    
    if err != nil {
        return ae, err
    }
    
    return ae.setFilePath(source).ffmpegExtract(ffmpeg, source)
}

func (ae *AudioExtractor) GetFilePath() string {
    return ae.filePath
}

func (ae *AudioExtractor) setFilePath(source string) *AudioExtractor {
    ae.filePath = strings.Trim(source, filepath.Ext(source)) + ".mp3"
    
    return ae
}

func (ae *AudioExtractor) ffmpegExtract(ffmpeg string, source string) (*AudioExtractor, error) {
    cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", source, "-vn", ae.filePath)
	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
	    return ae, err
	}
	
	return ae, nil
}
