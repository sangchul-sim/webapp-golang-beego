// https://github.com/joewalnes/reconnecting-websocket/
if (!"WebSocket" in window) {
    alert("WebSocket NOT supported by your Browser!");
}

var host = location.protocol.replace("http", "ws") + "//" + location.host + "/chat/ws";
var box = new ReconnectingWebSocket(host);

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
            appendMessage(data.user, data.text);
        }
    }
};

// 연결 종료
box.onclose = function() {
    console.log('box closed');
    this.box = new WebSocket(box.url);
};

var app = angular.module("app", [
    // "ngAnimate"
    "ui.bootstrap"
]);

app.config(function($sceDelegateProvider) {
    $sceDelegateProvider.resourceUrlWhitelist(["self", "**"]);
});

app.controller("ChatController", function($scope, $rootScope, $log, $http, $timeout) {
    $scope.users = [
        "1번유저",
        "2번유저",
        "3번유저",
        "4번유저",
        "5번유저",
        "6번유저"
    ];
    $scope.rooms = [
        "room1",
        "room2",
        "room3",
        "room4",
        "room5",
    ];

    $scope.sendMessage = function($event) {
        $event.preventDefault();
        var handle = "room.message";
        var user = $scope.selectedUser;
        var room = $scope.selectedRoom;
        var text = $scope.message;
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
        $scope.message = ""
        $("#input-message").focus();
    };

    $scope.systemMessage = function($event) {
        var message = {
            handle: "system.message",
            user: "system",
            text: "이것은 system 메시지 입니다."
        };
        console.log("on.submit")
        console.dir(message);

        box.send(JSON.stringify(message));
        $scope.message = ""
        $("#input-message").focus();
    }
});
