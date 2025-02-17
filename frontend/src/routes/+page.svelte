

<script>
	let counterValue = 0;

	async function incrementCounter() {
		const response = await fetch('/increment', { method: 'POST' });
		const data = await response.json();
		if (data.newValue !== undefined) {
			counterValue = data.newValue;
		}
	}
</script>

<h1>Counter: {counterValue}</h1>
<button on:click={incrementCounter}>Increment</button>
