<template>
  <div class="app-container">
    <div id="heatmapchart" :class="className" :style="{height:height,width:width}" />
  </div>
</template>

<script>
// echart教程 https://echarts.apache.org/zh/tutorial.html#5%20%E5%88%86%E9%92%9F%E4%B8%8A%E6%89%8B%20ECharts
// 引入echarts组件
import echarts from 'echarts'
import 'echarts/map/js/china.js'
// 单独引入
// require('echarts/theme/macarons')
// import resize from './mixins/resize'
// import { searchCeSu } from '@/api/business'

export default {
//   mixins: [resize],
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
      chart: null,
      mapJson: []
    }
  },
  mounted() {
    // mounted生命周期函数中实例化echarts对象。因为我们要确保dom元素已经挂载到页面中
    this.initMapJson()
  },

  created() {
    // this.initMapJson()
    setInterval(this.initMapJson, 21600000)
  },

  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },

  methods: {
    // 初始化地图数据
    initMapJson() {
      // searchCeSu().then(response => {
      //   this.mapJson = response.data
      //   this.renderChinaMap()
      // })
      this.mapJson = [{ 'name': '辽宁', 'value': 0.183 }, { 'name': '湖南', 'value': 0.232 }, { 'name': '广西', 'value': 0.229 }, { 'name': '安徽', 'value': 0.13 }, { 'name': '黑龙江', 'value': 0.183 }, { 'name': '浙江', 'value': 0.148 }, { 'name': '江西', 'value': 0.112 }, { 'name': '重庆', 'value': 0.131 }, { 'name': '福建', 'value': 0.134 }, { 'name': '江苏', 'value': 0.229 }, { 'name': '湖北', 'value': 0.179 }, { 'name': '上海', 'value': 0.081 }, { 'name': '山东', 'value': 0.116 }, { 'name': '甘肃', 'value': 0.23 }, { 'name': '北京', 'value': 0.033 }, { 'name': '四川', 'value': 0.15 }, { 'name': '广东', 'value': 0.176 }, { 'name': '陕西', 'value': 0.18 }, { 'name': '美国', 'value': 0.967 }, { 'name': '河南', 'value': 0.435 }, { 'name': '内蒙古', 'value': 0.089 }, { 'name': '加拿大', 'value': 0.473 }, { 'name': '香港', 'value': 0.044 }, { 'name': '山西', 'value': 0.127 }, { 'name': '新疆', 'value': 0.387 }, { 'name': '青海', 'value': 0.495 }, { 'name': '云南', 'value': 0.181 }, { 'name': '新加坡', 'value': 0.338 }]
      this.renderChinaMap()
    },

    renderChinaMap() {
      // this.$el是<template>标签里的所有内容
      // document.getElementById('mychart')只是到id为heatmapchart的<div>
      // console.log('this.$el', this.$el, document.getElementById('mychart'))
      this.chart = echarts.init(document.getElementById('heatmapchart'), 'macarons')
      window.onresize = this.chart.resize
      this.chart.setOption({
        tooltip: {},
        grid: {
          show: false,
          top: 2,
          right: 10,
          left: 10,
          bottom: 20
        },
        // 地图坐标系组件
        geo: {
          map: 'china',
          roam: false,
          layoutCenter: ['50%', '50%'],
          layoutSize: 450,
          // 地图上文本标签
          label: {
            normal: {
              show: true, // 是否显示对应地名及颜色设置
              textStyle: {
                color: '#000000' // 地图轮廓线黑色
              }
            }
          },
          // 地图区域多边形、图形样式
          itemStyle: {
            borderColor: '#cccccc', // 区域默认背景 深灰色
            areaColor: '#b0b5bf',
            emphasis: {
              areaColor: null, // 点击时显示区域颜色
              shadowOffsetX: 0,
              shadowOffsetY: 0,
              shadowBlur: 20,
              borderWidth: 0,
              shadowColor: '#69b7b5'
            }
          }
        },
        // 系列列表设置,每个系列通过type觉得图标类型
        series: [{
          name: '响应时间', // 浮动框的标题
          type: 'map',
          geoIndex: 0,
          data: this.mapJson
        }],
        //  数据映射组件
        visualMap: [{
          'show': true,
          textStyle: {
            color: '#aaa'
          },
          'type': 'piecewise',
          'left': '5%',
          'bottom': '5%',
          'pieces': [
            { 'gte': 0, 'lt': 1, label: '好(0~1s)', 'color': '#008000' },
            { 'gte': 1, 'lt': 2, label: '较好(1~2s)', 'color': '#00ff40' },
            { 'gte': 2, 'lt': 3, label: '警告(2~3s)', 'color': '#ecec00' },
            { 'gte': 3, 'lt': 5, label: '较差(3~5s)', 'color': '#ff6600' },
            { 'gte': 5, label: '(差>5s)', 'color': '#d50000' }
          ]
        }]
      })
    }
  }
}
</script>
