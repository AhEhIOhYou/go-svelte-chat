<script lang="ts">
	import { isUsernameAvailableHTTPRequest } from '@/lib/services/api-service';
	import { createEventDispatcher } from 'svelte';

	let username: string = '';
	let password: string = '';
	let nameAvailable = {
		available: false,
		message: ''
	};

	const dispatch = createEventDispatcher();

	const regisrtation = () => {
		dispatch('regisrtation', { username, password });
	};

	$: username && checkNameAvailability();

	async function checkNameAvailability() {
		const res = await isUsernameAvailableHTTPRequest(username);
		if (res) {
			nameAvailable.available = res.available;
			nameAvailable.message = res.message;
		} else {
			console.log('Something went wrong');
		}
	}
</script>

<div class="reg-container">
	<div class="form-row">
		<label for="username">Username</label>
		<input bind:value={username} type="text" id="username" placeholder="Username" />
	</div>
	<div class="form-row">
		<label for="password">Password</label>
		<input bind:value={password} type="password" id="password" placeholder="Password" />
	</div>
	<div class="btn">
		<button type="submit" on:click={regisrtation}>Registrate</button>
	</div>
	{#if username}
		{#if nameAvailable.available}
			<div class="form-ok">
				{nameAvailable.message}
			</div>
		{:else}
			<div class="form-dang">
				{nameAvailable.message}
			</div>
		{/if}
	{/if}
</div>

<style lang="scss">
	.reg-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
	}

	.form-row {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin: 10px;
	}

	.form-ok {
		color: green;
		padding: 5px;
	}

	.form-dang {
		color: red;
		padding: 5px;
	}

	.btn {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin: 10px;
	}
</style>
