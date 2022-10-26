/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/26 06:42:53
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/26 06:42:53
 */

package paylinx

import (
	"encoding/xml"
	"github.com/spf13/cast"
	"reflect"
)

type WechatCreatTransactionReq struct {
	XMLName        xml.Name `xml:"xml"`
	MchID          int      `xml:"mch_id"`
	StoreID        int      `xml:"store_id"`
	NonceStr       string   `xml:"nonce_str"`
	NotifyURL      string   `xml:"notify_url"`
	ReturnURL      string   `xml:"return_url,omitempty"`
	OutTradeNo     string   `xml:"out_trade_no"`
	Body           string   `xml:"body"`
	SpBillCreateIP string   `xml:"spbill_create_ip"`
	TotalFee       int      `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	TradeType      string   `xml:"trade_type"`
	Sign           string   `xml:"sign"`
}

func (e *WechatCreatTransactionReq) toMap(skipEmpty bool) map[string]string {
	obj1 := reflect.TypeOf(e).Elem()
	obj2 := reflect.ValueOf(e).Elem()

	var data = make(map[string]string)
	for i := 0; i < obj1.NumField(); i++ {
		if skipEmpty && obj2.Field(i).IsZero() {
			continue
		}
		data[obj1.Field(i).Tag.Get("xml")] = cast.ToString(obj2.Field(i).Interface())
	}
	return data
}

type WechatCreateTransactionResp struct {
	XMLName        xml.Name `xml:"xml"`
	MchID          int      `xml:"mch_id"`
	StoreID        int      `xml:"store_id"`
	NonceStr       string   `xml:"nonce_str"`
	NotifyURL      string   `xml:"notify_url"`
	ReturnURL      string   `xml:"return_url,omitempty"`
	OutTradeNo     string   `xml:"out_trade_no"`
	Body           string   `xml:"body"`
	SpBillCreateIP string   `xml:"spbill_create_ip"`
	TotalFee       int      `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	TradeType      string   `xml:"trade_type"`
	Sign           string   `xml:"sign"`
	PrepayID       string   `xml:"prepay_id"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`
	ResultCode     string   `xml:"result_code"`
	ErrCodeDes     string   `xml:"err_code_des"`
}

type AlipayCreateTransactionReq struct {
	XMLName       xml.Name `xml:"xml"`
	Partner       int      `xml:"partner"`
	StoreID       int      `xml:"store_id"`
	NonceStr      string   `xml:"nonce_str"`
	OutTradeNo    string   `xml:"out_trade_no"`
	TotalFee      int      `xml:"total_fee"`
	TransCurrency string   `xml:"trans_currency"`
	NotifyURL     string   `xml:"notify_url"`
	ReturnURL     string   `xml:"return_url,omitempty"`
	Subject       string   `xml:"subject"`
	Detail        string   `xml:"detail,omitempty"`
	Sign          string   `xml:"sign"`
}
