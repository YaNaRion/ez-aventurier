export interface UserStatusPayload {
	userId: string;
	online: boolean;
}

export interface LoginRequestPayload {
	uniqueID: string;
}

export interface LoginRequestPayload {
	uniqueID: string;
	isConnection: boolean;
}
