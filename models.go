package vxgo

type VxAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type VxNewsForm struct {
	Articles []*VxNews `json:"articles"`
}

type VxNews struct {
	Title              string `json:"title"`
	ThumbMediaId       string `json:"thumb_media_id"`
	Authod             string `json:"authod"`
	ShowCoverPic       string `json:"show_cover_pic"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	NeedOpenComment    uint32 `json:"need_open_comment"`
	OnlyFansCanComment uint32 `json:"only_fans_can_comment"`
}

type MediaId struct {
	MediaId string `json:"media_id"`
}

type VxMaterial struct {
	MediaId string `json:"media_id"`
	URL     string `json:"url"`
}

type NewCommitPoint struct {
	CommitID string
	Files    []string
}

type DailyPoetry struct {
	Status    string     `json:"status"`
	Data      PoetryData `json:"data"`
	IpAddress string     `json:"ipAddress"`
	Token     string     `json:"token"`
	Warning   string     `json:"warning"`
}

type PoetryData struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	Popularity int    `json:"popularity"`
	Origin     struct {
		Title     string   `json:"title"`
		Dynasty   string   `json:"dynasty"`
		Author    string   `json:"author"`
		Content   []string `json:"content"`
		Translate string   `json:"translate,omitempty"`
	} `json:"origin"`
	MatchTags         []string `json:"matchTags"`
	RecommendedReason string   `json:"recommendedReason"`
	CacheAt           string   `json:"cacheAt"`
}

type MessSendFilter struct {
	IsToAll bool `json:"is_to_all"`
	TagId   int  `json:"tag_id"`
}

type NewsSend struct {
	Filter            *MessSendFilter `json:"filter,omitempty"`
	MpNews            MediaId         `json:"mpnews"`
	MsgType           string          `json:"msgtype"`
	SendIgnoreReprint int             `json:"send_ignore_reprint"`
}
