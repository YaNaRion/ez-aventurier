// src/hooks/useSession.ts
import { useState, useEffect, useCallback } from 'react';

export const useSession = () => {
	const [connectionId, setConnectionId] = useState('');
	const [sessionExpiry, setSessionExpiry] = useState<number | null>(null);
	const [timeLeft, setTimeLeft] = useState<number | null>(null);
	const [timerInterval, setTimerInterval] = useState<NodeJS.Timeout | null>(null);

	const getCookie = useCallback((name: string): string | null => {
		const nameEQ = name + "=";
		const ca = document.cookie.split(';');
		for (let i = 0; i < ca.length; i++) {
			let c = ca[i];
			while (c.charAt(0) === ' ') c = c.substring(1, c.length);
			if (c.indexOf(nameEQ) === 0) return decodeURIComponent(c.substring(nameEQ.length, c.length));
		}
		return null;
	}, []);

	const setCookie = useCallback((name: string, value: string, minutes: number) => {
		const expirationDate = new Date();
		expirationDate.setTime(expirationDate.getTime() + (minutes * 60 * 1000));
		document.cookie = `${name}=${encodeURIComponent(value)}; expires=${expirationDate.toUTCString()}; path=/; SameSite=Strict`;
	}, []);

	const deleteCookie = useCallback((name: string) => {
		document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
	}, []);

	const saveSession = useCallback((id: string, remember: boolean) => {
		const expiryTime = Date.now() + (10 * 60 * 1000); // 10 minutes

		// Set cookies
		setCookie('connectionId', id, 10);
		setCookie('sessionExpiry', expiryTime.toString(), 10);

		// Store in localStorage for additional persistence
		if (remember) {
			localStorage.setItem('lastConnectionId', id);
		} else {
			localStorage.removeItem('lastConnectionId');
		}

		setConnectionId(id);
		setSessionExpiry(expiryTime);
	}, [setCookie]);

	const validateSession = useCallback(() => {
		const savedConnectionId = getCookie('connectionId');
		const savedSessionExpiry = getCookie('sessionExpiry');

		if (savedConnectionId && savedSessionExpiry) {
			const expiryTime = parseInt(savedSessionExpiry);
			if (Date.now() < expiryTime) {
				setConnectionId(savedConnectionId);
				setSessionExpiry(expiryTime);
				return savedConnectionId;
			} else {
				clearSession();
			}
		}
		return null;
	}, [getCookie]);

	const clearSession = useCallback(() => {
		deleteCookie('connectionId');
		deleteCookie('sessionExpiry');
		localStorage.removeItem('lastConnectionId');

		if (timerInterval) {
			clearInterval(timerInterval);
			setTimerInterval(null);
		}

		setConnectionId('');
		setSessionExpiry(null);
		setTimeLeft(null);
	}, [deleteCookie, timerInterval]);

	const startSessionTimer = useCallback(() => {
		if (!sessionExpiry) return;

		// Clear existing interval
		if (timerInterval) {
			clearInterval(timerInterval);
		}

		const updateTimer = () => {
			const now = Date.now();
			const remaining = Math.max(0, sessionExpiry - now);
			setTimeLeft(remaining);

			if (remaining <= 0) {
				clearSession();
			}
		};

		updateTimer(); // Initial update
		const interval = setInterval(updateTimer, 1000);
		setTimerInterval(interval);

		// Cleanup
		return () => {
			if (interval) {
				clearInterval(interval);
			}
		};
	}, [sessionExpiry, timerInterval, clearSession]);

	// Cleanup on unmount
	useEffect(() => {
		return () => {
			if (timerInterval) {
				clearInterval(timerInterval);
			}
		};
	}, [timerInterval]);

	return {
		connectionId,
		setConnectionId,
		sessionExpiry,
		timeLeft,
		saveSession,
		validateSession,
		clearSession,
		startSessionTimer,
		getCookie,
		setCookie,
		deleteCookie
	};
};
