package schemas

import (
	"time"
)

type Part struct {
	ID   int64  `json:"id"`
	Salt string `json:"salt"`
}

type FileQuery struct {
	Name          string     `form:"name"`
	Search        string     `form:"search"`
	Type          string     `form:"type"`
	Path          string     `form:"path"`
	Op            string     `form:"op"`
	Starred       *bool      `form:"starred"`
	ParentID      string     `form:"parentId"`
	Category      string     `form:"category"`
	UpdatedAt     *time.Time `form:"updatedAt"`
	Sort          string     `form:"sort"`
	Order         string     `form:"order"`
	PerPage       int        `form:"perPage"`
	NextPageToken string     `form:"nextPageToken"`
}

type FileIn struct {
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Parts     []Part `json:"parts,omitempty"`
	MimeType  string `json:"mimeType"`
	ChannelID int64  `json:"channelId"`
	Path      string `json:"path" binding:"required"`
	Size      int64  `json:"size"`
	ParentID  string `json:"parentId"`
	Encrypted bool   `json:"encrypted"`
}

type FileOut struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	MimeType   string    `json:"mimeType"`
	Category   string    `json:"category,omitempty"`
	Encrypted  bool      `json:"encrypted"`
	Path       string    `json:"path,omitempty"`
	Size       int64     `json:"size,omitempty"`
	Starred    bool      `json:"starred"`
	ParentID   string    `json:"parentId,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type FileOutFull struct {
	*FileOut
	Parts     []Part `json:"parts,omitempty"`
	ChannelID int64  `json:"channelId"`
	Encrypted bool   `json:"encrypted"`
}

type FileUpdate struct {
	Name      string    `json:"name,omitempty"`
	Type      string    `json:"type,omitempty"`
	Path      string    `json:"path,omitempty"`
	Starred   *bool     `json:"starred,omitempty"`
	ParentID  string    `json:"parentId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type FileResponse struct {
	Files         []FileOut `json:"results"`
	NextPageToken string    `json:"nextPageToken,omitempty"`
}

type FileOperation struct {
	Files       []string `json:"files"  binding:"required"`
	Destination string   `json:"destination,omitempty"`
}

type DirMove struct {
	Source      string `json:"source" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

type MkDir struct {
	Path string `json:"path" binding:"required"`
}

type Copy struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

type FileCategoryStats struct {
	TotalFiles int    `json:"totalFiles"`
	TotalSize  int    `json:"totalSize"`
	Category   string `json:"category"`
}
