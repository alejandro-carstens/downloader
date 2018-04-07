package downloader

import "io"

type Downloader struct {
	audioExtractor  *AudioExtractor
	audioUploader   *AudioUploader
	fileCleaner     *FileCleaner
	videoDownloader *VideoDownloader
}

func New(storage string) *Downloader {
	return new(Downloader).Init(storage)
}

func (d *Downloader) Download(identifier string) error {
	_, err := d.videoDownloader.Download(identifier)

	if err != nil {
		return err
	}

	_, err = d.audioExtractor.Extract(d.videoDownloader.GetTempFileName())

	if err != nil {
		return err
	}

	err = d.audioUploader.
		Upload(d.videoDownloader.GetVideoMeta().GetTitle(), d.audioExtractor.GetFilePath())

	if err != nil {
		return err
	}

	return d.fileCleaner.SetAudioFilePath(d.audioExtractor.GetFilePath()).
		SetVideoFilePath(d.videoDownloader.GetTempFileName()).
		Clean()
}

func (d *Downloader) Init(storage string) *Downloader {
	d.audioExtractor = new(AudioExtractor)
	d.audioUploader = new(AudioUploader).Init(storage)
	d.fileCleaner = new(FileCleaner)
	d.videoDownloader = new(VideoDownloader)

	return d
}

func (d *Downloader) Get(key string) (io.ReadCloser, error) {
	return d.audioUploader.Get(key)
}

func (d *Downloader) GetVideoMeta() *VideoMeta {
	return d.videoDownloader.GetVideoMeta()
}

func (d *Downloader) GetPath() string {
	return d.audioUploader.GetPath()
}

func (d *Downloader) GetKey() string {
	return d.audioUploader.GetKey()
}

func (d *Downloader) GetDownloadId() string {
	return d.videoDownloader.GetDownloadId()
}
