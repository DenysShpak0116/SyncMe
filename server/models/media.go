package models

type XPhoto struct {
	XPhotoId int `db:"XPhotoId" json:"x_photo_id"`
	URL      string `db:"URL" json:"url"`
	PostId   int    `db:"PostId" json:"post_id"`
}

type XVideo struct {
    XVideoId int `db:"XVideoId" json:"x_video_id"`
    URL      string `db:"URL" json:"url"`
    PostId   int    `db:"PostId" json:"post_id"`
}