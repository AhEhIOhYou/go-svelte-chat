const WS_ENDPOINT = 'ws://localhost:8080';

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
			eventName: 'message',
			eventPayload: messagePayload
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
			switch (socketPayload.eventName) {
				case 'chatlist-res':
					if (!socketPayload.eventPayload) {
						return
					}
					callback(socketPayload.eventPayload);
					break;
				case 'disconnect':
					if (!socketPayload.eventPayload) {
						return
					}
					callback(socketPayload.eventPayload);
					break;
				case 'message-res':
					if (!socketPayload.eventPayload) {
						return
					}
					callback(socketPayload.eventPayload);
					break;
				default:
					break;
			}
		} catch (error) {
			console.log(error)
			console.warn('Something went wrong while decoding the Message Payload')
		}
	};
}	