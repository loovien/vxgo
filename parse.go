package vxgo

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var weChatTPL = `
<div style="text-align: center;">
<h1 style="margin:20px 0px; font-weight: 600;">撸完诗歌，撸技术， 做一个有逼格的程序员</h1>
<hr>
{{ if .Data }}
    <h4 style="margin: 10px 0; font-size:15px">今日诗歌推荐《{{.Data.Origin.Title}}》</h4>
	{{ range .Data.Origin.Content}}
		<p style="text-align:left; font-size:13px; text-indent:12px">{{.}}</p>
	{{end}}
	<p style="text-align: right;font-size:12px;">{{.Data.Origin.Dynasty}}-{{.Data.Origin.Author}}</p>
{{else}}
    <h4 style="margin: 10px 0; font-size:15px">今日诗歌歇菜了</h4>
	<p>................</p>
{{end}}
<strong> 点击左下角阅读更多, 进入正题</strong>
</div>
`

func ParseVxNews(file string) (*VxNews, error) {

	fd, err := os.Open(file)
	if err != nil {
		log.Printf("open file: %s failure: %v\n", file, err)
		return nil, err
	}
	vxNews := &VxNews{
		ThumbMediaId: "drwaZ2CgYKBpJE7GXmYSXPNSX_O5SLf4P5oyx_aiMLo",
		Content:      parsePoetry(),
		ShowCoverPic: "1",
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(txt, "<!--more-->") {
			break
		}
		if strings.HasPrefix(txt, "title") {
			title := strings.Split(txt, ":")
			if len(title) >= 2 {
				vxNews.Title = strings.TrimSpace(title[1])
			}
			continue
		}
		if strings.HasPrefix(txt, "date") {
			date := strings.Split(txt, "date:")
			if len(date) < 2 {
				vxNews.ContentSourceUrl = blogURL
				continue
			}
			blogTime, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(date[1]))
			if err != nil {
				continue
			}
			baseName := strings.TrimRight(filepath.Base(file), ".md")
			vxNews.ContentSourceUrl = fmt.Sprintf(
				"%s/%d/%s/%s/%s",
				blogURL,
				blogTime.Year(),
				blogTime.Format("01"),
				blogTime.Format("02"),
				baseName,
			)
			continue
		}
		if strings.HasPrefix(txt, "desc") {
			desc := strings.Split(txt, ":")
			if len(desc) >= 2 {
				vxNews.Digest = strings.TrimSpace(desc[1])
			}
			continue
		}
	}
	return vxNews, nil

}

func parsePoetry() string {
	poetry, _ := GetDailyPoetry()
	writer := new(bytes.Buffer)
	tpl, err := template.New("WeChatTPL").Parse(weChatTPL)
	if err != nil {
		log.Printf("template parse failure %v\n", err)
	}
	err = tpl.Execute(writer, poetry)
	if err != nil {
		log.Printf("template execute failure %v\n", err)
	}
	return writer.String()
}
