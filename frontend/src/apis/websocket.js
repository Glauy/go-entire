
let socket;

export function connectWebSocket() {
  socket = new WebSocket('ws://localhost:8080/ws');

  socket.addEventListener('open', () => {
    console.log('WebSocket connection opened');
  });

  socket.addEventListener('message', (event) => {
    const message = JSON.parse(event.data);
    if (message.type === 'response' && message.id === 'greet') {
      console.log('Greet response:', message.data);
    }
  });

  socket.addEventListener('close', () => {
    console.log('WebSocket connection closed');
  });
}

export function greet(name) {
  const request = {
    type: 'request',
    id: 'greet',
    method: 'Greet',
    arguments: [name],
  };

  socket.send(JSON.stringify(request));
}
