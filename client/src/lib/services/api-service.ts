const API_ENDPOINTS = 'https://localhost:8081/api';

export async function loginHTTPRequest(username, password) {
	const response = await fetch(`${API_ENDPOINTS}/users/login`, {
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
	const response = await fetch(`${API_ENDPOINTS}/users/register`, {
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
	const response = await fetch(`${API_ENDPOINTS}/users/is-username-available/${username}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}

export async function getUserByID(userID) {
	const response = await fetch(`${API_ENDPOINTS}/users/${userID}`, {
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

export async function searchUserByUserName(username) {
	const response = await fetch(`${API_ENDPOINTS}/users/search?username=${username}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}

export async function addToContacts(userID, contactID) {
	const response = await fetch(`${API_ENDPOINTS}/contacts/add/`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			"userID": userID,
			"contactUserID": contactID
		})
	});
	return await response.json();
}

export async function deleteContact(userID, contactID) {
	const response = await fetch(`${API_ENDPOINTS}/contacts/delete/`, {
		method: 'DELETE',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			"userID": userID,
			"contactUserID": contactID
		})
	});
	return await response.json();
}

export async function getContacts(userID) {
	const response = await fetch(`${API_ENDPOINTS}/contacts/${userID}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});
	return await response.json();
}