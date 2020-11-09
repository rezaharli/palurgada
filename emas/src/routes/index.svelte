<script>
	import Chart from '../components/d3/Chart.svelte';
	import points from '../datas/data.js';

	let fixedData = points.reverse()

	let weekDays = 7;

	let weeks = {
		show1week: weekDays * 1,
		show2week: weekDays * 2,
		show4week: weekDays * 4,
		show7week: weekDays * 7,
		show14week: weekDays * 14,
		show28week: weekDays * 28,
	}

	let showDays = 49;
</script>

<label>
	<input type=radio bind:group={showDays} value={weeks.show1week}>
	Show last 1 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={weeks.show2week}>
	Show last 2 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={weeks.show4week}>
	Show last 4 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={weeks.show7week}>
	Show last 7 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={weeks.show14week}>
	Show last 14 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={weeks.show28week}>
	Show last 28 weeks
</label>

<label>
	<input type=radio bind:group={showDays} value={-1}>
	Show All Days
</label>

<h1>5 gram Buy Price</h1>
<Chart points="{fixedData.map(v => ({ date: v.date_price, price: v.buy_price * 5 }))}" show="{showDays}" />

<h1>5 gram Sell Price</h1>
<Chart points="{fixedData.map(v => ({ date: v.date_price, price: v.sell_price * 5 }))}" show="{showDays}" />