<script>

import {GetChart} from "../../wailsjs/go/main/App.js";
import LineChart from "./LineChart.vue";
import BarChart from "./BarChart.vue";
import Counter from "./Counter.vue";

export default {
  components: {Counter, BarChart, LineChart},
  methods: {
    async getChart() {
      const result = await GetChart()
      const data = {
          labels: [],
          datasets: [
            {data: [], fill: false, label: 'EngineSpeed', borderColor: 'red', lineTension: 0.4},
            {data: [], fill: false, label: 'PressureActual', borderColor: 'orange', lineTension: 0.4},
            {data: [], fill: false, label: 'PressureWanted', borderColor: 'green', lineTension: 0.4},
          ]
        };

        result.forEach((item) => {
          data.labels.push(item.Time)
          // data.datasets[0].data.push({x: item.Time, y: item.EngineSpeed})
          // data.datasets[1].data.push({x: item.Time, y: item.PressureActual})
          // data.datasets[2].data.push({x: item.Time, y: item.PressureWanted})
          data.datasets[0].data.push({x: item.Time, y: item.EngineSpeed})
          data.datasets[1].data.push({x: item.Time, y: item.PressureActual})
          data.datasets[2].data.push({x: item.Time, y: item.PressureWanted})
        })

      return data

    }
  },
  async mounted() {
    this.loaded = false
    this.chart = await this.getChart()
    this.loaded = true

    function getRandomInt(min, max) {
      min = Math.ceil(min);
      max = Math.floor(max);
      return Math.floor(Math.random() * (max - min) + min); // The maximum is exclusive and the minimum is inclusive
    }

    setInterval(() => {
      this.engineSpeed = getRandomInt(1000, 5000)
    }, 100)

    setInterval(() => {
      this.pressureActual = getRandomInt(1000, 5000)
    }, 100)

    setInterval(() => {
      this.pressureWanted = getRandomInt(1000, 5000)
    }, 100)

  },
  data() {
    return {
      chart: null,
      engineSpeed: 0,
      pressureActual: 0,
      pressureWanted: 0,
      loaded: false
    }
  }
}

</script>

<template>
  <div class="counters">
    <Counter v-if="engineSpeed" title="Engine speed" :value="engineSpeed" />
    <Counter v-if="pressureActual" title="Pressure actual" :value="pressureActual" />
    <Counter v-if="pressureWanted" title="Pressure wanted" :value="pressureWanted" />
  </div>
<LineChart v-if="loaded" :chart-data="chart"/>
</template>