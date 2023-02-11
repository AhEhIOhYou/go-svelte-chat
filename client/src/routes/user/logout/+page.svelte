<script lang="ts">
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import { closeWebSocketConnection } from '@/lib/services/ws-service';
	closeWebSocketConnection();
	removeItemInLS('userDetails');
	window.location.href = '/';
</script>
