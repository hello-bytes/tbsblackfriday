package blackfriday

import (
	"bytes"
	"log"
)

const IMAGEITEM int8 = 1
const HTMLITEM int8 = 2
const URLITEM int8 = 3
const URLAUTOITEM int8 = 4

//MakedownItem markdown的元素
type MakedownItem struct {
	ItemType int8
	Data     interface{}
}

//ImageItem markdown的图片
type ImageItem struct {
	Title string
	URL   string
}

type URLItemData struct {
	Title string
	URL   string
}

type HtmlItem struct {
	Html string
}

type Filter struct {
	items []MakedownItem
}

func FilterRenderer() Renderer {
	return &Filter{}
}

func (options *Filter) BlockCode(out *bytes.Buffer, text []byte, infoString string) {
	// log.Println("BlockCode ")
}

func (options *Filter) BlockQuote(out *bytes.Buffer, text []byte) {
	// log.Println("BlockQuote ")
}
func (options *Filter) BlockHtml(out *bytes.Buffer, text []byte) {
	log.Println("BlockHtml ")

	var item MakedownItem
	item.ItemType = HTMLITEM

	var htmlItem HtmlItem
	htmlItem.Html = string(text)
	item.Data = &htmlItem
	options.items = append(options.items, item)

	//log.Println(string(text))
}

func (options *Filter) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	// log.Println("Header ")
	text()
}
func (options *Filter) HRule(out *bytes.Buffer) {
	// log.Println("HRule ")
}
func (options *Filter) List(out *bytes.Buffer, text func() bool, flags int) {
	// log.Println("List ")
	text()
}
func (options *Filter) ListItem(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("ListItem ")
}
func (options *Filter) Paragraph(out *bytes.Buffer, text func() bool) {
	// log.Println("Paragraph ")
	text()
}
func (options *Filter) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	// log.Println("Table ")
}
func (options *Filter) TableRow(out *bytes.Buffer, text []byte) {
	// log.Println("TableRow ")
}
func (options *Filter) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("TableHeaderCell ")
}
func (options *Filter) TableCell(out *bytes.Buffer, text []byte, flags int) {
	// log.Println("TableCell ")
}
func (options *Filter) Footnotes(out *bytes.Buffer, text func() bool) {
	// log.Println("Footnotes ")
}
func (options *Filter) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	// log.Println("FootnoteItem ")
}
func (options *Filter) TitleBlock(out *bytes.Buffer, text []byte) {
	// log.Println("TitleBlock ")
}

func (options *Filter) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	//log.Println("AutoLink", kind)

	var item MakedownItem
	item.ItemType = URLAUTOITEM

	var urlItem URLItemData
	urlItem.URL = string(link)
	item.Data = &urlItem

	options.items = append(options.items, item)
}
func (options *Filter) CodeSpan(out *bytes.Buffer, text []byte) {
	// log.Println("CodeSpan ")
}
func (options *Filter) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	// log.Println("DoubleEmphasis ")
}
func (options *Filter) Emphasis(out *bytes.Buffer, text []byte) {
	// log.Println("Emphasis ")
}
func (options *Filter) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	// log.Println("Image : ", string(link))

	var item MakedownItem
	item.ItemType = IMAGEITEM

	var imageItem ImageItem
	imageItem.Title = string(title)
	//imageItem.Title = string(alt)
	imageItem.URL = string(link)
	item.Data = &imageItem
	options.items = append(options.items, item)
}
func (options *Filter) LineBreak(out *bytes.Buffer) {
	// log.Println("NormalText ")
}
func (options *Filter) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

	var item MakedownItem
	item.ItemType = URLITEM

	var urlItem URLItemData
	urlItem.Title = string(content)
	urlItem.URL = string(link)
	item.Data = &urlItem

	options.items = append(options.items, item)
}
func (options *Filter) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	// log.Println("RawHtmlTag ")
}
func (options *Filter) TripleEmphasis(out *bytes.Buffer, text []byte) {
	// log.Println("TripleEmphasis ")
}
func (options *Filter) StrikeThrough(out *bytes.Buffer, text []byte) {
	// log.Println("StrikeThrough ")
}
func (options *Filter) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	// log.Println("FootnoteRef ")
}

func (options *Filter) Entity(out *bytes.Buffer, entity []byte) {
	// log.Println("Entity ")
}
func (options *Filter) NormalText(out *bytes.Buffer, text []byte) {
	// log.Println("NormalText ", string(text))
	//if options.GetFlags()&HTML_USE_SMARTYPANTS != 0 {
	//	options.Smartypants(out, text)
	//} else {
	attrEscape(out, text)
	//}
}

/*
func (options *Filter) Smartypants(out *bytes.Buffer, text []byte) {
	smrt := smartypantsData{false, false}

	// first do normal entity escaping
	var escaped bytes.Buffer
	attrEscape(&escaped, text)
	text = escaped.Bytes()

	mark := 0
	for i := 0; i < len(text); i++ {
		if action := options.smartypants[text[i]]; action != nil {
			if i > mark {
				out.Write(text[mark:i])
			}

			previousChar := byte(0)
			if i > 0 {
				previousChar = text[i-1]
			}
			i += action(out, &smrt, previousChar, text[i:])
			mark = i + 1
		}
	}

	if mark < len(text) {
		out.Write(text[mark:])
	}
}*/

func (options *Filter) DocumentHeader(out *bytes.Buffer) {
	// log.Println("document header")
}
func (options *Filter) DocumentFooter(out *bytes.Buffer) {
	// log.Println("DocumentFooter")
}

func (options *Filter) GetFlags() int {
	return 0 |
		HTML_USE_XHTML |
		HTML_USE_SMARTYPANTS |
		HTML_SMARTYPANTS_FRACTIONS |
		HTML_SMARTYPANTS_DASHES |
		HTML_SMARTYPANTS_LATEX_DASHES | //-------
		EXTENSION_NO_INTRA_EMPHASIS |
		EXTENSION_TABLES |
		EXTENSION_FENCED_CODE |

		EXTENSION_STRIKETHROUGH |
		EXTENSION_SPACE_HEADERS |
		EXTENSION_HEADER_IDS |
		EXTENSION_BACKSLASH_LINE_BREAK |
		EXTENSION_DEFINITION_LISTS
}

//EXTENSION_AUTOLINK |
func AnalyzeMarkdown(text string) ([]MakedownItem, error) {
	var filterObj Filter
	//renderer := FilterRenderer()
	MarkdownOptions([]byte(text), &filterObj, Options{
		Extensions: commonExtensionsFilter})
	return filterObj.items, nil
}
