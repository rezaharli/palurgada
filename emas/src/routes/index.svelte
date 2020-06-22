<script>
	import Chart from '../components/d3/Chart.svelte';
	import points from '../datas/data.js';

	let fixedData = points.reverse()

	let showDays = 30;
</script>

<label>
	<input type=radio bind:group={showDays} value={30}>
	Show last 30 Days
</label>

<label>
	<input type=radio bind:group={showDays} value={60}>
	Show last 60 Days
</label>

<label>
	<input type=radio bind:group={showDays} value={-1}>
	Show All Days
</label>

<h1>Buy</h1>
<Chart points="{fixedData.map(v => ({ date: v.date_price, price: v.buy_price }))}" show="{showDays}" />

<h1>Sell</h1>
<Chart points="{fixedData.map(v => ({ date: v.date_price, price: v.sell_price }))}" show="{showDays}" />