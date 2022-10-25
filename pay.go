/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/26 06:48:36
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/26 06:48:36
 */

package paylinx

import (
	"bytes"
	"encoding/xml"
	"github.com/google/uuid"
	"github.com/shop-r1/royalpay"
	"net/http"
)

const (
	WechatPayURL = "https://paylinx.cn/wxpay/gateway/unifiedorder/"
	AlipayURL    = "https://paylinx.cn/alipay/gateway/create/"
)

type Pay struct {
	MchID   int
	StoreID int
	Key     string
}

func NewPay(mchID, storeID int, key string) *Pay {
	return &Pay{
		MchID:   mchID,
		StoreID: storeID,
		Key:     key,
	}
}

// CreateWechatOrder 创建微信订单
func (p *Pay) CreateWechatOrder(money int,
	body, returnURL, notifyURL, spbillCreateIP string,
	currency royalpay.Currency) (*WechatCreateTransactionResp, error) {
	req := &WechatCreatTransactionReq{
		MchID:          p.MchID,
		StoreID:        p.StoreID,
		NotifyURL:      notifyURL,
		NonceStr:       uuid.New().String(),
		ReturnURL:      returnURL,
		Body:           body,
		SpBillCreateIP: spbillCreateIP,
		TotalFee:       money,
		FeeType:        string(currency),
		TradeType:      "JSAPI",
	}
	data := req.toMap(true)
	req.Sign = sign(data, p.Key)
	var buf bytes.Buffer
	err := xml.NewEncoder(&buf).Encode(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(WechatPayURL, "application/xml", &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData WechatCreateTransactionResp
	err = xml.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}
