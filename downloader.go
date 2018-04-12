package downloader

import "io"

type Downloader struct {
	audioExtractor  *AudioExtractor
	audioUploader   *AudioUploader
	fileCleaner     *FileCleaner
	videoDownloader *VideoDownloader
	fragmentor      *Fragmentor
}

func New(storage string) *Downloader {
	return new(Downloader).Init(storage)
}

func (d *Downloader) Download(identifier string) error {
	if _, err := d.videoDownloader.Download(identifier); err != nil {
		return err
	}

	if _, err := d.audioExtractor.Extract(d.videoDownloader.GetTempFileName()); err != nil {
		return err
	}

	if err := d.audioUploader.
		Upload(d.videoDownloader.GetVideoMeta().GetTitle(), d.audioExtractor.GetFilePath()); err != nil {
		return err
	}

	d.fileCleaner.AddPath(d.audioExtractor.GetFilePath()).
		AddPath(d.videoDownloader.GetTempFileName())

	return nil
}

func (d *Downloader) Fragment(file io.Reader, from string, to string) error {
	if err := d.fragmentor.Fragment(file, from, to); err != nil {
		return err
	}

	d.fileCleaner.AddPath(d.fragmentor.GetInputPath()).
		AddPath(d.fragmentor.GetTempPath()).
		AddPath(d.fragmentor.GetOutputPath())

	return nil
}

func (d *Downloader) Clean() error {
	return d.fileCleaner.Clean()
}

func (d *Downloader) GetFileContents() (io.Reader, int64, error) {
	return d.audioUploader.GetFileContents(d.audioExtractor.GetFilePath())
}

func (d *Downloader) Init(storage string) *Downloader {
	d.videoDownloader = new(VideoDownloader).SetDownloadId()
	d.audioUploader = new(AudioUploader).Init(storage)
	d.fileCleaner = new(FileCleaner)
	d.fragmentor = new(Fragmentor)
	d.audioExtractor = new(AudioExtractor)

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
