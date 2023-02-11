const WS_ENDPOINT = 'ws://localhost:8080';

import EventEmiter from 'eventemitter3';

let wsConnection = null;
const eventEmitter = new EventEmiter();

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

export function listenToWebSocketMessages() {
	if (wsConnection === null) {
		return;
	}

	wsConnection.onclose = (event) => {
		eventEmitter.emit('disconnect', event);
	};

	wsConnection.onmessage = (event) => {
		try {
			const socketPayload = JSON.parse(event.data);
			switch (socketPayload.eventName) {
				case 'chatlist-resp':
					if (!socketPayload.eventPayload) {
						return
					}
					eventEmitter.emit('chatlist-res', socketPayload.eventPayload);

					break;

				case 'disconnect':
					if (!socketPayload.eventPayload) {
						return
					}
					eventEmitter.emit('chatlist-res', socketPayload.eventPayload);

					break;

				case 'message-res':

					if (!socketPayload.eventPayload) {
						return
					}

					eventEmitter.emit('message-res', socketPayload.eventPayload);
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