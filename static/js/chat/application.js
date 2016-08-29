// https://github.com/joewalnes/reconnecting-websocket/
if ("WebSocket" in window) {
    var host = location.protocol.replace("http", "ws") + "//" + location.host + "/chat/ws";
    var box = new ReconnectingWebSocket(host);

    function systemMessage() {
        var message = {
            handle: "system.message",
            text: "이것은 클라이언트에서 보내는 메시지 입니다."
        }
        box.send(JSON.stringify(message));
    }

    function appendMessage(handle, message) {
        var html = '<div>';
        html += '<dl class="dl-horizontal">';
        html += '<dt><span class="glyphicon glyphicon-user" aria-hidden="true"></span><span>' + handle + '</span></dt>';
        html += '<dd><span class="label label-info">' + message + '</span></dd>';
        html += '</dl>';
        html += '</div>'
        $("#chat-text").append(html);

        $("#chat-text").stop().animate({
            scrollTop: $('#chat-text')[0].scrollHeight
        }, 800);
    }

    // https://developer.mozilla.org/ko/docs/WebSockets/Writing_WebSocket_client_applications 참고
    // 서버로부터 데이터 수신
    box.onmessage = function(message) {
        var data = JSON.parse(message.data);
        console.log("box.onmessage")
        console.dir(data);

        if (data.hasOwnProperty("handle")) {
            if (data.handle == "system.message" || data.handle == "room.message") {
                appendMessage(data.handle, data.text);
            }
        }
    };

    // 연결 종료
    box.onclose = function() {
        console.log('box closed');
        this.box = new WebSocket(box.url);
    };

    $("#input-form").on("submit", function(event) {
        event.preventDefault();
        var handle = $("#input-handle")[0].value;
        var user = $("#input-user")[0].value;
        var room = $("#input-room")[0].value;
        var text = $("#input-text")[0].value;
        if (handle == "" ||
            user == "" ||
            room == "" ||
            text == ""
        ) {
            return false;
        }

        // 서버에 데이터 전송하기
        var message = {
            handle: handle,
            user: user,
            room: room,
            text: text
        };
        console.log("on.submit")
        console.dir(message);

        box.send(JSON.stringify(message));
        $("#input-text")[0].value = "";
    });
} else {
    alert("WebSocket NOT supported by your Browser!");
}
