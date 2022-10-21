<template>
  <Line
      :chart-options="chartOptions"
      :chart-data="chartData"
      :chart-id="chartId"
      :dataset-id-key="datasetIdKey"
      :plugins="plugins"
      :css-classes="cssClasses"
      :styles="styles"
      :width="width"
      :height="height"
  />
</template>

<script>
import { Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement)

export default {
  name: 'LineChart',
  components: { Line },
  props: {
    chartId: {
      type: String,
      default: 'line-chart'
    },
    datasetIdKey: {
      type: String,
      default: 'label'
    },
    width: {
      type: Number,
      default: 400
    },
    height: {
      type: Number,
      default: 400
    },
    cssClasses: {
      default: '',
      type: String
    },
    styles: {
      type: Object,
      default: () => {}
    },
    plugins: {
      type: Object,
      default: () => {}
    },
    chartData: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    console.log(this.chartData)
    return {
      // chartData: {
      //   labels: [ 'January', 'February', 'March' ],
      //   datasets: [ { data: [40, 20, 12] } ]
      // },
      chartOptions: {
        interaction: {
          intersect: false,
          mode: 'index',
        },
        responsive: true,
        scales: {
          x: {
            type: 'linear',
            ticks: {
              stepSize: 10
            }
          }
        }
      },
    }
  }
}
</script>