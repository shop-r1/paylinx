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
	"fmt"
	"sort"
	"strings"
)

func sign(params map[string]string, key string) string {
	s := mapToUrl(params, key)
	fmt.Println(s)
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func mapToUrl(params map[string]string, key string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var s string
	for i := range keys {
		s += keys[i] + "=" + params[keys[i]] + "&"
	}
	s += "key=" + key
	return s
}
