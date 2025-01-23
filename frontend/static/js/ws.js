let ws;

function setupWs() {
  ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = () => {
    console.log("is connected");
  };
}

export    { setupWs };
