// src/components/ConnectionForm.tsx
import React, { useState, useEffect } from 'react';
// import StatusMessage from '../status-message/StatusMessage';
import { EventType, type WebSocketClient } from '../../services/websocket';
import type { LoginRequestPayload } from '../../services/websocket_event';


const SessionInformation: React.FC = () => (
	<div className="connection-info">
		<div className="info-item">
			<i className="icon-shield">🛡️</i>
			<span>Les portes se fermeront dans 10 minutes</span>
		</div>
	</div>
);

interface ConnectionFormProps {
	initialConnectionId?: string;
	isLoading?: boolean;
	websocket: WebSocketClient;
}

const ConnectionForm: React.FC<ConnectionFormProps> = ({
	initialConnectionId = '',
	isLoading = false,
	websocket
}) => {
	const [connectionId, setConnectionId] = useState(initialConnectionId);
	const [rememberMe, setRememberMe] = useState(!!initialConnectionId);
	const [error, setError] = useState('');

	useEffect(() => {
		// setConnectionId(initialConnectionId);
		// setRememberMe(!!initialConnectionId);
	}, [initialConnectionId]);

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();

		websocket.send<LoginRequestPayload>(EventType.LOGIN_REQUEST, {
			uniqueID: connectionId
		});
	};

	const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const value = e.target.value;
		setConnectionId(value);

		// Clear error when user starts typing
		if (error && value.trim()) {
			setError('');
		}
	};

	return (
		<div className="connection-card">
			<div className="connection-header">
				<h1 className="connection-title">Code secret</h1>
				<p className="connection-subtitle">Veuillez entrer le code secret</p>
			</div>

			<div className="connection-body">
				<form id="connectionForm" className="connection-form" onSubmit={handleSubmit}>
					<div className="form-group">
						<label htmlFor="connectionId" className="form-label">
							Identifiant
						</label>
						<div className="input-wrapper">
							<input
								type="text"
								id="connectionId"
								className={`form-input ${error ? 'error' : ''}`}
								placeholder="Entrer votre code personnalisé"
								value={connectionId}
								onChange={handleInputChange}
								required
								autoComplete="off"
								autoFocus
								disabled={isLoading}
							/>
							<div className="input-icon">
								<svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
									<path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
									<circle cx="12" cy="7" r="4"></circle>
								</svg>
							</div>
						</div>
						{error && <div className="error-message">{error}</div>}
					</div>

					<div className="form-options">
						<label className="checkbox-label">
							<input
								type="checkbox"
								id="rememberMe"
								checked={rememberMe}
								onChange={(e) => setRememberMe(e.target.checked)}
								disabled={isLoading}
							/>
							<span className="checkbox-custom"></span>
							<span className="checkbox-text">
								Garder votre session enregistrée pour 10 minutes
							</span>
						</label>
					</div>

					<button
						type="submit"
						className="connect-button"
						disabled={isLoading || !connectionId.trim()}
					>
						{isLoading ? (
							<span className="button-loader">
								<div className="spinner"></div>
							</span>
						) : (
							<span className="button-text">⏎ Enter the Citadel</span>
						)}
					</button>
				</form>
				<SessionInformation />

				<div className="connection-footer">
					<p className="footer-text">
						Pour toutes questions, veuillez les poser au ...
					</p>
				</div>
			</div>
		</div>
	);
};

export default ConnectionForm;
