// Import modules
import { ControllerService } from './service/controller.js';
import { WebSocketClient } from './service/websocket.js';

interface GlobalState {
	controllerService: ControllerService;
	websocketClient: WebSocketClient;
}


const global_state: GlobalState = {
	controllerService: new ControllerService(),
	websocketClient: new WebSocketClient({
		url: 'ws://localhost:3000/ws',
		reconnectInterval: 5000,
		maxReconnectAttempts: 10,
		heartbeatInterval: 25000,
	}),
};

// Define message types
interface ChatMessage {
	userId: string;
	username: string;
	content: string;
	timestamp: number;
}

interface UserStatus {
	userId: string;
	online: boolean;
}

function InitialiseWebsocketHandler() {
	global_state.websocketClient.onConnect(() => {
		console.log('Connected to server');
		document.getElementById('connection-status')!.textContent = 'Connected';
		document.getElementById('connection-status')!.className = 'status-connected';
	});

	global_state.websocketClient.onDisconnect(() => {
		console.log('Disconnected from server');
		document.getElementById('connection-status')!.textContent = 'Disconnected';
		document.getElementById('connection-status')!.className = 'status-disconnected';
	});

	// Subscribe to specific message types
	const unsubscribeChat = global_state.websocketClient.on<ChatMessage>('chat.message', (message) => {
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
	InitialiseWebsocketHandler();
	console.log(global_state);
	// Initialise top bar
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
	document.addEventListener('DOMContentLoaded', init);
} else {
	init();
}

