package wechat
import (
	"./config"
	"./wechat"
	"net/http"
)

func main() {
	wechat.Debug = true

	cfg := &wechat.WxConfig{
		Token:          config.Token,
		AppId:          config.AppId,
		Secret:         config.Secret,
		EncodingAESKey: config.EncodingAESKey,
	}

	app := wechat.New(cfg)
	app.SendText("@all", "Hello,World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = app.VerifyURL(w, r).NewText("客服消息1").Send().NewText("客服消息2").Send().NewText("查询OK").Reply()
	})

	_ = http.ListenAndServe("localhost:9090", nil)
}