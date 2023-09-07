// client.js

const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/ws');

ws.on('open', () => {
  console.log('Connessione WebSocket aperta.');
  // Invia un messaggio al server
  ws.send('Ciao, server!');
});

ws.on('message', (message) => {
  console.log('Messaggio dal server:', message);
});

ws.on('close', () => {
  console.log('Connessione WebSocket chiusa.');
});
