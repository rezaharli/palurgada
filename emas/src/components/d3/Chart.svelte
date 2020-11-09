<script>
  import { scaleLinear } from "d3-scale";

  export let points;
  export let show = -1;

  //y -> price
  //x -> date

  const padding = { top: 20, right: 15, bottom: 20, left: 25 };

  let width = 700;
  let height = 300;

  $: ma14points = points
    .map(mapPointsByThresholdCallback(14))
    .filter((v) => v != null);
  $: ma84points = points
    .map(mapPointsByThresholdCallback(84))
    .filter((v) => v != null);

  $: filteredPoints = show == -1 ? points : points.slice(show * -1);
  $: filteredMa14Points = show == -1 ? ma14points : ma14points.slice(show * -1);
  $: filteredMa84Points = show == -1 ? ma84points : ma84points.slice(show * -1);

  $: minX = Math.min(...filteredPoints.map((v) => Date.parse(v.date) / 1000));
  $: maxX = Math.max(...filteredPoints.map((v) => Date.parse(v.date) / 1000));

  $: minY =
    Math.floor(
      Math.min(
        ...filteredPoints
          .concat(filteredMa14Points)
          .concat(filteredMa84Points)
          .map((v) => v.price)
      ) / 10000
    ) * 10000;
  $: maxY =
    Math.ceil(
      Math.max(
        ...filteredPoints
          .concat(filteredMa14Points)
          .concat(filteredMa84Points)
          .map((v) => v.price)
      ) / 10000
    ) * 10000;

  $: xScale = scaleLinear()
    .domain([minX, maxX])
    .range([padding.left, width - padding.right]);

  $: yScale = scaleLinear()
    .domain([Math.min.apply(null, yTicks), Math.max.apply(null, yTicks)])
    .range([height - padding.bottom, padding.top]);

  $: path = `M${filteredPoints
    .map((v) => `${xScale(Date.parse(v.date) / 1000)}, ${yScale(v.price)}`)
    .join("L")}`;
  $: path14 = `M${filteredMa14Points
    .map((v) => `${xScale(Date.parse(v.date) / 1000)}, ${yScale(v.price)}`)
    .join("L")}`;
  $: path84 = `M${filteredMa84Points
    .map((v) => `${xScale(Date.parse(v.date) / 1000)}, ${yScale(v.price)}`)
    .join("L")}`;

  $: area = `${path}L${xScale(maxX)},${yScale(0)}L${xScale(minX)},${yScale(
    0
  )}Z`;

  $: xTicks = filteredPoints.map((v) => Date.parse(v.date) / 1000).reverse();
  $: yTicks = (() => {
    let middlePrices = (range) => {
      let ps = [];
      for (let i = minY; i < maxY; i += range) {
        let mn = Math.floor(i / range) * range;
        if (mn > minY) ps.push(mn);

        let mx = Math.ceil(i / range) * range;
        if (mx < maxY) ps.push(mx);
      }

      return ps.filter((v, i, self) => {
        return self.indexOf(v) === i;
      });
    };

    return [minY, ...middlePrices(21000), maxY];
  })();

  function mapPointsByThresholdCallback(threshold) {
    return function (v, i) {
      var prices = points.map((v) => v.price).slice(i - threshold, i);
      var sum = prices.reduce((a, b) => a + b, 0);
      var avg = sum / threshold;

      return i < threshold
        ? null
        : {
            date: v.date,
            price: avg,
          };
    };
  }

  function timeConverter(UNIX_timestamp) {
    var a = new Date(UNIX_timestamp * 1000);
    var months = [
      "Jan",
      "Feb",
      "Mar",
      "Apr",
      "May",
      "Jun",
      "Jul",
      "Aug",
      "Sep",
      "Oct",
      "Nov",
      "Dec",
    ];
    var year = a.getFullYear();
    var month = months[a.getMonth()];
    var date = a.getDate();
    return [date, month, year];
  }
</script>

<div class="chart" bind:clientWidth="{width}" bind:clientHeight="{height}">
  <svg>
    <!-- y axis -->
    <g class="axis y-axis" transform="translate(0, {padding.top})">
      {#each yTicks as tick}
      <g
        class="tick tick-{tick}"
        transform="translate(0, {yScale(tick) - padding.bottom})"
      >
        <line x2="100%"></line>
        <text y="-4">{tick}</text>
      </g>
      {/each}
    </g>

    <!-- x axis -->
    <g class="axis x-axis">
      {#each xTicks as tick, i} {#if show == -1 } {#if i % 10 == 0 }
      <g
        class="tick tick-{ tick }"
        transform="translate({xScale(tick)},{height})"
      >
        <line y1="-{height}" y2="-{padding.bottom}" x1="0" x2="0"></line>

        {#if i == 0 }
        <text y="-2" x="25">{timeConverter(tick).join(" ")}</text>
        {:else} {#if timeConverter(tick)[0] == 1 }
        <text y="-2">
          {timeConverter(tick)[0] + " " + timeConverter(tick)[1]}
        </text>
        {:else}
        <text y="-2">{timeConverter(tick)[0]}</text>
        {/if} {/if}
      </g>
      {/if} {:else}
      <g
        class="tick tick-{ tick }"
        transform="translate({xScale(tick)},{height})"
      >
        <line y1="-{height}" y2="-{padding.bottom}" x1="0" x2="0"></line>

        {#if i == 0 }
        <text y="-2" x="25">{timeConverter(tick).join(" ")}</text>
        {:else} {#if timeConverter(tick)[0] == 1 }
        <text y="-2">
          {timeConverter(tick)[0] + " " + timeConverter(tick)[1]}
        </text>
        {:else}
        <text y="-2">{timeConverter(tick)[0]}</text>
        {/if} {/if}
      </g>
      {/if} {/each}
    </g>

    <!-- data -->
    <path class="path-line" d="{path}"></path>
    <path class="path-line path-line-14" d="{path14}"></path>
    <path class="path-line path-line-84" d="{path84}"></path>
  </svg>
</div>

<style>
  .chart,
  h2,
  p {
    width: 100%;
    margin-left: auto;
    margin-right: auto;
  }

  svg {
    position: relative;
    width: 100%;
    height: 300px;
    overflow: visible;
  }

  .tick {
    font-size: 0.725em;
    font-weight: 300;
  }

  .tick line {
    stroke: #aaa;
    stroke-dasharray: 2;
  }

  .tick text {
    fill: #666;
    text-anchor: start;
  }

  .tick.tick-0 line {
    stroke-dasharray: 0;
  }

  .x-axis .tick text {
    text-anchor: middle;
  }

  .path-line {
    fill: none;
    stroke: blue;
    stroke-linejoin: round;
    stroke-linecap: round;
    stroke-width: 2;
  }

  .path-line-14 {
    stroke: greenyellow;
  }

  .path-line-84 {
    stroke: green;
  }

  .path-area {
    fill: rgba(0, 100, 100, 0.2);
  }
</style>
