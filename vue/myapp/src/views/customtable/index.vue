<template>
  <div class="app-container">

    <div class="filter-container">
      <el-input v-model="search" placeholder="团队/成员" class="filter-item" style="width: 400px;" @keyup.enter.native="handleFilter" />
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
        新增
      </el-button>
    </div>

    <el-table
      v-if="list"
      :key="tableKey"
      v-loading="listLoading"
      :data="list.slice((currentPage - 1) * pageSize,currentPage * pageSize)"
      border
      fit
      highlight-current-row
      style="width: 80%"
    >

      <el-table-column label="团队名称" prop="usergroup" align="center" width="200" sortable>
        <template slot-scope="scope">
          <span>{{ scope.row.usergroup }}  </span>
        </template>
      </el-table-column>

      <el-table-column label="团队成员" prop="email" align="center" width="600">
        <template slot-scope="scope">
          {{ scope.row.email | arrToString() }}
        </template>
      </el-table-column>

      <el-table-column label="通知方式" prop="notifiers" align="center" width="150">
        <template slot-scope="scope">
          <div v-for="(way) in scope.row.notifiers" :key="way" style="display: inline-block; width: 15%;">
            <i v-if="way == 'qcloud_voiceprompt_configs'" class="el-icon-phone" />
            <svg-icon v-if="way == 'email_configs'" icon-class="email" />
            <svg-icon v-if="way == 'internal_api_configs'" icon-class="wechat" />
          </div>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="mini"
            @click="handleUpdate(scope.row)"
          >
            编辑
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      :total="total"
      :page.sync="pageSize"
      :current.sync="currentPage"
      :page-sizes="pageSizes"
      @pagination="getList"
    />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width:600px; margin-left:10px;">
        <el-form-item label="团队名称" prop="usergroup">
          <el-input v-model="temp.usergroup" :readonly="dialogStatus=='update'" />
        </el-form-item>
        <el-form-item label="成员">
          <el-select
            v-model="selectedUsers"
            multiple
            filterable
            allow-create
            default-first-option
            style="width:100%"
            placeholder="请选择"
          >
            <el-option
              v-for="item in usersList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="通知方式" prop="condition">
          <template>
            <el-checkbox-group v-model="notifyList">
              <el-checkbox label="email_configs"><svg-icon icon-class="email" /></el-checkbox>
              <el-checkbox label="internal_api_configs"><svg-icon icon-class="wechat" /></el-checkbox>
              <el-checkbox label="qcloud_voiceprompt_configs"><i class="el-icon-phone" /></el-checkbox>
            </el-checkbox-group>
          </template>
        </el-form-item>

      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogStatus === 'create'? createData(): updateData()">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { arrToString } from '@/utils/index'
// 自定义分页组件
import Pagination from '@/components/Pagination'
import { getAllUser, createUsergroup, updateUsergroup, deleteUsergroup } from '@/api/alertm'

export default {
  name: 'LogTable',
  filters: {
    //   自定义过滤函数
    arrToString
  },
  components: { Pagination },
  data() {
    return {
      tableKey: 0,
      // checkbox
      notifyList: [],
      // select
      usersList: null,
      selectedUsers: [],
      // dialog
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      temp: {
        usergroup: '',
        notifiers: [],
        email: []
      },
      // pagination
      total: 0,
      currentPage: 1,
      // bug: pageSize和pageSizes的第一个值设置为一样
      pageSize: 2,
      pageSizes: [2, 20, 30],
      // search
      search: '',
      // raw list data
      rawList: null,
      // table list data
      list: null,
      listLoading: true
    }
  },
  created() {
    this.getList()
    this.getUsersList()
  },
  methods: {
    getUsersList() {
      getAllUser()
        .then(response => {
          var users = []
          for (var i in response.data.data) {
            users.push(response.data.data[i].email)
          }
          this.usersList = users.map(item => {
            return { value: `${item}`, label: `${item}` }
          })
        })
        .catch(error => { console.log(error) })
    },
    getList() {
      this.listLoading = true

      // 模拟后端返回数据格式
      this.rawList = [
        {
          'usergroup': 'default-weixin-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        },
        {
          'usergroup': 'canal-group',
          'notifiers': [
            'internal_api_configs'
          ],
          'email': [
            'aaa@xxx.com',
            'bbb@xxx.com'
          ]
        }
      ]
      this.list = this.rawList
      this.total = this.list.length
      this.listLoading = false

    //   getAllUsergroup()
    //     .then(response => {
    //       this.rawList = response.data.data
    //       this.list = this.rawList
    //       this.total = this.list.length
    //     })
    //     .catch(error => { console.log(error) })
    //     .finally(() => { this.listLoading = false })
    },

    handleFilter() {
      if (this.search.length === 0) {
        this.list = this.rawList
      } else {
        this.list = this.rawList.filter(data => !this.search || data.usergroup.toLowerCase().includes(this.search.toLowerCase()))
        if (this.list.length === 0) {
          // 多个邮箱搜索
          for (var i in this.rawList) {
            if (this.rawList[i].email.join(',').indexOf(this.search.toLowerCase()) !== -1) {
              this.list.push(this.rawList[i])
            }
          }
        }
      }
      this.total = this.list.length
    },

    resetTemp() {
      this.temp = {
        usergroup: '',
        email: [],
        notifiers: []
      }
    },
    handleCreate() {
      this.resetTemp()
      this.selectedUsers = []
      this.notifyList = []
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.dialogFormVisible = false
          this.temp.email = this.selectedUsers
          this.temp.notifiers = this.notifyList
          createUsergroup(this.temp)
            .then(res => {
              if (res.data.code === 200) {
                this.$notify({
                  title: res.data.status,
                  type: 'success',
                  duration: 2000
                })
              } else {
                this.$notify({
                  title: res.data.error,
                  type: 'error',
                  duration: 2000 })
              }
              this.getList()
            })
            .catch(error => { console.log(error) })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row)
      this.selectedUsers = this.temp.email
      this.notifyList = this.temp.notifiers
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.dialogFormVisible = false
          this.temp.notifiers = this.notifyList
          this.temp.email = this.selectedUsers
          updateUsergroup(this.temp.usergroup, this.temp)
            .then(res => {
              if (res.data.code === 200) {
                this.$notify({
                  title: res.data.status,
                  type: 'success',
                  duration: 2000
                })
              } else {
                this.$notify({
                  title: res.data.error,
                  type: 'error',
                  duration: 2000 })
              }
              this.getList()
            })
            .catch(error => { console.log(error) })
        }
      })
    },
    handleDelete(index, row) {
      this.temp = Object.assign({}, row)
      deleteUsergroup(this.temp.usergroup)
        .then(res => {
          if (res.data.code === 200) {
            this.$notify({
              title: res.data.status,
              type: 'success',
              duration: 2000
            })
            this.list.splice(index, 1)
            this.getList()
          } else {
            this.$notify({
              title: res.data.error,
              type: 'error',
              duration: 2000 })
          }
        })
        .catch(error => { console.log(error) })
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
