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
          if (this.list[i].inBytes === 0) {
            continue
          }
          var domain = this.list[i].domain.substring(0, 3)
          var inBytes = this.list[i].inBytes
          arr.push({ 'name': domain, 'inBytes': inBytes })
        }

        var compare = function(obj1, obj2) {
          var val1 = obj1.inBytes
          var val2 = obj2.inBytes
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
          var kbyte = arr1[i].inBytes / 1024
          this.inByteList.push(kbyte.toFixed(0))
        }

        console.log(this.inByteList)

        this.renderChart()
      })
    },

    renderChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      window.onresize = this.chart.resize
      this.chart.setOption({
        tooltip: {
          trigger: 'axis',
          axisPointer: { // ??????????????????????????????????????????
            type: 'shadow' // ??????????????????????????????'line' | 'shadow'
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
          splitLine: { show: false }, // ???????????????
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
          splitArea: { show: false }// ??????????????????
        }],
        yAxis: [{
          splitLine: { show: false },
          splitArea: { show: false }, // ??????????????????
          type: 'value',
          axisTick: {
            show: false
          }
        }],
        legend: {
          show: true,
          right: '0',
          orient: 'vertical',
          // bottom: '10',
          icon: 'none',
          textStyle: {
            color: '#aaa',
            fontSize: 10
          },
          itemStyle: {
            normal: {
              barBorderRadius: 10
            }
          },
          data: ['????????????(KB)']
        },

        series: [{
          name: '????????????(KB)',
          type: 'bar',
          // stack: 'vistors',
          // barWidth: '60%',
          label: {
            show: true,
            position: 'top',
            fontSize: 10
          },
          data: this.inByteList.slice(0, 5)

        }]
      })
    }
  }
}
</script>
