package modules

import (
	"io/ioutil"
	"os"
	"regexp"
)

type Blog struct {
	*Header
	mdfile  string
	Content []byte
}

// NewBlog ...
func NewBlog(file string) (b *Blog, err error) {
	b = new(Blog)
	b.mdfile = file
	head, body, err := b.separateFile()
	if err != nil {
		return
	}
	//log.Debugf("head: %s \nbody: %s", head, body)

	b.Content = body
	b.Header = new(Header)
	err = b.Header.Load(head)
	if err != nil {
		return 
	}
	return
}

// separateFile 分割ms文件，返回头部json和 md内容两部分
func (b *Blog) separateFile() (header, body []byte, err error) {
	f, err := os.OpenFile(b.mdfile, os.O_RDONLY, 0444)
	if err != nil {
		log.Errorf("打开文件 %s 失败， %s",b.mdfile, err)
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorf("读取文件 %s 失败, %s",b.mdfile, err)
	}
	reg_s := regexp.MustCompile(`<!--`)
	reg_e := regexp.MustCompile(`-->`)
	start := reg_s.FindStringIndex(string(bs))
	end := reg_e.FindStringIndex(string(bs))
	if len(start) < 2 || len(end) < 2 {
		return
	}
	header =  bs[start[1]:end[0]]
	body = bs[end[1]:]
	return
}
