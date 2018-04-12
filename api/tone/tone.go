package tone

import (
	"g/x/web"
	"io/ioutil"

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
			}
			tones = append(tones, tone)
		}
	}
	s.SendData(c, tones)
	// c.JSON(200, tones)
}

type Tone struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	URL         string `json:"url"`
}
