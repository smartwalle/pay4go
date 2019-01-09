package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/pay4go"
	"github.com/smartwalle/xid"
	"net/http"
)

var (
	appID     = "2016073100129537"
	partnerID = "2088102169227503"

	// RSA2(SHA256)
	aliPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2MhEVUp+rRRyAD9HZfiS
g8LLxRAX18XOMJE8/MNnlSSTWCCoHnM+FIU+AfB+8FE+gGIJYXJlpTIyWn4VUMte
wh/4C8uwzBWod/3ilw9Uy7lFblXDBd8En8a59AxC6c9YL1nWD7/sh1szqej31VRI
2OXQSYgvhWNGjzw2/KS1GdrWmdsVP2hOiKVy6TNtH7XnCSRfBBCQ+LgqO1tE0NHD
DswRwBLAFmIlfZ//qZ+a8FvMc//sUm+CV78pQba4nnzsmh10fzVVFIWiKw3VDsxX
PRrAtOJCwNsBwbvMuI/ictvxxjUl4nBZDw4lXt5eWWqBrnTSzogFNOk06aNmEBTU
hwIDAQAB
-----END PUBLIC KEY-----`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAv8dXxi8wNAOqBNOh8Dv5rh0BTb5KNgk62jDaS536Z1cDqq2J
mpBYkBnzJXHAXEgBwPXgX8bGruMMjZKW8P4uv3Rvj8Am9ewWwUK2U7m2ZB3Oo9MW
tyYoiLGX1IA4FFenXzpPgm0WyzaeLX4yJ8j+hVrRbgwbZzb9Aq0MyepnK5PVoSPL
APXxvWrIBTok1+liughxwD/7R+ldaQQCtWC7nHBwOOChLkX6jenCOqi6LrTxJ4yc
GTWTctngFMJO4YtMmq/2zrw+ovNqmxHJQAZwuRFnKlZuFoEKPWyMGYtbvK9AWIfC
8ubn30O5F9kfLMIHwAHCh0UipPSbKDwQ2BnWswIDAQABAoIBAH7QyfkSsTRkC+Sf
MaGTd1qscXVAVQCAf/tSfLeuIqx9PL57fNfJhdbcYg2rt8EOGKLJtHKBFlcFawKf
IdMAslcGHtOXA+xxDucDP2AEGVkA4OkyJ/46bGlfzn/Fvc+t2s6812Du1DjSyCxb
G711SuFSGdVEikZpdUt0tVU7/LcyKAEZd45Ct+F9MvrPECbSsfODvTOVDHO2k42f
iwSzLPVmM4wVUc2xA15O87jtDhRiAK/RveQ7J2TWcarkyCR8J+bf5GGA79LdE3vR
Kr/HAk7INVX4T6U9QuDF30mqNRsloQbNGdvqO65nafNHvuVzUiqPdSX7XQwg/cOO
mhSsUbkCgYEA8BQXaHn3psHUZx8zEwQFVyd6rzxb+7jmVlUT+jG1pSiZ4WAWxxqx
YVXhn2dbfatDxWoGOMsrDM/Qp8g81nMG01jtmJr2RKFhAbQl93ipGvvaCNoJ8Lx7
HpFSq7dETcCCAE7tYMk0LlcVwxeaIUWakDyBHgEy4Zp6lLwdwsh115UCgYEAzH8/
E5dTOcYdcxk7HLupEC9MCb+FshZT5UIN9I7zLNljQX2O/8m2THb+oZUoy30RVot+
kYjh5H8M5CYiP0Kkm0Ovq5KC0loyt5SfzWbgwHEldQUVp8woE0YdaJzGB/UnmI9m
dJBON1t3qbMWjlguXOD8bfriDRuefaZd9oVSQycCgYBcz+ecxEoxdY2fsDgWid9m
qiSLylHlJr4lcg6fEsieaOvUbUlg/7jDYGgxL8v28Vbp4us02ZZzBYQs2QRsA1wI
KMDx1jaOobTW68YhvcviWqsX8PMW1kbislu7dsY5KMsZQ2oRmLdLku8e1OkJI9d1
G27vIpeBEC+DgJYgz05/YQKBgQCStWNiQbkihKBSF7LR3Uvf4Z6yi6V16xDLM8Vh
Q0DwVxEfRd3WYjcXynLJJ4J54kMTDMaD0GkHDaMI9taw/bWr8jZQZ67VDILAM68l
o/3v8fyGZFxx4kSJ905X48kqolWC3LYLQA/tJQDHTUUMX/T7CynuGQQdlUfyKu3U
Uzd+FwKBgHW9Nur4eTxK1nIOZyGgCqL1duYsJQcPWyIcRMTSjOoQZ5ZUhQZTw1Hd
2CW0Iu2fXExESTIjwXJ0ZJXnCgFU8acQX5vtItC1BlMaucw9XTx1RBCVQdTZ7DSX
vTlWbWwZHVDP85dioLE9mfo5+Hh3SmHDi3TaVXjxeJsUgHkRgOX7
-----END RSA PRIVATE KEY-----
`
)

func main() {
	var ap = pay4go.NewAliPay(appID, aliPublicKey, privateKey, false)
	ap.ReturnURL = "http://tw.smartwalle.tk/pay/return"
	ap.CancelURL = "http://tw.smartwalle.tk/pay/cancel"
	ap.NotifyURL = "http://tw.smartwalle.tk/pay/notify"

	var pp = pay4go.NewPayPal("AS8XSa9JrOJ3rf0kxVqCgRLIlMpgaKhLTShpYxISysR1VpnN6AMLfrvj-upOMuNkXdb9bTIzsFH4umB5", "ECA3_usif2DUgGxgcBTddOKgg2rbjUT7J3B3-Ud9z9y54AK9mYTDDFyadmMLSo1QOiO2rci99FSq1PbZ", false)
	pp.ReturnURL = "http://tw.smartwalle.tk/pay/return"
	pp.CancelURL = "http://tw.smartwalle.tk/pay/cancel"
	pp.WebHookId = "6WJ221414R474672F"

	var wp = pay4go.NewWXPal("wx20fa044851046bbf", "1v4h5g4s8u1x25tf451d025e10geagf2", "1299730801", false)
	wp.NotifyURL = "http://tw.smartwalle.tk/pay/notify"

	var ps = pay4go.NewService()
	ps.RegisterChannel(ap)
	ps.RegisterChannel(pp)
	ps.RegisterChannel(wp)

	http.HandleFunc("/pay/notify", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("notification", req.FormValue("channel"), req.FormValue("order_no"))

		var noti, err = ps.NotifyRequestHandler(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		bs, _ := json.Marshal(noti)
		fmt.Println(string(bs))
	})

	http.HandleFunc("/pay/cancel", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("cancel", req.FormValue("channel"), req.FormValue("order_no"))
	})

	http.HandleFunc("/pay/return", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("return", req.FormValue("channel"), req.FormValue("order_no"))

		trade, err := ps.ReturnRequestHandler(req)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		tradeByte, _ := json.Marshal(trade)
		w.Write(tradeByte)
	})

	http.HandleFunc("/pay", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		var channel = req.FormValue("c")
		var method = req.FormValue("m")

		var p = &pay4go.Order{}
		p.TradeMethod = method
		p.OrderNo = xid.NewXID().Hex()
		p.Currency = "USD"
		p.Discount = 10.33
		for i := 0; i < 3; i++ {
			p.AddProduct("test", "sku001", 1, 14.99, 0)
		}
		p.Timeout = 3

		var url, err = ps.CreatePayment(channel, p)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(channel, method, url)
		http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	})
	http.ListenAndServe(":5000", nil)
}
