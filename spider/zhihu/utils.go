package zhihu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"github.com/golang/glog"
)

const (
	userAgent    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36"
	baseZhihuURL = "https://www.zhihu.com"
	pageSize     = 20
)

var (
	reQuestionURL    = regexp.MustCompile("^(http|https)://www.zhihu.com/question/[0-9]{8}$")
	reCollectionURL  = regexp.MustCompile("^(http|https)://www.zhihu.com/collection/[0-9]{8,9}$") // bugfix: for private collection
	reTopicURL       = regexp.MustCompile("^(http|https)://www.zhihu.com/topic/[0-9]{8}$")
	reGetNumber      = regexp.MustCompile(`([0-9])+`)
	reAvatarReplacer = regexp.MustCompile(`_(s|xs|m|l|xl|hd).(png|jpg)`)
	reIsEmail        = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

func validQuestionURL(value string) bool {
	return reQuestionURL.MatchString(value)
}

func validCollectionURL(value string) bool {
	return reCollectionURL.MatchString(value)
}

func validTopicURL(value string) bool {
	return reTopicURL.MatchString(value)
}

func reMatchInt(raw string) int {
	matched := reGetNumber.FindStringSubmatch(raw)
	if len(matched) == 0 {
		return 0
	}
	rv, _ := strconv.Atoi(matched[0])
	return rv
}

func validateAvatarSize(size string) bool {
	for _, x := range []string{"s", "xs", "m", "l", "xl", "hd"} {
		if size == x {
			return true
		}
	}
	return false
}

func replaceAvatarSize(origin string, size string) string {
	return reAvatarReplacer.ReplaceAllString(origin, fmt.Sprintf("_%s.$2", size))
}

func isEmail(value string) bool {
	return reIsEmail.MatchString(value)
}

func newHTTPHeaders(isXhr bool) http.Header {
	headers := make(http.Header)
	headers.Set("Accept", "*/*")
	headers.Set("Connection", "keep-alive")
	headers.Set("Host", "www.zhihu.com")
	headers.Set("Origin", "http://www.zhihu.com")
	headers.Set("Pragma", "no-cache")
	headers.Set("User-Agent", userAgent)
	headers.Set("authorization","Bearer Mi4wQUFDQVRzNG5BQUFBQU1LcEx4UGtDaGNBQUFCaEFsVk5WSm1qV1FCaWN2VGtxRU4tNWxYY1Q0Z3JucFM5Qi1KNGRR|1501301844|85596a24b9e4d91aadd03098de887166321e1f8a")
	if isXhr {
		headers.Set("X-Requested-With", "XMLHttpRequest")
	}
	return headers
}

func strip(s string) string {
	return strings.TrimSpace(s)
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("获取 CWD 失败：" + err.Error())
	}
	return cwd
}

func save(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0666)
}

func saveString(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}

func openCaptchaFile(filename string) error {
	glog.Info("调用外部程序渲染验证码……")
	var args []string
	switch runtime.GOOS {
	case "linux":
		args = []string{"xdg-open", filename}
	case "darwin":
		args = []string{"open", filename}
	case "freebsd":
		args = []string{"open", filename}
	case "netbsd":
		args = []string{"open", filename}
	case "windows":
		var (
			cmd      = "url.dll,FileProtocolHandler"
			runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
		)
		args = []string{runDll32, cmd, filename}
	default:
		fmt.Printf("无法确定操作系统，请自行打开验证码 %s 文件，并输入验证码。", filename)
	}

	glog.Info("Command: %s", strings.Join(args, " "))

	err := exec.Command(args[0], args[1:]...).Run()
	if err != nil {
		return err
	}

	return nil
}

func readCaptchaInput() string {
	var captcha string
	fmt.Print("请输入验证码：")
	fmt.Scanf("%s", &captcha)
	return captcha
}

func makeZhihuLink(path string) string {
	return urlJoin(baseZhihuURL, path)
}

func urlJoin(base, path string) string {
	if strings.HasSuffix(base, "/") {
		base = strings.TrimRight(base, "/")
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimLeft(path, "/")
	}
	return base + "/" + path
}



