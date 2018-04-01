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
    mp3File *os.File
}

func (se *AudioExtractor) ExtractAudio(source string) (*AudioExtractor, error) {
    ffmpeg, err := exec.LookPath(FFMPEG)        
    
    if err != nil {
        return se, err
    }
    
    return se.setFileName(source).ffmpegExtract(ffmpeg, source)
}

func (se *AudioExtractor) setFileName(source string) *AudioExtractor {
    se.filePath = strings.Trim(source, filepath.Ext(source)) + ".mp3"
    
    return se
}

func (se *AudioExtractor) ffmpegExtract(ffmpeg string, source string) (*AudioExtractor, error) {
    cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", source, "-vn", se.filePath)
	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
	    return se, err
	}
	
	return se, nil
}
