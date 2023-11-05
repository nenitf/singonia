<!-- vim: set ft=html: -->
{{template "base" .}}

{{define "content"}}
<div class="container">
    <div class="row">
        <div class="col">
            <small>Logged as {{.User.Name}}</small>
        </div>
    </div>
</div>

<input id="input" type="text" />
<button onclick="send()">Send</button>
<pre id="output"></pre>
<script>
    var input = document.getElementById("input");
    var output = document.getElementById("output");
    var socket = new WebSocket("ws://localhost:8000/ws");

    socket.onopen = function () {
        output.innerHTML += "Status: Connected\n";
    };

    socket.onmessage = function (e) {
        output.innerHTML += "Server: " + e.data + "\n";
    };

    function send() {
        socket.send(input.value);
        input.value = "";
    }
</script>
{{end}}
