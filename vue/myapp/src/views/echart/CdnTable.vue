<template>
  <el-table
    :data="list"
    :style="{background:'#062238'}"
    :row-style="{background:'#062238', height:'12px'}"
    :header-cell-style="{background:'#062238',color:'#aaa'}"
  >
    <el-table-column label="域名" width="130" min-width="170">
      <template slot-scope="scope">
        <span class="table_txt"> {{ scope.row.domain | domainFilter }} </span>
      </template>
    </el-table-column>
    <el-table-column label="流量(M)" width="70" align="left">
      <template slot-scope="scope">
        <span class="table_txt"> {{ scope.row.flux | fluxFilter }}</span>
      </template>
    </el-table-column>
    <el-table-column label="访问量" width="90" align="left">
      <template slot-scope="{row}">
        <span class="table_txt"> {{ row.requests }}</span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { searchCDN } from '@/api/business'

export default {
  filters: {
    domainFilter(str) {
      return str ? str.substring(0, 14) : ''
    },
    fluxFilter(num) {
      var MBytes = (num / 1024 / 1024).toFixed(0)
      return MBytes
    }
  },

  data() {
    return {
      list: []
    }
  },

  created() {
    this.fetchData()
    setInterval(this.fetchData, 86400000)
  },

  methods: {
    fetchData() {
      // searchCDN().then(response => {
      //   this.list = response.data.slice(0, 10)
      // })
      // 模拟后端返回
      this.list = [{ 'domain': 'photo.zastatic.com', 'flux': 35418421928, 'requests': 377512798 }, { 'domain': 'quyuehui-1251661065.image.myqcloud.com', 'flux': 1476989064, 'requests': 44800193 }, { 'domain': 'file.zastatic.com', 'flux': 1507367040, 'requests': 31376499 }, { 'domain': 'zhenai4saylove-1251661065.file.myqcloud.com', 'flux': 722806880, 'requests': 22570920 }, { 'domain': 'i.zhenai.com', 'flux': 303044336, 'requests': 18862834 }, { 'domain': 'images.zastatic.com', 'flux': 7623930080, 'requests': 11544275 }, { 'domain': 'zhenai4video-1251661065.image.myqcloud.com', 'flux': 235114784, 'requests': 6885040 }, { 'domain': 'shortvideo.zhenai.com', 'flux': 10504445736, 'requests': 3760503 }, { 'domain': 'static.jiebao.zhenai.com', 'flux': 883139664, 'requests': 3739191 }, { 'domain': 'zhenai4youth-1251661065.image.myqcloud.com', 'flux': 129552776, 'requests': 1257723 }, { 'domain': 'm.quyuehui.com', 'flux': 5989192, 'requests': 854855 }, { 'domain': 'photo12.zastatic.com', 'flux': 41552280, 'requests': 574170 }, { 'domain': 'photo11.zastatic.com', 'flux': 42470504, 'requests': 556400 }, { 'domain': 'photo13.zastatic.com', 'flux': 40381416, 'requests': 539376 }, { 'domain': 'photo16.zastatic.com', 'flux': 40912768, 'requests': 535578 }, { 'domain': 'photo14.zastatic.com', 'flux': 41610640, 'requests': 530120 }, { 'domain': 'photo15.zastatic.com', 'flux': 40342560, 'requests': 522400 }, { 'domain': 'quyuehui-1251661065.file.myqcloud.com', 'flux': 36310704, 'requests': 190220 }, { 'domain': 'm.virover.com', 'flux': 723104, 'requests': 58037 }, { 'domain': 'images.virover.com', 'flux': 860048, 'requests': 36062 }, { 'domain': 'bidataimages.zhenai.com', 'flux': 5020480, 'requests': 33289 }, { 'domain': 'zhenaicrm-1251661065.file.myqcloud.com', 'flux': 10178720, 'requests': 27304 }, { 'domain': 'photo1.zastatic.com', 'flux': 6024280, 'requests': 16783 }, { 'domain': 'pubser-res.zhenai.com', 'flux': 2011792, 'requests': 15173 }, { 'domain': 'photo3.zastatic.com', 'flux': 5971344, 'requests': 14009 }, { 'domain': 'photo2.zastatic.com', 'flux': 5751736, 'requests': 12447 }, { 'domain': 'imagescrm.zhenai.com', 'flux': 82128, 'requests': 12343 }, { 'domain': 'photo4.zastatic.com', 'flux': 6389624, 'requests': 11500 }, { 'domain': 'www.virover.com', 'flux': 58344, 'requests': 11284 }, { 'domain': 'photo0.zastatic.com', 'flux': 6093440, 'requests': 10943 }, { 'domain': 'zhenai4saylove-1251661065.image.myqcloud.com', 'flux': 1731304, 'requests': 8318 }, { 'domain': 'www.quyuehui.com', 'flux': 137696, 'requests': 4996 }, { 'domain': 'www.lecreater.com', 'flux': 5352, 'requests': 1524 }, { 'domain': 'cdn.imjieli.com', 'flux': 1984, 'requests': 1444 }, { 'domain': 'zaw-10009734.file.myqcloud.com', 'flux': 12767624, 'requests': 1341 }, { 'domain': 'g.virover.com', 'flux': 90464, 'requests': 945 }, { 'domain': 'photo.zhenai.com', 'flux': 139160, 'requests': 620 }, { 'domain': 'photo1.zhenai.com', 'flux': 15496, 'requests': 86 }, { 'domain': 'zaw-10009734.image.myqcloud.com', 'flux': 12767624, 'requests': 61 }, { 'domain': 'photo4.zhenai.com', 'flux': 336, 'requests': 51 }, { 'domain': 'cdn.jiebao.zhenai.com', 'flux': 1984, 'requests': 48 }, { 'domain': 'photo16.zhenai.com', 'flux': 40912768, 'requests': 44 }, { 'domain': 'photo15.zhenai.com', 'flux': 40342560, 'requests': 40 }, { 'domain': 'cos.zastatic.com', 'flux': 0, 'requests': 39 }, { 'domain': 'photo3.zhenai.com', 'flux': 5971344, 'requests': 38 }, { 'domain': 'photo0.zhenai.com', 'flux': 0, 'requests': 37 }, { 'domain': 'photo2.zhenai.com', 'flux': 5751736, 'requests': 33 }, { 'domain': 'photo14.zhenai.com', 'flux': 41610640, 'requests': 31 }, { 'domain': 'photo11.zhenai.com', 'flux': 42470504, 'requests': 26 }, { 'domain': 'photo12.zhenai.com', 'flux': 41552280, 'requests': 24 }, { 'domain': 'photo13.zhenai.com', 'flux': 40381416, 'requests': 22 }, { 'domain': 'cdn.wx.zhenai.com', 'flux': 1984, 'requests': 18 }, { 'domain': 'next.zhenai.com', 'flux': 723104, 'requests': 17 }, { 'domain': 'test.images.zhenai.com', 'flux': 883139664, 'requests': 15 }, { 'domain': 'static.sweet.zhenaioa.com', 'flux': 883139664, 'requests': 10 }, { 'domain': 'cdn.vxq.zhenai.com', 'flux': 1984, 'requests': 8 }, { 'domain': 'picture1-10009734.image.myqcloud.com', 'flux': 336, 'requests': 8 }, { 'domain': 'zhenai4zone-1251661065.file.myqcloud.com', 'flux': 129552776, 'requests': 4 }, { 'domain': 'wxmm-10009734.file.myqcloud.com', 'flux': 58344, 'requests': 4 }, { 'domain': 'wxmm1-10009734.image.myqcloud.com', 'flux': 58344, 'requests': 4 }, { 'domain': 'zhenai4divorce-1251661065.image.myqcloud.com', 'flux': 12767624, 'requests': 2 }, { 'domain': 'm.kaojin.com.cn', 'flux': 82128, 'requests': 1 }, { 'domain': 'realphoto2.zhenai.com', 'flux': 1476989064, 'requests': 1 }, { 'domain': 'i.kaojin.com.cn', 'flux': 90464, 'requests': 0 }, { 'domain': 'realphoto3.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'static-apioa.zhenai.com', 'flux': 10504445736, 'requests': 0 }, { 'domain': 'realphoto16.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'vxq-10009734.file.myqcloud.com', 'flux': 883139664, 'requests': 0 }, { 'domain': 'www.bluewhaledt.com', 'flux': 883139664, 'requests': 0 }, { 'domain': 'crm01.zhenai.com', 'flux': 0, 'requests': 0 }, { 'domain': 'www.sayloveapp.com', 'flux': 137696, 'requests': 0 }, { 'domain': 'wxapp-10068901.file.myqcloud.com', 'flux': 58344, 'requests': 0 }, { 'domain': 'realphoto13.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'realphoto.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'dwx.zhenai.com', 'flux': 0, 'requests': 0 }, { 'domain': 'zhenaicrm-1251661065.image.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'zhenairecord-10009734.file.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'i.ieagiqfxis.lecreater.com', 'flux': 90464, 'requests': 0 }, { 'domain': 'zhenaicrm-10075562.image.myqcloud.com', 'flux': 129552776, 'requests': 0 }, { 'domain': 'vip.zhenai.com', 'flux': 883139664, 'requests': 0 }, { 'domain': 'zhenaicrmyt-1251661065.image.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'zhenaicrmpic-1251661065.image.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'images.zhenai.com', 'flux': 7623930080, 'requests': 0 }, { 'domain': 'zhenai4twelvemeet-1251661065.file.myqcloud.com', 'flux': 1731304, 'requests': 0 }, { 'domain': 'youqun.lecreater.com', 'flux': 0, 'requests': 0 }, { 'domain': 'decoderesult-10009734.file.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'piccheck-10075117.file.myqcloud.com', 'flux': 336, 'requests': 0 }, { 'domain': 'cdn.sayloveapp.com', 'flux': 1984, 'requests': 0 }, { 'domain': 'i.yuanlai.com', 'flux': 90464, 'requests': 0 }, { 'domain': 'realphoto15.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'divorce-10009734.image.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'zhenairecord2-10009734.file.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'zhenai4live-1251661065.file.myqcloud.com', 'flux': 12767624, 'requests': 0 }, { 'domain': 'zbsh.zhenai.com', 'flux': 12767624, 'requests': 0 }, { 'domain': 'm.sayloveapp.com', 'flux': 5989192, 'requests': 0 }, { 'domain': 'realphoto14.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'realphoto11.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'zhenai4twelvemeet-1251661065.image.myqcloud.com', 'flux': 1731304, 'requests': 0 }, { 'domain': 'zhenai-oa-1251661065.file.myqcloud.com', 'flux': 12767624, 'requests': 0 }, { 'domain': 'realphoto0.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'faceid-10074391.image.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'zhenai4zone-1251661065.image.myqcloud.com', 'flux': 129552776, 'requests': 0 }, { 'domain': 'zhenaioa-1251661065.image.myqcloud.com', 'flux': 10178720, 'requests': 0 }, { 'domain': 'jiebao01.zhenai.com', 'flux': 82128, 'requests': 0 }, { 'domain': 'pubser-1251661065.file.myqcloud.com', 'flux': 336, 'requests': 0 }, { 'domain': 'zhenai4video-1251661065.file.myqcloud.com', 'flux': 1731304, 'requests': 0 }, { 'domain': 'bidata-1251661065.image.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'oa-common-1251661065.file.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'jiebao.zhenai.com', 'flux': 82128, 'requests': 0 }, { 'domain': 'oa-inventory-1251661065.file.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'realphoto12.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'oa-inventory-1251661065.image.myqcloud.com', 'flux': 0, 'requests': 0 }, { 'domain': 'realphoto4.zhenai.com', 'flux': 1476989064, 'requests': 0 }, { 'domain': 'realphoto1.zhenai.com', 'flux': 1476989064, 'requests': 0 }]
    }
  }
}
</script>

<style scoped>
.table_txt {
  color: #3c7c97;
}
</style>
