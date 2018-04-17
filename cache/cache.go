package cache

import (
	"fmt"
	"g/x/web"
	"io/ioutil"
	"tone/global"
	"tone/o/tone"
)

func init() {
	var tones = make([]*tone.Tone, 0)
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
			var tone = &tone.Tone{
				Name:     f.Name(),
				Category: file.Name(),
				URL:      "http://192.168.11.2:3006/static/" + file.Name() + "/" + f.Name(),
			}
			tones = append(tones, tone)
		}
	}
	fmt.Printf("cached %d tones", len(tones))
	global.CacheTone = tones
}
