 
function setupWs() {
  let ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = () => {
    console.log("is connected");
  };
  return ws
 }

export    { setupWs };
