// import React from 'react';
import ConnectionContainer from './components/connection/connection';
import { EventType, WebSocketClient } from './services/websocket';
import './App.css';
import { useState } from 'react';
import type { LoginRequestPayload } from './services/websocket_event';
import UserPageComponent from './components/user-page/UserPage';
//
// Decorative components
const KnightThemeDecorations: React.FC = () => (
  <>
    <div className="knight-decor knight-left">♞</div>
    <div className="knight-decor knight-right">♘</div>
  </>
);

const Banner: React.FC = () => (
  <div className="knight-banner">
    <div className="banner-content">
      <div className="banner-title">Registre des ordres</div>
    </div>
  </div>
);

// Initialize WebSocket (adjust as needed)
const websocket = new WebSocketClient({
  url: 'ws://localhost:3000/ws',
  reconnectInterval: 5000,
  maxReconnectAttempts: 10,
  heartbeatInterval: 25000,
});

interface AppState {
  isConnected: boolean;
}

function App() {
  const [appState, setAppState] = useState<AppState | null>({
    isConnected: true,
  });


  websocket.on(EventType.LOGIN_RESPONSE, (payload: LoginRequestPayload) => {
    console.log(payload.uniqueID);
    setAppState({
      isConnected: true,
      ...appState,
    });
  });


  const user = {
    id: "ID",
    name: "Yann",
    userID: "1234678",
    team: 1,
    group: 1
  }

  return (
    <div className="App">
      <KnightThemeDecorations />
      <Banner />
      {appState?.isConnected ? (
        <UserPageComponent websocket={websocket} user={user} />
      ) : (
        <ConnectionContainer websocket={websocket} />
      )}
    </div>
  );
}

export default App;
