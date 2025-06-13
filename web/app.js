const ws = new WebSocket(`ws://${location.host}/ws`);
const chat = document.getElementById("chat");
const form = document.getElementById("form");
const input = document.getElementById("input");
const username = document.getElementById("username");

ws.onmessage = function (event) {
  const msg = JSON.parse(event.data);
  appendMessage(msg.username, msg.text, msg.timestamp);
};

form.addEventListener("submit", function (e) {
  e.preventDefault();
  if (!input.value || !username.value) return;

  const message = {
    username: username.value,
    text: input.value,
  };
  ws.send(JSON.stringify(message));
  input.value = "";
});

function appendMessage(user, text, timestamp) {
  const el = document.createElement("div");
  const time = timestamp ? ` [${timestamp}]` : "";
  el.textContent = `${user}${time}: ${text}`;
  chat.appendChild(el);
  chat.scrollTop = chat.scrollHeight;
}
