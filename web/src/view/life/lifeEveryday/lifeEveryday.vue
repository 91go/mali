<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="prefix">
          <el-select
            v-model="searchInfo.prefix"
            placeholder="请选择"
            clearable
            filterable
            @blur="selectBlur"
            @clear="selectClear"
            @change="selectChange"
          >
            <el-option
              v-for="item in prefixOptions"
              :key="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询
          </el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增
        </el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">取消
            </el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <!--        <el-table-column align="left" label="日期" width="180">-->
        <!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="前缀" prop="prefix" min-width="10%" />
        <el-table-column align="left" label="task字段" prop="task" min-width="30%" />
        <el-table-column align="left" label="任务备注" prop="remark" min-width="30%" />
        <el-table-column align="left" label="timeStub" prop="timeStub" min-width="10%" />
        <el-table-column align="left" label="按钮组" min-width="10%">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateLifeEverydayFunc(scope.row)"
            >变更
            </el-button>
            <el-button
              type="text"
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="前缀:" prop="prefix">
          <el-select v-model="formData.prefix" placeholder="请选择">
            <el-option
              v-for="item in prefixOptions"
              :key="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="task字段:">
          <el-input v-model="formData.task" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:">
          <el-select v-model="formData.timeStub" placeholder="请选择">
            <el-option
              v-for="item in timeStubOptions"
              :key="item.key"
              :label="item.value"
              :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="任务备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="more字段:">
          <Md v-model="formData.more" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'LifeEveryday',
  data() {
    return {
      prefixOptions: ['早起', '中午', '晚上'],
      value: '请选择',
    }
  },
  methods: {
    selectBlur(e) {
      // 意见类型
      if (e.target.value !== '') {
        this.value = e.target.value + '(其他)'
        this.$forceUpdate() // 强制更新
      }
    },
    selectClear() {
      this.value = ''
      this.$forceUpdate()
    },
    selectChange(val) {
      this.value = val
      this.$forceUpdate()
    },
  }
}
</script>

<script setup>
import {
  createLifeEveryday,
  deleteLifeEveryday,
  deleteLifeEverydayByIds,
  updateLifeEveryday,
  findLifeEveryday,
  getLifeEverydayList
} from '@/api/lifeEveryday'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { generateTimeStub } from '@/utils/date'
import Md from '@/components/md/md.vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  prefix: '',
  task: '',
  remark: '',
  timeStub: '',
  more: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getLifeEverydayList({
    page: page.value,
    pageSize: pageSize.value, ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteLifeEverydayFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
	multipleSelection.value.map(item => {
	  ids.push(item.ID)
	})
  const res = await deleteLifeEverydayByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLifeEverydayFunc = async(row) => {
  const res = await findLifeEveryday({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.relifeEveryday
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteLifeEverydayFunc = async(row) => {
  const res = await deleteLifeEveryday({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    prefix: '',
    task: '',
    remark: '',
    timeStub: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createLifeEveryday(formData.value)
      break
    case 'update':
      res = await updateLifeEveryday(formData.value)
      break
    default:
      res = await createLifeEveryday(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}

const timeStubOptions = generateTimeStub(15)
</script>

<style>
</style>
