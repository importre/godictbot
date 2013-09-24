package godictbot

import (
	"appengine"
	"appengine/xmpp"
	"fmt"
	"net/http"
	"strings"
)

func init() {
	xmpp.Handle(handleChat)

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/_ah/xmpp/subscription/subscribe", aHandler)
	http.HandleFunc("/_ah/xmpp/subscription/subscribed", aHandler)
	http.HandleFunc("/_ah/xmpp/subscription/unsubscribe", aHandler)
	http.HandleFunc("/_ah/xmpp/subscription/unsubscribed", aHandler)
	http.HandleFunc("/_ah/xmpp/presence/available", aHandler)
	http.HandleFunc("/_ah/xmpp/presence/unavailable", aHandler)
	http.HandleFunc("/_ah/xmpp/presence/probe", aHandler)
	http.HandleFunc("/_ah/xmpp/error", aHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "godictbot")
}

func handleChat(c appengine.Context, m *xmpp.Message) {
	body := "뭔 단어인지 모르겠Go..."
	dicts := DictList(c, m.Body)

	if dicts != nil {
		data := make([]string, len(dicts))
		for i, dict := range dicts {
			data[i] = dict.String()
		}
		body = strings.Join(data, "\n")
	}

	reply := &xmpp.Message{
		To:   []string{m.Sender},
		Body: body,
	}
	if err := reply.Send(c); err != nil {
		c.Errorf("Sending reply: %v", err)
	}
}

func aHandler(w http.ResponseWriter, r *http.Request) {
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	m := &xmpp.Message{
		To:   []string{"jaeweheo@gmail.com"},
		Body: "hello\nworld",
	}
	err := m.Send(c)
	if err != nil {
		// Send an email message instead...
		fmt.Fprint(w, err)
	}
}
