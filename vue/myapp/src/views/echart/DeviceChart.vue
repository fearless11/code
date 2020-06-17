<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
// import resize from './mixins/resize'
import { searchDevice } from '@/api/business'

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
      default: '350px'
    }
  },

  data() {
    return {
      device: {
        Type1: 'andorid',
        Type2: 'ios',
        Type3: 'pc',
        Type4: 'other'
      },
      chart: null,
      list: [],
      dayList: [],
      monthList: [],
      peopleList: []
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
      searchDevice().then(res => {
        var body = res.data
        for (let i = 0; i < body.length; i++) {
          var name = body[i]['device']
          var desc = body[i]['describe']
          var value = body[i]['value']
          if (desc === '日活跃') {
            this.dayList[name] = value
          }
          if (desc === '月活跃') {
            this.monthList[name] = value
          }
          if (desc === '在线人数') {
            this.peopleList[name] = value
          }
        }
        // console.log('yyy', name, desc, value, this.dayList, this.monthList, this.peopleList)
        this.RenderChart()
      })
    },

    RenderChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      window.onresize = this.chart.resize
      this.chart.setOption({
        title: {
          show: false,
          text: '端设备活跃度',
          textStyle: {
            fontWeight: 'normal',
            fontSize: 12,
            x: '10',
            y: '10',
            color: '#DDD'
          }
        },

        tooltip: {
        },

        grid: {
          show: false,
          top: 7
        //   height: '75%'
        //   bottom: 40
        },

        legend: {
          // right: '0',
          // orient: 'vertical',
          bottom: '1',
          itemWidth: 20,
          itemHeight: 10,
          itemStyle: {
            normal: {
              barBorderRadius: 10
            }
          },
          textStyle: {
            color: '#aaa',
            fontSize: 10
          },

          data: ['andorid', 'ios', 'pc', 'other']
        },

        xAxis: {
          splitArea: { show: false }, // 保留网格区域
          axisLine: {
            show: true,
            lineStyle: {
              color: '#008acd'
            }
          },
          axisTick: {
            show: false
          },
          splitLine: {
            show: false
          },
          axisLabel: {
            show: true,
            textStyle: {
              fontSize: 12,
              color: '#008acd'
            }
          }
        },

        yAxis: {
          type: 'category',
          axisLabel: {
            // interval: 0,
            // rotate: 20,
            // Y轴标签垂直显示
            formatter: function(value) {
              return value.split('').join('\n')
            },
            textStyle: {
              fontSize: 12,
              color: '#008acd'
            },
            lineHeight: 15
          },
          axisLine: {
            show: true,
            lineStyle: {
              length: 1,
              color: '#008acd'
            }
          },
          axisTick: {
            show: false
          },
          splitLine: {
            show: false
          },
          data: ['日活跃', '月活跃']
        },

        series: [
          {
            name: 'andorid',
            type: 'bar',
            // barWidth: 15,
            // barCategoryGap: '30%',
            itemStyle: {
              normal: {
                barBorderRadius: 20
              }
            },
            label: {
              show: true,
              position: 'right'
            },
            data: [
              { value: this.dayList[this.device.Type1], name: '日活跃' },
              { value: this.monthList[this.device.Type1], name: '月活跃' }
            ]
          },
          {
            name: this.device.Type4,
            type: 'bar',
            // barWidth: 15,
            // barCategoryGap: '30%',
            itemStyle: {
              normal: {
                barBorderRadius: 20
              }
            },
            label: {
              show: true,
              position: 'right'
            },
            data: [
              { value: this.dayList[this.device.Type4], name: '日活跃' },
              { value: this.monthList[this.device.Type4], name: '月活跃' }
            ]
          },
          {
            name: this.device.Type2,
            type: 'bar',
            // barWidth: 15,
            // barCategoryGap: '30%',
            itemStyle: {
              normal: {
                barBorderRadius: 20
              }
            },
            label: {
              show: true,
              position: 'right'
            },
            data: [
              { value: this.dayList[this.device.Type2], name: '日活跃' },
              { value: this.monthList[this.device.Type2], name: '月活跃' }
            ]
          },
          {
            name: this.device.Type3,
            type: 'bar',
            // barWidth: 15,
            // barCategoryGap: '30%',
            itemStyle: {
              normal: {
                barBorderRadius: 20
              }
            },
            label: {
              show: true,
              position: 'right'
            },
            data: [
              { value: this.dayList[this.device.Type3], name: '日活跃' },
              { value: this.monthList[this.device.Type3], name: '月活跃' }
            ]
          }
        ]
      })
    }
  }
}
</script>
