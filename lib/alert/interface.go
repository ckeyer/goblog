package alert

type Alerter interface {
	Alert(int, string) error
}

const (
	ALERT_LEVEL_WARING = iota
	ALERT_LEVEL_MINOR
	ALERT_LEVEL_MAJOR
	ALERT_LEVEL_CRITICAL
	ALERT_LEVEL_LEAVE_MSG
)

var (
	AlertTitle map[int]string
)

func init() {
	monitor_title := "监控告警信息"
	blog_title := "Blog"
	AlertTitle = make(map[int]string)
	AlertTitle[ALERT_LEVEL_WARING] = `【警告】` + monitor_title
	AlertTitle[ALERT_LEVEL_MINOR] = `【次要】` + monitor_title
	AlertTitle[ALERT_LEVEL_MAJOR] = `【主要】` + monitor_title
	AlertTitle[ALERT_LEVEL_CRITICAL] = `【严重】` + monitor_title

	AlertTitle[ALERT_LEVEL_LEAVE_MSG] = `【留言】` + blog_title
}
