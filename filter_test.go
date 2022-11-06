package blackfriday

import (
	"log"
	"testing"
)

func TestFilterData(t *testing.T) {
	/*var input = "![](https://oss-cn-hangzhou.aliyuncs.com/codingsky/cdn/img/2020-01-01/2486514945123c1f10d8429851fb5f4d.png)"

	renderer := FilterRenderer()
	text := MarkdownOptions([]byte(input), renderer, Options{
		Extensions: commonExtensions})

	t.Log("output the markdown:")
	t.Log(text)

	//output := blackfriday.Run(input, blackfriday.WithNoExtensions())*/

	var input = "![](https://oss-cn-hangzhou.aliyuncs.com/codingsky/cdn/img/2020-01-01/2486514945123c1f10d8429851fb5f4d.png)"
	items, err := AnalyzeMarkdown(input)
	if err != nil {
		t.Error(err)
		return
	}

	for index, item := range items {
		if item.ItemType == IMAGEITEM {
			imageData := item.Data.(*ImageItem)
			log.Println(index, ",", imageData.Title, ",", imageData.URL)
		}

	}

	t.Error(nil)
}
