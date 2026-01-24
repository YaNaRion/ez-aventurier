// Import modules
import { time } from 'node:console';
import { ControllerService } from './service/controller.js';
import { EventType, WebSocketClient } from './service/websocket.js';
import ConnectionView from './component/connection.js';

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

interface LoginRequest {
	uniqueID: string;
}

function InitialiseWebsocketHandler() {
	global_state.websocketClient.onConnect(() => {
		console.log('Connected to server');
	});

	global_state.websocketClient.onDisconnect(() => {
		console.log('Disconnected from server');
	});


	const element = document.createElement("button");
	element.onclick = () => {
		global_state.websocketClient.send<LoginRequest>(EventType.LOGIN_REQUEST, {
			uniqueID: "I Am unique",
		});
	}
	element.textContent = "CLICK TO SEND A LOGIN REQUEST";
	document.body.appendChild(element);

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
	InitialiseWebsocketHandler();
	ConnectionView.createHTML();
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
	document.addEventListener('DOMContentLoaded', init);
} else {
	init();
}

