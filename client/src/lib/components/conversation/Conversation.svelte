<script lang="ts">
	import { sendWebSocketMessage } from '@/lib/services/ws-service';
	import { onMount, afterUpdate, beforeUpdate } from 'svelte';

	let div;
	let autoscroll;

	export let conversation;
	export let newMessage = null;
	let messages = conversation.messages ?? [];

	onMount(() => {
		div.scrollTo(0, div.scrollHeight);
	});

	beforeUpdate(() => {
		autoscroll = div && div.offsetHeight + div.scrollTop > div.scrollHeight - 20;
	});

	afterUpdate(() => {
		if (autoscroll) div.scrollTo(0, div.scrollHeight);
	});

	$: if (newMessage) {
		messages = messages.concat(newMessage);
		newMessage = null;
	}

	function handleKeydown(event) {
		if (event.key === 'Enter') {
			const text = event.target.value;
			if (!text) return;

			const payload = {
				fromUserID: conversation.details.toId,
				toUserID: conversation.details.fromId,
				message: text
			};

			messages = messages.concat(payload);
			sendWebSocketMessage(payload);
			event.target.value = '';
		}
	}
</script>

<div class="conv">
	<div class="chat-scrollable" bind:this={div}>
		{#each messages as message}
			<div class="message">
				<div class="message__outer">
					<div class="message__name">
						<span
							>{message.fromUserID == conversation.details.toId
								? 'me'
								: conversation.details.fromName}</span
						>
					</div>
					<div
						class="message__inner {message.fromUserID == conversation.details.toId
							? 'me'
							: 'buddy'} "
					>
						<div class="message__bubble">{message.message}</div>
						<div class="message__spacer" />
					</div>
				</div>
			</div>
		{/each}
	</div>
	<div class="input">
		<input on:keydown={handleKeydown} />
	</div>
</div>

<style lang="scss">
	.chat-scrollable {
		height: 300px;
		overflow-y: auto;

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
