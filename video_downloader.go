package downloader

import (
	"github.com/rylio/ytdl"
	"os"
	"strings"
)

const DELIMITER string = "^"
const DATE_FORMAT string = "Jan 2 2006"

type VideoDownloader struct {
	tempFileName string
	videoMeta    *VideoMeta
	downloader   *ytdl.VideoInfo
	outputFile   *os.File
}

func (vd *VideoDownloader) Download(identifier string) (*VideoDownloader, error) {
	video, err := ytdl.GetVideoInfo(identifier)

	if err != nil {
		return vd, err
	}

	if err := vd.fillVideoMeta(video).setTempFileName().setOutputFile(); err != nil {
		return vd, err
	}

	video.Download(vd.getFormat(video), vd.outputFile)

	defer vd.outputFile.Close()

	return vd, nil
}

func (vd *VideoDownloader) GetVideoMeta() *VideoMeta {
    return vd.videoMeta
}

func (d *VideoDownloader) fillVideoMeta(video *ytdl.VideoInfo) *VideoDownloader {
	vd.videoMeta = new(VideoMeta).
		SetTitle(strings.Replace(video.Title, DELIMITER, "")).
		SetAuthor(video.Author).
		SetDatePublished(video.DatePublished.Format(DATE_FORMAT)).
		SetDuration(video.Duration)

	return vd
}

func (vd *VideoDownloader) setTempFileName() *VideoDownloader {
	vd.tempFileName = strings.Replace(vd.VideoMeta.GetTitle(), " ", DELIMITER, -1) + ".mp4"

	return vd
}

func (d *VideoDownloader) setOutputFile() error {
	vd.outputFile, err = os.Create(vd.tempFileName)

	if err != nil {
		return err
	}

	return nil
}

func (vd *VideoDownloader) getFormat(video *ytdl.VideoInfo) ytdl.Format {
	return video.Formats.Best(ytdl.FormatAudioEncodingKey)[0]
}
