// Import modules
import { ControllerService } from './service/controller.js';
import { EventType, WebSocketClient } from './service/websocket.js';
import ConnectionView from './component/connection.js';
import { LoginRequest, UserStatus } from './service/websocket_event.js';

interface GlobalState {
	controllerService: ControllerService;
	websocketClient: WebSocketClient;
	connectionComponent: ConnectionView;
}

function NewGlobalState(): GlobalState {
	let websocketClient: WebSocketClient = new WebSocketClient({
		url: 'ws://localhost:3000/ws',
		reconnectInterval: 5000,
		maxReconnectAttempts: 10,
		heartbeatInterval: 25000,
	});

	return {
		controllerService: new ControllerService(),
		websocketClient: websocketClient,
		connectionComponent: new ConnectionView(websocketClient),
	}
}

let global_state: GlobalState;

function InitialiseWebsocketHandler() {
	global_state.websocketClient.onConnect(() => {
		console.log('Connected to server');
	});

	global_state.websocketClient.onDisconnect(() => {
		console.log('Disconnected from server');
	});

	// Subscribe to specific message types
	const loginRespoinse = global_state.websocketClient.on<LoginRequest>('login.response', (message) => {
		console.log('New chat message:', message);
	});

	const unsubscribeUserStatus = global_state.websocketClient.on<UserStatus>('user.status', (status) => {
		console.log('User status update:', status);
	});

	// Error handling
	global_state.websocketClient.onError((error) => {
		console.error('WebSocket error:', error);
	});
}

// Initialize application
function init() {
	console.log('Application initializing...');
	global_state = NewGlobalState()
	InitialiseWebsocketHandler();
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
	document.addEventListener('DOMContentLoaded', init);
} else {
	init();
}

