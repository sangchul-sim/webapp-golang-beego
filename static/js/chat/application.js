// https://github.com/joewalnes/reconnecting-websocket/
if (!"WebSocket" in window) {
    alert("WebSocket NOT supported by your Browser!");
}

var host = location.protocol.replace("http", "ws") + "//" + location.host + "/chat/ws";
var box = new ReconnectingWebSocket(host);
var chat = {
    "appendMessage": function(handle, message) {
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
    },
    "getUserInfo": function (user_id) {
        for (var i in chat.users) {
          var user = chat.users[i];

          if (user.id == user_id) {
              return user;
          }
        }
    },
    "users": [
        {"name": "1번유저", "id": 1},
        {"name": "2번유저", "id": 2},
        {"name": "3번유저", "id": 3},
        {"name": "4번유저", "id": 4},
        {"name": "5번유저", "id": 5},
        {"name": "6번유저", "id": 6},
    ],
    "rooms": [
        {"name": "room1", "id": 1},
        {"name": "room2", "id": 2},
        {"name": "room3", "id": 3},
        {"name": "room4", "id": 4},
        {"name": "room5", "id": 5},
        {"name": "room6", "id": 6},
    ]
};

// https://developer.mozilla.org/ko/docs/WebSockets/Writing_WebSocket_client_applications 참고
// 서버로부터 데이터 수신
box.onmessage = function($event) {
    var data = JSON.parse($event.data);

    if (data.hasOwnProperty("handle")) {
        if (data.handle == "system.message" || data.handle == "room.message") {
            if (data.handle == "room.message") {
                // socket.broadcast.to(room_id).emit() 구현
                // 현재 입장한 방에만 메시지를 보낸다.
                var selectedRoom = $("#room option:selected").val();
                if (selectedRoom == data.room_id) {
                    var user = chat.getUserInfo(data.user_id);
                    chat.appendMessage(user.name, data.message);
                }
            }
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
    $scope.users = chat.users;
    $scope.rooms = chat.rooms;
    $scope.selectedUser = $scope.users[0];
    $scope.selectedRoom = $scope.rooms[0];

    $scope.roomMessage = function($event) {
        $event.preventDefault();
        var handle = "room.message";
        var user_id = $scope.selectedUser.id;
        var room_id = $scope.selectedRoom.id;
        var message = $scope.message;
        if (handle == "" ||
            user_id == "" ||
            room_id == "" ||
            message == ""
        ) {
            return false;
        }

        // 서버에 데이터 전송하기
        var data = {
            "user_id": user_id,
            "room_id": room_id,
            "message": message
        };
        console.log("on.submit")
        console.dir(data);

        $scope.sendMessage(handle, data);
    };

    $scope.systemMessage = function($event) {
        var handle = "system.message";
        var data = {
            "user_id": "system",
            "message": "이것은 system 메시지 입니다."
        };
        console.log("on.submit")
        console.dir(data);

        $scope.sendMessage(handle, data);
    };

    $scope.focusMessage = function() {
        $("#input-message").focus();
    };

    $scope.sendMessage = function(handle, data) {
        socketSend(handle, data);
        $scope.message = ""
        $scope.focusMessage();
    };

    var socketSend = function(handle, data) {
        data.handle = handle;
        box.send(JSON.stringify(data));
    };

    $scope.changeRoom = function($event) {
        var handle = "room.in";
        var data = {
            "room_id":$scope.selectedRoom.id,
            "user_id":$scope.selectedUser.id,
        };
        socketSend(handle, data);
        $scope.message = ""
        $scope.focusMessage();
    };

    $scope.changeUser = function($event) {
      var handle = "room.in";
      var data = {
          "room_id":$scope.selectedRoom.id,
          "user_id":$scope.selectedUser.id,
      };
      socketSend(handle, data);
      $scope.message = ""
      $scope.focusMessage();
    };
});
