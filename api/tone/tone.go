package tone

import (
	"g/x/web"
	"io/ioutil"
	"tone/global"

	"github.com/gin-gonic/gin"
)

type ToneServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewToneServer(parent *gin.RouterGroup, name string) *ToneServer {
	var s = ToneServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("/list", s.handleListTones)
	return &s
}

func (s *ToneServer) handleListTones(c *gin.Context) {
	// var tones = getTones()
	// var page, _ = strconv.ParseInt(c.Query("page"), 10, 32)
	// var skip, _ = strconv.ParseInt(c.Query("skip"), 10, 32)
	// var start = (page - 1) * skip
	// var end = (page) * skip
	// if int(start) > len(tones) {
	// 	s.SendData(c, []*Tone{})
	// } else {
	// 	if int(end) >= len(tones) {
	// 		s.SendData(c, tones[start:len(tones)])
	// 	} else {
	// 		s.SendData(c, tones[start:end])
	// 	}
	// }
	s.SendData(c, global.CacheTone)
	// c.JSON(200, tones)
}
func getTones() []*Tone {
	var tones = make([]*Tone, 0)
	files, err := ioutil.ReadDir("./upload")
	if err != nil {
		web.AssertNil(err)
	}
	for _, file := range files {
		var first, err = ioutil.ReadDir("./upload/" + file.Name())
		if err != nil {
			web.AssertNil(err)
		}
		for _, f := range first {
			var tone = &Tone{
				Name:     f.Name(),
				Category: file.Name(),
				URL:      "http://192.168.11.2:3006/static/" + file.Name() + "/" + f.Name(),
			}
			tones = append(tones, tone)
		}
	}
	return tones
}

type Tone struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	URL         string `json:"url"`
}
