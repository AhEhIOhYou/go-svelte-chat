<script lang="ts">
	import { getContacts } from '@/lib/services/api-service';
	import { sendWebSocketEvent } from '@/lib/services/ws-service';

	export let userID = null;
	let participantsID = [];
	let chatName = '';

	async function loadContacts() {
		let contactsRes = await getContacts(userID);
		return contactsRes.contacts;
	}

	let contacts = loadContacts();

	async function handleCreateChat(e) {
		if (participantsID.length < 2 || chatName == '') {
			throw new Error('Chat create error');
		}
		participantsID.push(userID);
		
		const payload = {
			name: chatName,
			participants: participantsID,
		}
		sendWebSocketEvent("new-chat", payload);
	}
</script>

<div class="contacts">
	<div class="chat-part">
		{#await contacts}
			<div>Loading...</div>
		{:then contacts}
			{#if !contacts}
				<div class="contact">
					<div class="name">No contacts</div>
				</div>
			{:else}
				{#each contacts as contact}
					<div class="contact">
						<label
							><input
								type="checkbox"
								bind:group={participantsID}
								value={contact.contactUserID}
							/>{contact.contactUsername}</label
						>
					</div>
				{/each}
			{/if}
		{/await}
	</div>
	<div class="chat-name">
		<label for="chat-name">Chat name:</label>
		<input id="chat-name" bind:value={chatName} />
	</div>
	<button on:click={(e) => handleCreateChat(e)} type="button">Create Chat</button>
</div>

<style lang="scss">
</style>
