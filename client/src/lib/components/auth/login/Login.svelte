<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	let username: string = '';
	let password: string = '';
	let isNameAvailable: boolean = false;

	const dispatch = createEventDispatcher();

	const login = () => {
		dispatch('login', { username, password });
	};
</script>

<div class="login-conttainer">
	<div class="form-row">
		<label for="username">Username</label>
		<input bind:value={username} type="text" id="username" placeholder="Username" />
	</div>
	<div class="form-row">
		<label for="password">Password</label>
		<input bind:value={password} type="password" id="password" placeholder="Password" />
	</div>
	<div class="btn">
		<button type="submit" on:click={login}>Login</button>
	</div>
</div>

<style lang="scss">
	.login-conttainer {
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

	.btn {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin: 10px;
	}
</style>
