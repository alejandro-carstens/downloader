package downloader

type Downloader struct {
	audioExtractor  *AudioExtractor
	audioUploader   *AudioUploader
	fileCleaner     *FileCleaner
	videoDownloader *VideoDownloader
}

func (d *Downloader) Init(storage string) *Downloader {
	d.audioExtractor = new(AudioExtractor)
	d.audioUploader = new(AudioUploader).Init(storage)
	d.fileCleaner = new(FileCleaner)
	d.videoDownloader = new(VideoDownloader)

	return d
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

func Download(storage string, identifier string) error {
	return new(Downloader).Init(storage).Download(identifier)
}
