<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
// vintage dark macarons
require('echarts/theme/walden') // echarts theme
// import resize from './mixins/resize'
import { searchNginxXX } from '@/api/business'

export default {
  // mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '160px'
    }
  },
  data() {
    return {
      list: [],
      chart: null
    }
  },
  mounted() {
    this.initChart()
    setInterval(this.initChart, 600000)
    // this.$nextTick(() => {
    //   this.initChart()
    // })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      searchNginxXX().then(res => {
        this.list = [
          // { 'name': '1xx', 'value': res.data.xx1 },
          { 'name': '2xx', 'value': res.data.xx2 },
          { 'name': '3xx', 'value': res.data.xx3 },
          { 'name': '4xx', 'value': res.data.xx4 },
          { 'name': '5xx', 'value': res.data.xx5 }
        ]

        this.RenderChart()
      })
    },

    RenderChart() {
      this.chart = echarts.init(this.$el, 'walden')
      window.onresize = this.chart.resize
      this.chart.setOption({
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          // right: '0',
          left: '0',
          orient: 'vertical',
          bottom: '0',
          // top: '0',
          itemWidth: 8,
          itemHeight: 5,
          itemGap: 3,
          data: ['1xx', '2xx', '3xx', '4xx', '5xx']
        },
        grid: {
          show: false
          // top: 10
        },
        series: [
          {
            name: '具体情况',
            type: 'pie',
            radius: '80%',
            itemStyle: {
              normal: {
                label: {
                  show: true,
                  formatter: '{d}%',
                  position: 'top'
                },
                labelLine: { // 调节饼图线的长度
                  show: true,
                  length: 1,
                  length2: 25
                }
              }
            },
            // roseType: 'radius',
            // radius: [90, 30],
            center: ['60%', '50%'],
            data: this.list,
            animationEasing: 'cubicInOut',
            animationDuration: 2600
          }
        ]
      })
    }
  }
}
</script>
