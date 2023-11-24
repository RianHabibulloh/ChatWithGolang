let chatSocket;
function connect() {
    chatSocket = new WebSocket('ws://' + window.location.host + '/ws/chat/');

    chatSocket.onmessage = function(e) {
        let data = JSON.parse(e.data);
        let li = document.createElement('li');
        li.textContent = data.message;
        document.querySelector('#chat ul').appendChild(li);
    };

    chatSocket.onclose = function(e) {
        console.error('Chat socket closed unexpectedly');
    };
}
document.querySelector('#message-form').addEventListener('submit', function(e) {
    e.preventDefault();
    let input = document.querySelector('#message');
    let message = input.value;
    chatSocket.send(JSON.stringify({'message': message}));
    input.value = '';
});

connect();