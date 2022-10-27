/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/26 09:23:02
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/26 09:23:02
 */

package paylinx

import (
	"github.com/google/uuid"
	"os"
	"strings"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/shop-r1/royalpay"
	"github.com/spf13/cast"
)

func TestPay_CreateWechatOrder(t *testing.T) {
	type fields struct {
		MchID   int
		StoreID int
		Key     string
	}
	type args struct {
		orderID        string
		money          int
		body           string
		returnURL      string
		notifyURL      string
		spbillCreateIP string
		currency       royalpay.Currency
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test0",
			fields{
				MchID:   cast.ToInt(os.Getenv("MCH_ID")),
				StoreID: cast.ToInt(os.Getenv("STORE_ID")),
				Key:     os.Getenv("SECRET_KEY"),
			},
			args{
				orderID:        strings.ReplaceAll(uuid.New().String(), "-", ""),
				money:          14,
				body:           "test",
				returnURL:      "https://www.baidu.com",
				notifyURL:      "https://www.baidu.com",
				spbillCreateIP: "49.87.93.233",
				currency:       royalpay.CNY,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pay{
				MchID:   tt.fields.MchID,
				StoreID: tt.fields.StoreID,
				Key:     tt.fields.Key,
			}
			got, err := p.CreateWechatOrder(tt.args.money, tt.args.orderID,
				tt.args.body, tt.args.returnURL, tt.args.notifyURL, tt.args.spbillCreateIP, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWechatOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			litter.Dump(got)
		})
	}
}
