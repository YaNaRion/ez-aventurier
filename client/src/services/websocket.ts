export enum WebSocketReadyState {
	CONNECTING = 0,
	OPEN = 1,
	CLOSING = 2,
	CLOSED = 3,
}

export interface WebSocketConfig {
	url: string;
	protocols?: string | string[];
	reconnectInterval?: number;
	maxReconnectAttempts?: number;
	heartbeatInterval?: number;
	heartbeatMessage?: string | object;
	autoConnect?: boolean;
}

export interface WebSocketMessage<T> {
	type: string;
	payload?: T;
	timestamp?: number;
	eventUniqueID?: string;
}

export type MessageHandler<T> = (message: T) => void;
export type ErrorHandler = (error: Event | Error) => void;
export type ConnectionHandler = () => void;

export const EventType = {
	LOGIN_REQUEST: "login.request",
	JOIN: "join",
	MESSAGE: "message",
};

export class WebSocketClient {
	private socket: WebSocket | null = null;
	private reconnectAttempts = 0;
	private reconnectTimeout: ReturnType<typeof setTimeout> | null = null;
	private heartbeatInterval: ReturnType<typeof setTimeout> | null = null;
	private isManuallyClosed = false;
	// ESlint disable
	private messageHandlers = new Map<string, MessageHandler<unknown>[]>();
	private connectionHandlers: ConnectionHandler[] = [];
	private disconnectionHandlers: ConnectionHandler[] = [];
	private errorHandlers: ErrorHandler[] = [];
	private config: WebSocketConfig

	constructor(config: WebSocketConfig) {
		this.config = {
			reconnectInterval: 3000,
			maxReconnectAttempts: 5,
			heartbeatInterval: 30000,
			heartbeatMessage: JSON.stringify({ type: 'heartbeat' }),
			autoConnect: true,
			...config,
		};

		if (this.config.autoConnect) {
			this.connect();
		}
	}

	/**
	 * Connect to WebSocket server
	 */
	public connect(): void {
		if (this.isConnected() || this.socket?.readyState === WebSocketReadyState.CONNECTING) {
			console.warn('WebSocket is already connecting or connected');
			return;
		}

		this.isManuallyClosed = false;

		try {
			this.socket = new WebSocket(this.config.url, this.config.protocols);
			this.setupEventListeners();
		} catch (error) {
			this.handleError(error as Error);
		}
	}

	/**
	 * Setup WebSocket event listeners
	 */
	private setupEventListeners(): void {
		if (!this.socket) return;

		this.socket.onopen = () => {
			console.log('WebSocket connected successfully');
			this.reconnectAttempts = 0;
			this.startHeartbeat();
			this.connectionHandlers.forEach(handler => handler());
		};

		this.socket.onmessage = (event) => {
			try {
				const message = this.parseMessage(event.data);
				this.handleMessage(message);
			} catch (error) {
				this.handleError(error as Error);
			}
		};

		this.socket.onclose = (event) => {
			console.log(`WebSocket disconnected: ${event.code} - ${event.reason}`);
			this.stopHeartbeat();
			this.disconnectionHandlers.forEach(handler => handler());

			if (!this.isManuallyClosed && this.shouldReconnect()) {
				this.scheduleReconnect();
			}
		};

		this.socket.onerror = (event) => {
			console.error('WebSocket error:', event);
			this.errorHandlers.forEach(handler => handler(event));
		};
	}

	/**
	 * Parse incoming message
	 */
	private parseMessage(data: string): WebSocketMessage {
		try {
			return JSON.parse(data) as WebSocketMessage;
		} catch {
			return { type: 'raw', payload: data };
		}
	}

	/**
	 * Handle incoming message
	 */
	private handleMessage(message: WebSocketMessage): void {
		const handlers = this.messageHandlers.get(message.type) || [];
		handlers.forEach(handler => handler(message.payload));
	}

	/**
	 * Send message to server
	 */
	public send<T>(type: string, payload?: T): boolean {
		if (!this.isConnected()) {
			console.error('Cannot send message: WebSocket is not connected');
			return false;
		}

		const message: WebSocketMessage<T> = {
			type,
			payload,
			timestamp: Date.now(),
			eventUniqueID: this.generateId(),
		};

		try {
			console.log(JSON.stringify(message));
			this.socket!.send(JSON.stringify(message));
			return true;
		} catch (error) {
			this.handleError(error as Error);
			return false;
		}
	}

	/**
	 * Send raw data
	 */
	public sendRaw(data: string | ArrayBuffer | Blob): boolean {
		if (!this.isConnected()) {
			console.error('Cannot send message: WebSocket is not connected');
			return false;
		}

		try {
			this.socket!.send(data);
			return true;
		} catch (error) {
			this.handleError(error as Error);
			return false;
		}
	}

	/**
	 * Subscribe to specific message types
	 */
	public on<T>(type: string, handler: MessageHandler<T>): () => void {
		if (!this.messageHandlers.has(type)) {
			this.messageHandlers.set(type, []);
		}
		this.messageHandlers.get(type)!.push(handler);

		// Return unsubscribe function
		return () => {
			const handlers = this.messageHandlers.get(type) || [];
			const index = handlers.indexOf(handler);
			if (index > -1) {
				handlers.splice(index, 1);
			}
		};
	}

	/**
	 * Subscribe to connection events
	 */
	public onConnect(handler: ConnectionHandler): void {
		this.connectionHandlers.push(handler);
	}

	/**
	 * Subscribe to disconnection events
	 */
	public onDisconnect(handler: ConnectionHandler): void {
		this.disconnectionHandlers.push(handler);
	}

	/**
	 * Subscribe to error events
	 */
	public onError(handler: ErrorHandler): void {
		this.errorHandlers.push(handler);
	}

	/**
	 * Close WebSocket connection
	 */
	public disconnect(code?: number, reason?: string): void {
		this.isManuallyClosed = true;
		this.stopHeartbeat();

		if (this.reconnectTimeout) {
			clearTimeout(this.reconnectTimeout);
			this.reconnectTimeout = null;
		}

		if (this.socket) {
			this.socket.close(code || 1000, reason);
		}
	}

	/**
	 * Schedule reconnection attempt
	 */
	private scheduleReconnect(): void {
		if (this.reconnectAttempts >= this.config.maxReconnectAttempts!) {
			console.error('Max reconnection attempts reached');
			return;
		}

		this.reconnectAttempts++;
		const delay = this.config.reconnectInterval! * Math.pow(1.5, this.reconnectAttempts - 1);

		console.log(`Attempting to reconnect in ${delay}ms (attempt ${this.reconnectAttempts})`);

		this.reconnectTimeout = setTimeout(() => {
			this.connect();
		}, delay);
	}

	/**
	 * Start heartbeat to keep connection alive
	 */
	private startHeartbeat(): void {
		if (!this.config.heartbeatInterval || this.config.heartbeatInterval <= 0) return;

		this.stopHeartbeat();

		this.heartbeatInterval = setInterval(() => {
			if (this.isConnected()) {
				if (typeof this.config.heartbeatMessage === 'string') {
					this.sendRaw(this.config.heartbeatMessage);
				} else {
					this.send('heartbeat', this.config.heartbeatMessage);
				}
			}
		}, this.config.heartbeatInterval);
	}

	/**
	 * Stop heartbeat
	 */
	private stopHeartbeat(): void {
		if (this.heartbeatInterval) {
			clearInterval(this.heartbeatInterval);
			this.heartbeatInterval = null;
		}
	}

	/**
	 * Check if should reconnect
	 */
	private shouldReconnect(): boolean {
		return this.reconnectAttempts < this.config.maxReconnectAttempts!;
	}

	/**
	 * Check if WebSocket is connected
	 */
	public isConnected(): boolean {
		return this.socket?.readyState === WebSocketReadyState.OPEN;
	}

	/**
	 * Get connection state
	 */
	public getState(): WebSocketReadyState {
		return this.socket?.readyState ?? WebSocketReadyState.CLOSED;
	}

	/**
	 * Generate unique ID for messages
	 */
	private generateId(): string {
		return Date.now().toString(36) + Math.random().toString(36).substr(2);
	}

	/**
	 * Handle errors
	 */
	private handleError(error: Error): void {
		console.error('WebSocket error:', error);
		this.errorHandlers.forEach(handler => handler(error));
	}

	/**
	 * Clean up resources
	 */
	public destroy(): void {
		this.disconnect();
		this.messageHandlers.clear();
		this.connectionHandlers = [];
		this.disconnectionHandlers = [];
		this.errorHandlers = [];
	}
}
