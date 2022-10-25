/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/26 06:41:50
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/26 06:41:50
 */

package paylinx

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"
)

func sign(params map[string]string, key string) string {
	s := mapToUrl(params, key)
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func mapToUrl(params map[string]string, key string) string {
	u := url.Values{}
	for k, v := range params {
		if v == "" {
			continue
		}
		u.Set(k, v)
	}
	s := u.Encode()
	s += "&key=" + key
	return s
}
