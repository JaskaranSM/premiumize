package types

import (
	"encoding/json"
)

type Transfer struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Message  string  `json:"message"`
	Status   string  `json:"status"`
	Progress float32 `json:"progress"`
	Src      string  `json:"src"`
	FolderId string  `json:"folder_id"`
	FileId   string  `json:"file_id"`
}

func (t *Transfer) IsComplete() bool {
	return t.Status == "finished"
}

type ListTransfersResponse struct {
	Status    string     `json:"status"`
	Transfers []Transfer `json:"transfers"`
}

func (l *ListTransfersResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, l)
}

type CreateTransferResponse struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

func (ct *CreateTransferResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, ct)
}

type Item struct {
	Id              string   `json:""id`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	Size            int64    `json:"size"`
	CreatedAt       int      `json:"created_at"`
	MimeType        string   `json:"mime_type"`
	TranscodeStatus string   `json:"transcode_status"`
	Link            string   `json:"link"`
	StreamLink      string   `json:"stream_link"`
	VirusScan       string   `json:"virus_scan"`
	Acodec          string   `json:"acodec"`
	Vcodec          string   `json:"vcodec"`
	FolderId        string   `json:"folder_id"`
	ResX            string   `json:"resx"`
	ResY            string   `json:"resy"`
	Duration        string   `json:"duration"`
	Bitrate         string   `json:"bitrate"`
	AudioTrackNames []string `json:"audio_track_names"`
}

type ItemResponse struct {
	Item
}

func (it *ItemResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, it)
}

type BreadCrumb struct {
	Id       string `json:"id"`
	ParentId string `json:"parent_id"`
	Name     string `json:"name"`
}

type ListFolderResponse struct {
	Status      string       `json:"status"`
	Content     []Item       `json:"content"`
	BreadCrumbs []BreadCrumb `json:"breadcrumbs"`
	Name        string       `json:"name"`
	ParentId    string       `json:"parent_id"`
	FolderId    string       `json:"folder_id"`
}

func (lf *ListFolderResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, lf)
}

type DeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (d *DeleteResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, d)
}

type Progress struct {
	Total      int64
	Completed  int64
	IsComplete bool
}

func (p *Progress) Write(chunk []byte) (int, error) {
	n := len(chunk)
	p.Completed += int64(n)
	return n, nil
}

func NewProgress(size int64) *Progress {
	return &Progress{Total: size}
}
