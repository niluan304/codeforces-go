package leetcode

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/levigross/grequests"
)

func Test_login(t *testing.T) {
	// TODO get error: POST https://leetcode.cn/accounts/login/ return code 500
	// found the source: generator.go:80
	// 	if !resp.Ok {
	//		return nil, fmt.Errorf("POST %s return code %d", loginURL, resp.StatusCode)
	//	}
	session, err := login(username, password)
	//session, err := login("zwei_elen_2023-10-03", "4Wc9FcuvtSshwfWP") // 临时注册的账号与密码
	if err != nil {
		t.Error(err)
	}
	t.Log(session)
}

// id > 0，指定具体的一场周赛
// id = 0，指定下一场或当前正在进行的周赛
// id < 0，指定上 |id| 场周赛（例如 id = -1 表示最近的一场结束的周赛）
func Test_genLeetcode_weekly(t *testing.T) {
	var (
		contestID = GetWeeklyContestID(362)
		tag       = GetWeeklyContestTag(contestID)
		dir       = fmt.Sprintf("../../../leetcode_zwei_elen/weekly/%d/", contestID) // 自定义生成目录
	)

	// 临时账号的 Cookie LEETCODE_SESSION
	var leetcodeSession = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNjYwMjUxMyIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNGYyZGE4MDljYWUyNWZlMTM4NjM3ZmMyMGJkMTRiYzZjMjc2OTA5MDEzZjRmOGM1MjNkZjMwYzdjY2MxZjViYyIsImlkIjo2NjAyNTEzLCJlbWFpbCI6IiIsInVzZXJuYW1lIjoiendlaV9lbGVuXzIwMjMtMTAtMDMiLCJ1c2VyX3NsdWciOiJ6d2VpX2VsZW5fMjAyMy0xMC0wMyIsImF2YXRhciI6Imh0dHBzOi8vYXNzZXRzLmxlZXRjb2RlLmNuL2FsaXl1bi1sYy11cGxvYWQvdXNlcnMvendlaV9lbGVuXzIwMjMtMTAtMDMvYXZhdGFyXzE2OTYzNDM5MTQucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2OTY1MTM3MjcuNDQxMywiZXhwaXJlZF90aW1lXyI6MTY5OTAzODAwMCwidmVyc2lvbl9rZXlfIjoxfQ._A9OHtBKIAqARuXHz5AVuu3hDxbcD0qXes_Sxv08Qk0"

	err := GenLeetCodeTestsBySession(NewSession(leetcodeSession), tag, true, dir, comment)
	if err != nil {
		t.Error(err)
	}
}

func Test_genLeetcode_biweekly(t *testing.T) {
	var (
		contestID = GetBiweeklyContestID(111)
		tag       = GetBiweeklyContestTag(contestID)
		dir       = fmt.Sprintf("../../../leetcode_zwei_elen/biweekly/%d/", contestID) // 自定义生成目录
	)

	err := GenLeetCodeTests(username, password, tag, true, dir, comment)
	if err != nil {
		t.Error(err)
	}
}

func NewSession(leetcodeSession string) *grequests.Session {
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		UseCookieJar: true,
	})
	u, _ := url.Parse("https://leetcode.cn/")
	session.HTTPClient.Jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "LEETCODE_SESSION",
			Value: leetcodeSession,
		},
	})
	return session
}

// 获取题目信息（含题目链接）
// contestTag 如 "weekly-contest-200"，可以从比赛链接中获取
func GenLeetCodeTestsBySession(session *grequests.Session, contestTag string, openWebPage bool, contestDir, customComment string) (err error) {
	//session, err := login(username, password)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("登录成功")
	////fmt.Println("登录成功", host, username)

	var problems []*problem
	for {
		problems, err = fetchProblemURLs(session, contestTag)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(500 * time.Millisecond)
	}

	if len(problems) == 0 {
		return nil
	}

	customComment = updateComment(customComment)

	for _, p := range problems {
		p.openURL = openWebPage
		p.customComment = customComment
		p.contestDir = contestDir
	}

	fmt.Println("题目链接获取成功，开始解析")
	return handleProblems(session, problems)
}
