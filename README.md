XMPP
====


Goal
----

구글톡 사전 봇 만들기

![exam]


XMPP Java API Overview
----------------------

### 기능

- 메시지 보내기 / 받기
- 채팅 초대 보내기
- 유저의 출석 / 상태


### Incomming Messages

request handler에 의해 처리 (web request와 유사)


### Types of Message in GAE

메시지 타입 | 가능여부
--- | ---
chat | **O**
error | X
groupchat | X
headline | X
normal | **O**

타입 참고 : [RFC3921](http://www.ietf.org/rfc/rfc3921.txt)


### Terms

term | desc
--- | ---
JID | Jabber ID
XML Stanza | XML package ?!


Sending chat messages
---------------------

language | doc
--- | ---
java | [link][1]
go | [link][2]

development server에서는 메시지를 보낼 수 없음  
로그로 찍힘


Handling incoming calls
-----------------------

- xmpp_message  
allows your application to receive chat messages from XMPP-compliant chat services.
- xmpp_subscribe  
allows subscriptions between users and your application for the purpose of exchanging data, such as chat messages, presence information, and status messages.
- xmpp_presence  
allows your application to determine a user's chat presence (available or unavailable).
- xmpp_error  
allows your application to receive error stanzas.

### in app.yaml

	inbound_services:
	- xmpp_message
	- xmpp_presence
	- xmpp_subscribe
	- xmpp_error


### Receiving chat messages

일단 XMPP 서비스를 사용하려면 필수로 아래를 설정해야 함

	inbound_services:
	- xmpp_message

이래 해놓고, 앱이 메시지를 받으면 앱 앤진은 HTTP POST request를 아래 URL에 생성함

	/_ah/xmpp/message/chat/

고로 위 URL에 맞는 핸들러를 구현해서 메시지를 처리할 수 있음  
위 경로는 자동으로 admin만 접근이 가능하도록 설정됨 (admin 설정 필수가 아니라는 말)

앞서 언급했듯, **chat**과 **normal** 타입만 받기 때문에, 다른 메시지는 무시됨(핸들러 호출 안됨)

language | doc
--- | ---
java | [link][3]
go | [link][4]


### Handling subscriptions

	/_ah/xmpp/subscription/subscribe/
	/_ah/xmpp/subscription/subscribed/
	/_ah/xmpp/subscription/unsubscribe/
	/_ah/xmpp/subscription/unsubscribed/


### Handling user presence

	/_ah/xmpp/presence/available/
	/_ah/xmpp/presence/unavailable/
	/_ah/xmpp/presence/probe/


### Handling error stanzas

	/_ah/xmpp/error/


XMPP addresses
--------------

	your_app_id@appspot.com
	
	or
	
	anything@your_app_id.appspotchat.com

`appspotchat.com` 도메인도 사용 가능









[1]: https://developers.google.com/appengine/docs/java/xmpp/#Java_Sending_chat_messages "send msg in java"
[2]: https://developers.google.com/appengine/docs/go/xmpp/#Go_Sending_chat_messages "send msg in go"
[3]: https://developers.google.com/appengine/docs/java/xmpp/#Java_Handling_incoming_calls "recv msg in java"
[4]: https://developers.google.com/appengine/docs/go/xmpp/#Go_Handling_incoming_calls "recv msg in go"
[exam]: https://lh4.googleusercontent.com/-dTB00peJeMo/UkGU62qDyMI/AAAAAAAAEnY/VUM9DnJns2g/w670-h484-no/xmpp.png
