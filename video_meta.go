package downloader

import "strings"

type VideoMeta struct {
	title         string
	author        string
	datePublished string
	duration      string
}

func (vi *VideoMeta) SetTitle(title string) *VideoMeta {
	vi.title = strings.Replace(strings.Trim(title, " "), "/", "-", -1)

	return vi
}

func (vi *VideoMeta) GetTitle() string {
	return vi.title
}

func (vi *VideoMeta) SetAuthor(author string) *VideoMeta {
	vi.author = strings.Trim(author, " ")

	return vi
}

func (vi *VideoMeta) GetAuthor() string {
	return vi.author
}

func (vi *VideoMeta) SetDatePublished(datePublished string) *VideoMeta {
	vi.datePublished = strings.Trim(datePublished, " ")

	return vi
}

func (vi *VideoMeta) GetDatePublished() string {
	return vi.datePublished
}

func (vi *VideoMeta) SetDuration(duration string) *VideoMeta {
	vi.duration = strings.Trim(duration, " ")

	return vi
}

func (vi *VideoMeta) GetDuration() string {
	return vi.duration
}
