// src/components/StatusMessage.tsx
import React, { useEffect, useState } from 'react';

interface StatusMessageProps {
	message: string;
	type: 'success' | 'error' | 'info' | 'warning' | null;
	onDismiss?: () => void;
	autoDismiss?: boolean;
	dismissDelay?: number; // in milliseconds
}

const StatusMessage: React.FC<StatusMessageProps> = ({
	message,
	type,
	onDismiss,
	autoDismiss = true,
	dismissDelay = 3000
}) => {
	const [isVisible, setIsVisible] = useState(true);

	useEffect(() => {
		if (!message) {
			// setIsVisible(false);
			return;
		}

		// setIsVisible(true);

		if (autoDismiss && type === 'success' && onDismiss) {
			const timer = setTimeout(() => {
				setIsVisible(false);
				onDismiss();
			}, dismissDelay);

			return () => clearTimeout(timer);
		}
	}, [message, type, autoDismiss, dismissDelay, onDismiss]);

	if (!message || !isVisible || !type) {
		return null;
	}

	const getIcon = () => {
		switch (type) {
			case 'success': return '✅';
			case 'error': return '❌';
			case 'warning': return '⚠️';
			case 'info': return 'ℹ️';
			default: return '💬';
		}
	};

	const getClassName = () => {
		switch (type) {
			case 'success': return 'status-success';
			case 'error': return 'status-error';
			case 'warning': return 'status-warning';
			case 'info': return 'status-info';
			default: return 'status-default';
		}
	};

	const handleDismiss = () => {
		setIsVisible(false);
		if (onDismiss) {
			onDismiss();
		}
	};

	return (
		<div className={`status-message ${getClassName()} ${isVisible ? 'visible' : 'hidden'}`}>
			<div className="status-content">
				<span className="status-icon">{getIcon()}</span>
				<span className="status-text">{message}</span>
			</div>

			{(type !== 'success' || !autoDismiss) && onDismiss && (
				<button
					className="status-dismiss-btn"
					onClick={handleDismiss}
					aria-label="Dismiss message"
				>
					×
				</button>
			)}
		</div>
	);
};

export default StatusMessage;
