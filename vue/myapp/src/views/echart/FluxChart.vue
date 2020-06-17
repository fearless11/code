<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
// import resize from './mixins/resize'
import { searchNginxDomain } from '@/api/business'

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
      default: '200px'
    }
  },
  data() {
    return {
      chart: null,
      list: [],
      domainList: [],
      requestList: [],
      inByteList: [],
      outByteList: []
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
      searchNginxDomain().then(res => {
        this.list = res.data

        for (let i = 0; i < this.list.length; i++) {
          this.domainList.push(this.list[i].domain.substring(0, 5))
          this.requestList.push(this.list[i].request)
          this.inByteList.push(this.list[i].inBytes)
          this.inByteList.push(this.list[i].outBytes)
        }

        this.renderChart()
      })
    },

    renderChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      window.onresize = this.chart.resize
      this.chart.setOption({

        // tooltip: {
        //   trigger: 'axis',
        //   axisPointer: { // 坐标轴指示器，坐标轴触发有效
        //     type: 'line' // 默认为直线，可选为：'line' | 'shadow'
        //   }
        // },

        grid: {
          top: 10,
          left: '2%',
          right: '2%',
          bottom: '3%',
          containLabel: true,
          show: false
        },

        xAxis: [{
          type: 'category',
          data: this.domainList.slice(0, 8),
          axisTick: {
            alignWithLabel: true
          }
        }],
        yAxis: [{
          splitLine: { show: false },
          splitArea: { show: false }, // 保留网格区域
          type: 'value',
          axisTick: {
            show: false
          }
        }],
        legend: {
          show: true,
          right: '1',
          orient: 'vertical',
          // bottom: '10',
          textStyle: {
            color: '#aaa'
          },
          itemStyle: {
            normal: {
              barBorderRadius: 10
            }
          },
          data: ['入口流量', '出口流量']
        },

        series: [{
          name: '入口流量',
          type: 'bar',
          // stack: 'vistors',
          // barWidth: '60%',
          label: {
            show: true,
            position: 'top'
          },
          data: this.inByteList.slice(0, 8)
        }, {
          name: '出口流量',
          type: 'bar',
          label: {
            show: true,
            position: 'top'
          },
          // stack: 'vistors',
          // barWidth: '60%',
          data: this.outByteList.slice(0, 8)
        }]
      })
    }
  }
}
</script>
