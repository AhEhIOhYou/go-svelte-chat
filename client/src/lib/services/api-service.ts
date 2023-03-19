const API_ENDPOINTS = 'http://localhost:8081';

export async function loginHTTPRequest(username, password) {
	const response = await fetch(`${API_ENDPOINTS}/login`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			username,
			password,
		}),
	});
	return await response.json();
}

export async function registerHTTPRequest(username, password) {
	const response = await fetch(`${API_ENDPOINTS}/registration`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			username,
			password,
		}),
	});
	return await response.json();
}

export async function isUsernameAvailableHTTPRequest(username) {
	const response = await fetch(`${API_ENDPOINTS}/is-username-available/${username}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}

export async function userSessionCheckHTTPRequest(username) {
	const response = await fetch(`${API_ENDPOINTS}/user-session-check?userID=${username}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}


export async function getConversationBetweenUsers(toUserID, fromUserID) {
	const response = await fetch(`${API_ENDPOINTS}/chat/${fromUserID}/${toUserID}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}

export async function searchUserByNameHTTPRequest(username) {
	const response = await fetch(`${API_ENDPOINTS}/search-user?username=${username}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}