const WS_ENDPOINT = 'wss://localhost:8081';

let wsConnection = null;

export function connectToWebSocket(userID) {
	if (userID === "" && userID === null && userID === undefined) {
		return {
			message: "You need User ID to connect to the Chat server",
			webSocketConnection: null
		}
	} else if (!window["WebSocket"]) {
		return {
			message: "Your Browser doesn't support Web Sockets",
			webSocketConnection: null
		}
	}
	wsConnection = new WebSocket(`${WS_ENDPOINT}/ws?userID=${userID}`);
	return {
		message: "You are connected to Chat Server",
		wsConnection
	}
}

export function sendWebSocketMessage(messagePayload) {
	if (wsConnection === null) {
		return;
	}
	wsConnection.send(
		JSON.stringify({
			Name: 'message',
			Payload: messagePayload
		})
	);
}

export function sendWebSocketEvent(name, payload) {
	if (wsConnection === null) {
		return;
	}
	wsConnection.send(
		JSON.stringify({
			Name: name,
			Payload: payload
		})
	);
}

export function closeWebSocketConnection() {
	if (wsConnection === null) {
		return;
	}
	console.log('Closing the WebSocket Connection');
	wsConnection.close(1000, 'User has left the chat');
}

export function listenToWebSocketMessages(callback) {
	if (wsConnection === null) {
		return;
	}

	wsConnection.onclose = (event) => {
		callback(event);
	};

	wsConnection.onmessage = (event) => {
		try {
			const socketPayload = JSON.parse(event.data);
			if (!socketPayload.payload) {
				return
			}
			callback(socketPayload);
		} catch (error) {
			console.log(error)
			console.warn('Something went wrong while decoding the Message Payload')
		}
	};
}	