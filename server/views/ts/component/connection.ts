export const ConnectionView: any = {
	createHTML: function(): void {
		// Add Google Fonts for knight theme
		const fontLink = document.createElement('link');
		fontLink.href = 'https://fonts.googleapis.com/css2?family=Cinzel:wght@400;600;700&family=Cinzel+Decorative:wght@700&display=swap';
		fontLink.rel = 'stylesheet';
		document.head.appendChild(fontLink);

		// Add knight theme CSS
		const styleLink = document.createElement('link');
		styleLink.href = 'connection-styles.css'; // Make sure this path is correct
		styleLink.rel = 'stylesheet';
		document.head.appendChild(styleLink);

		// Replace entire body with connection page
		document.body.innerHTML = `
            <!-- Knight decorative elements -->
            <div class="knight-decor knight-left">♞</div>
            <div class="knight-decor knight-right">♘</div>
            
            <!-- Banner -->
            <div class="knight-banner">
                <div class="banner-content">
                    <div class="banner-title">Knight's Realm</div>
                    <div class="banner-subtitle">Secure Citadel Access</div>
                </div>
            </div>
            
            <!-- Main Connection Container -->
            <div class="connection-container">
                <div class="connection-card">
                    <div class="connection-header">
                        <h1 class="connection-title">⚔️ Citadel Gate</h1>
                        <p class="connection-subtitle">Present your sigil to enter the realm</p>
                    </div>
                    
                    <div class="connection-body">
                        <form id="connectionForm" class="connection-form">
                            <div class="form-group">
                                <label for="connectionId" class="form-label">
                                    <i class="icon-user"></i> Sigil of Passage
                                </label>
                                <div class="input-wrapper">
                                    <input 
                                        type="text" 
                                        id="connectionId" 
                                        class="form-input"
                                        placeholder="Enter your unique sigil (e.g., KNGHT-123, SIR-LANCELOT)"
                                        required
                                        autocomplete="off"
                                        autofocus
                                    />
                                    <div class="input-icon">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                                            <circle cx="12" cy="7" r="4"></circle>
                                        </svg>
                                    </div>
                                </div>
                                <div class="input-hint">
                                    <i class="icon-info"></i> Your sigil must be unique to your station
                                </div>
                            </div>
                            
                            <div class="form-options">
                                <label class="checkbox-label">
                                    <input type="checkbox" id="rememberMe">
                                    <span class="checkbox-custom"></span>
                                    <span class="checkbox-text">Keep this gateway open for 10 minutes</span>
                                </label>
                            </div>
                            
                            <button type="submit" class="connect-button" id="connectBtn">
                                <span class="button-text">⏎ Enter the Citadel</span>
                                <span class="button-loader" id="buttonLoader" style="display: none;">
                                    <div class="spinner"></div>
                                </span>
                            </button>
                            
                            <div class="connection-status" id="connectionStatus"></div>
                        </form>
                        
                        <div class="connection-info">
                            <div class="info-item">
                                <i class="icon-shield"></i>
                                <span>The gateway will seal in 10 minutes for security</span>
                            </div>
                            <div class="info-item">
                                <i class="icon-clock"></i>
                                <span>Ancient runes (cookies) maintain your passage</span>
                            </div>
                        </div>
                    </div>
                    
                    <div class="connection-footer">
                        <p class="footer-text">
                            Seek the scribe at 
                            <a href="mailto:scribe@knightsrealm.com" class="footer-link">scribe@knightsrealm.com</a>
                            for guidance
                        </p>
                    </div>
                </div>
                
                <div class="session-timer" id="sessionTimer" style="display: none;">
                    <div class="timer-text">Gateway seals in: <span id="timerCountdown">10:00</span></div>
                    <div class="timer-progress">
                        <div class="timer-bar" id="timerBar"></div>
                    </div>
                </div>
            </div>
        `;

		// Initialize the connection functionality
		this.initializeConnection();

		// Check for existing session
		this.checkExistingSession();
	},

	initializeConnection: function(): void {
		const form = document.getElementById('connectionForm') as HTMLFormElement;
		const connectBtn = document.getElementById('connectBtn') as HTMLButtonElement;
		const buttonLoader = document.getElementById('buttonLoader') as HTMLDivElement;
		const buttonText = connectBtn.querySelector('.button-text') as HTMLSpanElement;
		const connectionStatus = document.getElementById('connectionStatus') as HTMLDivElement;

		form.addEventListener('submit', (e) => {
			e.preventDefault();
			this.handleConnection();
		});

		// Also handle button click for safety
		connectBtn.addEventListener('click', (e) => {
			e.preventDefault();
			this.handleConnection();
		});
	},

	handleConnection: function(): void {
		const connectionIdInput = document.getElementById('connectionId') as HTMLInputElement;
		const rememberMeCheckbox = document.getElementById('rememberMe') as HTMLInputElement;
		const connectBtn = document.getElementById('connectBtn') as HTMLButtonElement;
		const buttonLoader = document.getElementById('buttonLoader') as HTMLDivElement;
		const buttonText = connectBtn.querySelector('.button-text') as HTMLSpanElement;
		const connectionStatus = document.getElementById('connectionStatus') as HTMLDivElement;

		const connectionId = connectionIdInput.value.trim();

		if (!connectionId) {
			this.showStatus('The sigil field cannot be empty', 'error');
			connectionIdInput.focus();
			return;
		}

		// Validate connection ID format (alphanumeric and hyphens)
		const idRegex = /^[a-zA-Z0-9\-_]+$/;
		if (!idRegex.test(connectionId)) {
			this.showStatus('Sigil may only contain letters, numbers, hyphens, and underscores', 'error');
			return;
		}

		// Show loading state
		connectBtn.disabled = true;
		buttonText.style.opacity = '0';
		buttonLoader.style.display = 'block';

		// Simulate API call
		setTimeout(() => {
			this.setCookie(connectionId, rememberMeCheckbox.checked);
			this.showStatus('Sigil accepted! The gates open before you...', 'success');

			// Simulate redirect or next step
			setTimeout(() => {
				this.startSessionTimer();
				this.showConnectedUI(connectionId);
			}, 1000);

		}, 1500);
	},

	setCookie: function(connectionId: string, remember: boolean): void {
		const expiresInMinutes = 10; // Session duration in minutes
		const expirationDate = new Date();
		expirationDate.setTime(expirationDate.getTime() + (expiresInMinutes * 60 * 1000));

		// Set cookie with connection ID
		document.cookie = `connectionId=${encodeURIComponent(connectionId)}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;

		// Set session expiration time
		const sessionExpiry = Date.now() + (expiresInMinutes * 60 * 1000);
		document.cookie = `sessionExpiry=${sessionExpiry}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;

		// Store in localStorage for additional persistence
		if (remember) {
			localStorage.setItem('lastConnectionId', connectionId);
		} else {
			localStorage.removeItem('lastConnectionId');
		}
	},

	getCookie: function(name: string): string | null {
		const nameEQ = name + "=";
		const ca = document.cookie.split(';');
		for (let i = 0; i < ca.length; i++) {
			let c = ca[i];
			while (c.charAt(0) === ' ') c = c.substring(1, c.length);
			if (c.indexOf(nameEQ) === 0) return decodeURIComponent(c.substring(nameEQ.length, c.length));
		}
		return null;
	},

	deleteCookie: function(name: string): void {
		document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
	},

	checkExistingSession: function(): void {
		const connectionId = this.getCookie('connectionId');
		const sessionExpiry = this.getCookie('sessionExpiry');

		if (connectionId && sessionExpiry) {
			const expiryTime = parseInt(sessionExpiry);

			if (Date.now() < expiryTime) {
				// Valid session exists
				this.startSessionTimer();
				this.showConnectedUI(connectionId);
				return;
			} else {
				// Session expired, clear cookies
				this.deleteCookie('connectionId');
				this.deleteCookie('sessionExpiry');
			}
		}

		// Check for last connection ID in localStorage
		const lastConnectionId = localStorage.getItem('lastConnectionId');
		if (lastConnectionId) {
			const connectionIdInput = document.getElementById('connectionId') as HTMLInputElement;
			const rememberMeCheckbox = document.getElementById('rememberMe') as HTMLInputElement;
			connectionIdInput.value = lastConnectionId;
			rememberMeCheckbox.checked = true;
		}
	},

	startSessionTimer: function(): void {
		const sessionExpiry = this.getCookie('sessionExpiry');
		if (!sessionExpiry) return;

		const expiryTime = parseInt(sessionExpiry);
		const totalDuration = 10 * 60 * 1000; // 10 minutes in milliseconds

		const timerElement = document.getElementById('sessionTimer') as HTMLDivElement;
		const countdownElement = document.getElementById('timerCountdown') as HTMLSpanElement;
		const timerBar = document.getElementById('timerBar') as HTMLDivElement;

		timerElement.style.display = 'block';

		const updateTimer = () => {
			const now = Date.now();
			const timeLeft = Math.max(0, expiryTime - now);
			const minutes = Math.floor(timeLeft / (1000 * 60));
			const seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);

			// Update countdown display
			countdownElement.textContent = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;

			// Update progress bar
			const progress = (timeLeft / totalDuration) * 100;
			timerBar.style.width = `${progress}%`;

			// Change color based on time left
			if (timeLeft < 60000) { // Less than 1 minute
				timerBar.style.background = 'linear-gradient(90deg, #8b0000 0%, #ff0000 100%)';
			} else if (timeLeft < 300000) { // Less than 5 minutes
				timerBar.style.background = 'linear-gradient(90deg, #8b4513 0%, #ffd700 100%)';
			}

			if (timeLeft <= 0) {
				clearInterval(timerInterval);
				this.disconnectUser();
			}
		};

		// Update immediately and then every second
		updateTimer();
		const timerInterval = setInterval(updateTimer, 1000);
	},

	showConnectedUI: function(connectionId: string): void {
		const connectionCard = document.querySelector('.connection-card') as HTMLDivElement;

		connectionCard.innerHTML = `
            <div class="connection-header">
                <div class="success-icon" style="font-size: 48px; margin-bottom: 20px;">🏰</div>
                <h1 class="connection-title">Welcome, Noble Knight!</h1>
                <p class="connection-subtitle">The citadel welcomes <strong>${connectionId}</strong></p>
            </div>
            
            <div class="connection-body">
                <div class="user-info">
                    <div class="info-card">
                        <div class="info-icon">⚔️</div>
                        <div class="info-content">
                            <h3>Your Sigil</h3>
                            <p>${connectionId}</p>
                        </div>
                    </div>
                    
                    <div class="info-card">
                        <div class="info-icon">🕯️</div>
                        <div class="info-content">
                            <h3>Torch Burns For</h3>
                            <p id="sessionExpiryDisplay">10:00</p>
                        </div>
                    </div>
                </div>
                
                <div class="action-buttons">
                    <button class="action-btn primary-btn" id="refreshBtn">
                        <span>⏳ Rekindle Torch</span>
                    </button>
                    <button class="action-btn secondary-btn" id="disconnectBtn">
                        <span>🚪 Leave Citadel</span>
                    </button>
                </div>
                
                <div class="session-warning">
                    <i class="icon-warning"></i>
                    <span>The gates will seal in 10 minutes to protect the realm</span>
                </div>
            </div>
        `;

		// Add event listeners for connected UI
		document.getElementById('refreshBtn')?.addEventListener('click', () => {
			this.refreshSession();
		});

		document.getElementById('disconnectBtn')?.addEventListener('click', () => {
			this.disconnectUser();
		});
	},

	refreshSession: function(): void {
		const connectionId = this.getCookie('connectionId');
		if (connectionId) {
			this.setCookie(connectionId, true);
			this.showStatus('Torch rekindled! Your passage extends.', 'success');
		}
	},

	disconnectUser: function(): void {
		// Clear cookies
		this.deleteCookie('connectionId');
		this.deleteCookie('sessionExpiry');

		// Show disconnected message
		const connectionCard = document.querySelector('.connection-card') as HTMLDivElement;
		connectionCard.innerHTML = `
            <div class="connection-header">
                <div class="disconnect-icon" style="font-size: 48px; margin-bottom: 20px;">🏮</div>
                <h1 class="connection-title">Gate Sealed</h1>
                <p class="connection-subtitle">You have left the citadel</p>
            </div>
            
            <div class="connection-body">
                <div class="disconnect-message">
                    <p>The gates have closed behind you for the safety of the realm.</p>
                    <p>Present your sigil once more to re-enter.</p>
                </div>
                
                <button class="reconnect-btn" id="reconnectBtn">
                    <span>⚔️ Approach Gate</span>
                </button>
            </div>
        `;

		// Add reconnect event listener
		document.getElementById('reconnectBtn')?.addEventListener('click', () => {
			this.createHTML(); // Restart the connection screen
		});

		// Hide timer
		const timerElement = document.getElementById('sessionTimer') as HTMLDivElement;
		if (timerElement) {
			timerElement.style.display = 'none';
		}
	},

	showStatus: function(message: string, type: 'success' | 'error'): void {
		const connectionStatus = document.getElementById('connectionStatus') as HTMLDivElement;
		connectionStatus.textContent = message;
		connectionStatus.className = 'connection-status ' + type;

		// Auto-hide success messages
		if (type === 'success') {
			setTimeout(() => {
				connectionStatus.style.display = 'none';
			}, 3000);
		}
	}
};

// Export for use in other files
export default ConnectionView;
