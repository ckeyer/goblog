package modules

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ckeyer/blackfriday"
	"time"
)

const (
	htmlFlags = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_HARD_LINE_BREAK
)

type Blog struct {
	*Header
	mdfile  string
	Content []byte
}

type Blogs []*Blog

// NewBlog ...
func NewBlog(file string) (b *Blog, err error) {
	b = new(Blog)
	b.mdfile = file
	head, body, err := b.separateFile()
	if err != nil {
		return
	}
	//log.Debugf("head: %s \nbody: %s", head, body)

	b.decode2html(body)
	b.Header = new(Header)
	err = b.Header.Load(head)
	if err != nil {
		log.Error("文件加载错误: ", file)
		return
	}
	return
}

// separateFile 分割ms文件，返回头部json和 md内容两部分
func (b *Blog) separateFile() (header, body []byte, err error) {
	f, err := os.OpenFile(b.mdfile, os.O_RDONLY, 0444)
	if err != nil {
		log.Errorf("打开文件 %s 失败， %s", b.mdfile, err)
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorf("读取文件 %s 失败, %s", b.mdfile, err)
	}
	reg_s := regexp.MustCompile(`<!--`)
	reg_e := regexp.MustCompile(`-->`)
	start := reg_s.FindStringIndex(string(bs))
	end := reg_e.FindStringIndex(string(bs))
	if len(start) < 2 || len(end) < 2 {
		return
	}
	header = bs[start[1]:end[0]]
	body = bs[end[1]:]
	return
}

// (b *Blog)Decode2html ...
func (b *Blog) decode2html(bs []byte) (err error) {
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	outputOpts := blackfriday.Options{
		Extensions: extensions,
	}
	// Render the markdown file into HTML and return a new []Byte
	b.Content = blackfriday.MarkdownOptions(bs,
		renderer,
		outputOpts,
	)
	return nil
}

//
func (b Blogs) Len() (count int) {
	return len(b)
}

//
func (b Blogs) Less(i, j int) bool {
	ti, _ := time.Parse("2006-01-02", b[i].Date)
	tj, _ := time.Parse("2006-01-02", b[j].Date)
	// log.Debug("ti: ", ti)
	// log.Debug("tj: ", tj)
	if ti.Unix() > tj.Unix() {
		return true
	}
	return false
}

// (g *Group)Sort ...
func (b Blogs) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
