package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Alertor interface{}

type WebHook struct {
	Header *HookHeader

	Repos  string
	Branch string
	User   string
	Script string
}

// webhook github发送http请求头
type HookHeader struct {
	Agent     string `json:"User-Agent"`
	Delivery  string `json:"X-GitHub-Delivery"`
	Event     string `json:"X-GitHub-Event"`
	Signature string `json:"X-Hub-Signature"`
}

//
type Payload struct {
	Ref     string      `json:"ref"` //"ref": "refs/heads/master",
	Commits []*Commit   `json:"commits"`
	Repo    *Repository `json:"repository"`
	Pusher  *GitUser    `json:"pusher"`
}

type Commit struct {
	Branch    string   `json:"branch"`
	Author    *GitUser `json:"author"`
	Committer *GitUser `json:"committer"`
}

type Repository struct {
	Url      string `json:"url"`
	FullName string `json:"full_name"`
}

type GitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func DoWebhook(res http.ResponseWriter, req *http.Request) {
	hook := &WebHook{}
	hook.HookHandle(res, req)
}

func (w *WebHook) HookHandle(res http.ResponseWriter, req *http.Request) {
	signature := req.Header.Get("X-Hub-Signature")
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("WebHook Error. %s", err)
		return
	}
	if signature != "sha1="+HmacSha1(bs, []byte("HOOK_SECRET")) {
		log.Notice("Check Error...")
		return
	}
	err = w.LoadPayload(bs)
	if err != nil {
		log.Errorf("WebHook Error. %s", err)
		return
	}
	if ok := w.Check(); ok {
		go w.Done()
	}
	res.Write([]byte("receive over"))
}

func (w *WebHook) LoadPayload(data []byte) (err error) {
	var p Payload
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Errorf("WebHook Error. %s", err)
		return
	}
	tmp := strings.Split(p.Ref, "/")
	if l := len(tmp); l == 0 {
		return
	} else {
		w.Branch = tmp[l-1]
	}
	w.Repos = p.Repo.FullName
	w.User = p.Pusher.Name
	return
}

func (w *WebHook) Done() {
	script := w.Script
	if _, err := os.Stat(script); err != nil {
		log.Errorf("WebHook Error. %s", err)
		return
	}
	out, err := exec.Command("bash", "-c", script).Output()
	if err != nil {
		log.Errorf("WebHook Error. %s", err)
		return
	}
	log.Notice(out)
}

// 检查是否应该触发
func (w *WebHook) Check() (ok bool) {
	// for _, h := range w.Config.Hooks {
	// 	if h.Repos != w.Repos {
	// 		continue
	// 	}
	// 	for _, m := range h.Monitors {
	// 		if m.Branch != w.Branch && m.Branch != "" {
	// 			continue
	// 		}
	// 		if m.User != w.User && m.User != "" {
	// 			continue
	// 		}
	// 		if m.Script != "" {
	// 			w.Script = m.Script
	// 			return true
	// 		}
	// 	}
	// }
	return false
}

func HmacSha1(message, key []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return fmt.Sprintf("%x", expectedMAC)
}
