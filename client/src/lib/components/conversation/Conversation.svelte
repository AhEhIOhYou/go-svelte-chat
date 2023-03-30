<script lang="ts">
	import { getChatData } from '@/lib/services/api-service';
	import { sendWebSocketMessage } from '@/lib/services/ws-service';
	import { onMount, afterUpdate, beforeUpdate } from 'svelte';

	export let user = null;
	export let chatID = null;
	export let newMessage = null;

	let div;
	let autoscroll;
	let conversation = null;
	let messages = [];

	onMount(async () => {
		if (!chatID) return;
		conversation = await getChatData(chatID);
		console.log(conversation);
		
		messages = conversation.messages;
		div.scrollTo(0, div.scrollHeight);
	});

	beforeUpdate(() => {
		if (!chatID) return;
		autoscroll = div && div.offsetHeight + div.scrollTop > div.scrollHeight - 20;
	});

	afterUpdate(() => {
		if (!chatID) return;
		if (autoscroll) div.scrollTo(0, div.scrollHeight);
	});

	function handleKeydown(e) {
		if (e.key === 'Enter' && e.target.value !== '') {
			const payload = {
				chatID: chatID,
				fromUserID: user.ID,
				fromUserName: user.Name,
				message: e.target.value,
				createdAt: new Date().toISOString()
			};
			sendWebSocketMessage(payload);
			messages === null ? (messages = [payload]) : (messages = [...messages, payload]);
			e.target.value = '';
		}
	}

	$: if (newMessage && newMessage.chatID === chatID) {
		messages === null ? (messages = [newMessage]) : (messages = [...messages, newMessage]);
		newMessage = null;
	}
</script>

<div class="conv">
	<div class="chat-title">Chat Title</div>
	{#if !chatID}
		<div class="chat-title">Choose a contact to start a conversation</div>
	{:else}
		<div class="chat-scrollable" bind:this={div}>
			{#if !messages}
				<div class="be-first">No messages</div>
			{:else}
				{#each messages as message}
					<div class="message__outer">
						<div class="message__name">
							{message.fromUserName}
						</div>
						<div class="message__inner {message.fromUserID === user.ID ? 'me' : 'buddy'}">
							<div class="message__bubble">
								{message.message}
							</div>
						</div>
					</div>
				{/each}
			{/if}
		</div>
		<div class="chat-input">
			<input on:keydown={(e) => handleKeydown(e)} />
		</div>
	{/if}
</div>

<style lang="scss">
	.conv {
		display: flex;
		flex-direction: column;
		height: 100%;
		width: 100%;

		.chat-title {
			font-size: 1.5rem;
			font-weight: 600;
			padding: 1rem;
		}

		.chat-input {
			padding: 1rem;
			display: flex;
			align-items: center;
			justify-content: center;

			input {
				width: 100%;
				height: 40px;
				border: 1px solid #ccc;
				border-radius: 4px;
				padding: 0.5rem;
				font-size: 1rem;
			}
		}
	}

	.chat-scrollable {
		height: 300px;
		overflow-y: auto;

		.be-first {
			display: flex;
			align-items: center;
			justify-content: center;
			height: 100%;
			font-size: 1.2rem;
			font-weight: 600;
		}

		.message__outer {
			display: flex;
			flex-direction: row;

			.message__name {
				width: 130px;
				display: flex;
				align-items: center;
				justify-content: center;
				color: #999;
			}

			.message__inner {
				flex: 1;
				display: flex;
				flex-direction: row;
				word-break: break-all;

				.message__bubble {
					border-radius: 15px;
					padding: 0.6rem;
					margin: 0.2rem;
					max-width: calc(100% - 67px);
					overflow-wrap: break-word;
					border-end-start-radius: 4px;
				}

				&.me {
					.message__bubble {
						background: #007bff;
						color: #fff;
					}
				}

				&.buddy {
					.message__bubble {
						background: #e6e6e6;
						color: #000;
					}
				}
			}
		}
	}
</style>
