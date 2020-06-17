<template>
  <div class="pagination-container">
    <el-pagination
      :current-page.sync="currentPage"
      :page-size.sync="pageSize"
      :page-sizes="pageSizes"
      layout="total,sizes,prev,pager,next,jumper"
      :total="total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script>

// https://element.eleme.cn/#/zh-CN/component/pagination
export default {
  name: 'Pagination',
  // props 接收来自父组件的数据
  props: {
    total: {
      required: true,
      type: Number
    },
    current: {
      type: Number,
      default: 1
    },
    page: {
      type: Number,
      default: 2
    },
    pageSizes: {
      type: Array,
      default() {
        return [10, 20, 30, 50]
      }
    }
  },
  computed: {
    // 计算属性currentPage的getter函数和setter函数
    currentPage: {
      get() {
        return this.current
      },
      set(val) {
        // 更新父组件的数据,配合父组件中的:current.sync使用
        this.$emit('update:current', val)
      }
    },
    pageSize: {
      get() {
        return this.page
      },
      set(val) {
        this.$emit('update:page', val)
      }
    }
  },
  methods: {
    handleSizeChange(val) {
      // 两种方式一样的效果
      // 直接更新父组件数据
      this.$emit('update:page', val)
      // 更新子组件数据，触发计算属性中的setter函数
    //   this.$emit('pagination', { current: this.currentPage, page: val })
    },
    handleCurrentChange(val) {
      this.$emit('update:current', val)
    //   this.$emit('pagination', { current: val, page: this.pageSize })
    }
  }

}
</script>

<style scoped>
.pagination-container {
  background: #fff;
  padding: 32px 16px;
}
</style>
