package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vicanso/go-axios"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	apiRepoInfo = "https://api.github.com/repos/%s"
)

var repoReg = regexp.MustCompile(`\* \[[\s\S]+\]\(\S+\)`)
var githubRepoReg = regexp.MustCompile(`\(https://github.com/(\S+)\)`)
var ins = axios.NewInstance(&axios.InstanceConfig{
	Timeout: time.Minute,
})
var defaultLogger *zap.Logger

// 用于保存repo的star数量，用于在获取失败时使用
var originalRepoStars = map[string]int{}

// github api的认证参数
var clientToken *string

type Repo struct {
	content string
	star    int
}
type Repos []*Repo

func (rs Repos) Len() int {
	return len(rs)
}
func (rs Repos) Less(i, j int) bool {
	return rs[i].star < rs[j].star
}
func (rs Repos) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

type GithubRepoInfo struct {
	StargazersCount int `json:"stargazers_count"`
}

func init() {
	c := zap.NewProductionConfig()
	c.DisableCaller = true
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 只针对panic 以上的日志增加stack trace
	l, err := c.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		panic(err)
	}
	defaultLogger = l
}

// 获取star的数量
func getStarCount(name string) (count int, err error) {
	url := fmt.Sprintf(apiRepoInfo, name)
	// 如果出错则认为0
	headers := make(http.Header)
	headers.Add("Authorization", fmt.Sprintf("token %s", *clientToken))
	resp, err := ins.Request(&axios.Config{
		URL:     url,
		Headers: headers,
	})
	if err != nil {
		return
	}
	info := &GithubRepoInfo{}
	err = resp.JSON(info)
	if err != nil {
		return
	}
	count = info.StargazersCount
	return
}

// 按star对同一分类的repo重排
func arrangeByStar(lines []string) (result []string, err error) {
	result = make([]string, 0, len(lines))
	newCatStart := false

	var repos Repos
	for _, line := range lines {
		if repoReg.MatchString(line) {
			// 新的分类开始
			if !newCatStart {
				newCatStart = true
				repos = make(Repos, 0)
			}
			subs := githubRepoReg.FindAllStringSubmatch(line, -1)
			repo := &Repo{
				content: line,
			}
			if len(subs) > 0 && len(subs[0]) == 2 {
				name := subs[0][1]
				count, e := getStarCount(name)
				if e != nil {
					defaultLogger.Error("get repo info fail",
						zap.String("repo", name),
						zap.Error(e),
					)
					count = originalRepoStars[name]
				}

				repo.star = count
				if repo.star != 0 {
					stars := strconv.Itoa(repo.star)
					if repo.star >= 1000 {
						stars = fmt.Sprintf("%.1fK", float32(repo.star)/1000)
					}
					repo.content += fmt.Sprintf(" Stars:`%s`.", stars)
					originalRepoStars[name] = repo.star
				}
				defaultLogger.Info("get repo info success",
					zap.String("repo", name),
					zap.Int("star", count),
				)
			}
			repos = append(repos, repo)
			continue
		}
		// 如果开始不匹配，则该分类结束
		if newCatStart {
			newCatStart = false
			// 对repo按star从高至低排序
			sort.Sort(sort.Reverse(repos))
			for _, item := range repos {
				result = append(result, item.content)
			}
		}
		result = append(result, line)
	}
	return
}

func main() {
	// 如果没有不匹配，调用github api的频率限制较低
	clientToken = flag.String("clientToken", "", "github personal access token")
	flag.Parse()

	resp, err := ins.Get("https://raw.githubusercontent.com/rust-unofficial/awesome-rust/main/README.md")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(string(resp.Data), "\n", -1)

	result, err := arrangeByStar(lines)
	if err != nil {
		panic(err)
	}
	buf, _ := json.Marshal(originalRepoStars)
	if len(buf) != 0 {
		_ = ioutil.WriteFile(filepath.Join(os.TempDir()+"awesome-rust.json"), buf, 0600)
	}
	err = ioutil.WriteFile("./README.md", []byte(strings.Join(result, "\n")), 0600)
	if err != nil {
		panic(err)
	}
}
