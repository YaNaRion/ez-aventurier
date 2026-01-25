// import React from 'react';
import ConnectionContainer from './components/connection/connection';
import { WebSocketClient } from './services/websocket';
import './App.css';

// Initialize WebSocket (adjust as needed)
const websocket = new WebSocketClient({
  url: 'ws://localhost:3000/ws',
  reconnectInterval: 5000,
  maxReconnectAttempts: 10,
  heartbeatInterval: 25000,
});

function App() {
  return (
    <div className="App">
      <ConnectionContainer websocket={websocket} />
    </div>
  );
}

export default App;
