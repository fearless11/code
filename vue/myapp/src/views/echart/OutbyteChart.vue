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
      default: '160px'
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
        var arr = []
        for (let i = 0; i < this.list.length; i++) {
          if (this.list[i].outBytes === 0) {
            continue
          }
          var domain = this.list[i].domain.substring(0, 3)
          var outBytes = this.list[i].outBytes
          arr.push({ 'name': domain, 'outBytes': outBytes })
        }

        var compare = function(obj1, obj2) {
          var val1 = obj1.outBytes
          var val2 = obj2.outBytes
          if (val1 < val2) {
            return 1
          } else if (val1 > val2) {
            return -1
          } else {
            return 0
          }
        }

        var arr1 = arr.sort(compare)
        for (let i = 0; i < arr1.length; i++) {
          this.domainList.push(arr1[i].name)
          var kbyte = arr1[i].outBytes / 1024
          this.outByteList.push(kbyte.toFixed(0))
        }

        console.log(this.outByteList)

        this.renderChart()
      })
    },

    renderChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      window.onresize = this.chart.resize
      this.chart.setOption({
        tooltip: {
          trigger: 'axis',
          axisPointer: { // 坐标轴指示器，坐标轴触发有效
            type: 'shadow' // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        grid: {
          top: 10,
          left: '2%',
          right: '2%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: [{
          type: 'category',
          splitLine: { show: false }, // 去除网格线
          data: this.domainList.slice(0, 5),
          axisTick: {
            alignWithLabel: false
          },
          axisLabel: {
            show: true,
            textStyle: {
              // color: '#aaa',
              fontSize: 10
            }
          },
          splitArea: { show: false }// 保留网格区域
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
          right: '0',
          orient: 'vertical',
          icon: 'none',
          // bottom: '10',
          textStyle: {
            color: '#aaa',
            fontSize: 10
          },
          itemStyle: {
            normal: {
              barBorderRadius: 10
            }
          },
          data: ['出口流量(KB)']
        },

        series: [{
          name: '出口流量(KB)',
          type: 'bar',
          // stack: 'vistors',
          // barWidth: '60%',
          label: {
            show: true,
            position: 'top',
            fontSize: 9
          },
          data: this.outByteList.slice(0, 5)

        }]
      })
    }
  }
}
</script>
