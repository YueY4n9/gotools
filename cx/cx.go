package cx

import (
	"errors"
	"net/http"
	"net/url"
)

// SendMsg https://cx.super4.cn/ can get how to use it.
func SendMsg(appKey, title, content string) error {
	params := url.Values{}
	params.Add("appkey", appKey)
	params.Add("title", title)
	params.Add("content", content)
	resp, err := http.Get("https://cx.super4.cn/push_msg" + "?" + params.Encode())
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
