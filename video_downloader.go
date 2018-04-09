package downloader

import (
	"github.com/rs/xid"
	"github.com/rylio/ytdl"
	"os"
)

const DATE_FORMAT string = "Jun 2 1992"

type VideoDownloader struct {
	downloadId   string
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

func (vd *VideoDownloader) GetTempFileName() string {
	return vd.tempFileName
}

func (vd *VideoDownloader) GetDownloadId() string {
	return vd.downloadId
}

func (vd *VideoDownloader) GetVideoMeta() *VideoMeta {
	return vd.videoMeta
}

func (vd *VideoDownloader) fillVideoMeta(video *ytdl.VideoInfo) *VideoDownloader {
	vd.videoMeta = new(VideoMeta).
		SetTitle(video.Title).
		SetAuthor(video.Author).
		SetDatePublished(video.DatePublished.Format(DATE_FORMAT)).
		SetDuration(video.Duration.String())

	return vd
}

func (vd *VideoDownloader) SetDownloadId() *VideoDownloader {
	vd.downloadId = xid.New().String()

	return vd
}

func (vd *VideoDownloader) setTempFileName() *VideoDownloader {
	vd.tempFileName = vd.downloadId + ".mp4"

	return vd
}

func (vd *VideoDownloader) setOutputFile() error {
	outputFile, err := os.Create(vd.tempFileName)

	if err != nil {
		return err
	}

	vd.outputFile = outputFile

	return nil
}

func (vd *VideoDownloader) getFormat(video *ytdl.VideoInfo) ytdl.Format {
	return video.Formats.Best(ytdl.FormatAudioEncodingKey)[0]
}
