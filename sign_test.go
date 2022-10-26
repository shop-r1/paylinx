/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/26 06:46:52
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/26 06:46:52
 */

package paylinx

import "testing"

func Test_sign(t *testing.T) {
	type args struct {
		paramsKv map[string]string
		key      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test01",
			args{
				paramsKv: map[string]string{
					"mch_id":   "aaa",
					"pb":       "bbb",
					"pc":       "ccc",
					"bizMchId": "6765",
				},
				key: "CDjAiTwflJzKTyMJ2C3Vuk5Ilv7AerES",
			},
			"12EA53784D1C0814BB3C1AE9B50ACE9C",
		},
		{
			"test02",
			args{
				paramsKv: map[string]string{
					"pa":       "aaa",
					"pb":       "bbb",
					"pc":       "ccc",
					"bizMchId": "7891",
				},
				key: "pSDsAZWvXCmZrbI2mIr0DdiPKErwNHWy",
			},
			"65BB240693036A8F7A83BECA0163E6AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sign(tt.args.paramsKv, tt.args.key)
			if got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
