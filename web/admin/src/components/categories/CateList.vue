<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-input-search
            v-model="queryParam.name"
            placeholder="输入分类名查找"
            enter-button
            allowClear
            @search="getCateList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addCateVisible = true"
            >新增</a-button
          >
        </a-col>
      </a-row>
      <a-table
        rowKey="name"
        :columns="columns"
        :pagination="pagination"
        :dataSource="catelist"
        bordered
        @change="handleTableChange"
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editCate(data.id)"
              >编辑</a-button
            >
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteCate(data.id)"
              >删除</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增分类的模块 -->
    <a-modal
      closable
      title="新增分类"
      width="60%"
      :visible="addCateVisible"
      @ok="addCateOk"
      @cancel="addCateCancel"
      destroyOnClose
    >
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item has-feedback label="分类名" prop="name">
          <a-input v-model="newCate.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑分类的模块 -->
    <a-modal
      closable
      title="编辑分类"
      width="60%"
      :visible="editCateVisible"
      @ok="editCateOk"
      @cancel="editCateCancel"
    >
      <a-form-model :model="cateInfo" :rules="cateRules" ref="editCateRef">
        <a-form-model-item has-feedback label="分类名" prop="name">
          <a-input v-model="cateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '5%',
    key: 'id',
    align: 'center',
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '25%',
    key: 'name',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      catelist: [],
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
      cateInfo: {
        id: 0,
        name: '',
      },
      newCate: {
        id: 0,
        name: '',
      },
      cateRules: {
        name: [
          {
            required: true,
            message: '请输入分类名',
            trigger: 'blur',
          },
          {
            min: 2,
            max: 12,
            message: '分类名位数应在2到12位之间',
            trigger: 'blur',
          },
        ],
      },
      addCateRules: {
        name: [
          {
            required: true,
            message: '请输入分类名',
            trigger: 'blur',
          },
          {
            min: 2,
            max: 12,
            message: '分类名位数应在2到12位之间',
            trigger: 'blur',
          },
        ],
      },
      visible: false,
      addCateVisible: false,
      editCateVisible: false,
    }
  },
  created() {
    this.getCateList()
  },
  methods: {
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category', {
        params: {
          name: this.queryParam.name,
          pagesize: this.pagination.pagesize,
          pagenum: this.pagination.pagenum,
        },
      })
      // console.log(res)
      if (res.status != 200) return this.$message.error(res.$message)
      this.catelist = res.data
      this.pagination.total = res.total
    },
    // 更改分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pageSize = pagination.pageSize
      this.queryParam.pagenum = pagination.current
      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getCateList()
    },
    // 删除分类
    deleteCate(ID) {
      this.$confirm({
        title: '提示：二次确认',
        content: '确认删除该分类吗？一旦删除无法恢复',
        onOk: async () => {
          // return new Promise((resolve, reject) => {
          //     setTimeout(Math.random() > 0.5 ? resolve : reject, 1000);
          // }).catch(() => console.log('Oops errors!'));
          const res = await this.$http.delete(`category/${ID}`)
          if (res.status != 200) return this.$message.error(res.$message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('已经取消删除')
        },
      })
    },
    // 新增分类
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name,
        })
        console.log(res)
        // if (res.status != 200) return this.$message.error(res.message)
        this.addCateVisible = false
        this.$message.success('添加分类成功')
        this.getCateList()
      })
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('添加已取消')
    },
    adminChange(value) {
      // 强制转换 字符串 为 数字
      this.cateInfo.role = Number(value)
    },

    adminChangeSwitch(checked) {
      if (checked) {
        this.cateInfo.role = 1
      } else {
        this.cateInfo.role = 2
      }
    },

    // 编辑分类
    async editCate(id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.cateInfo = res.data
      this.cateInfo.id = id
    },
    editCateOk() {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(
          `category/${this.cateInfo.id}`,
          {
            name: this.cateInfo.name,
            role: this.cateInfo.role,
          },
        )
        if (res.status != 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.success('更新分类信息成功')
        this.getCateList()
      })
    },
    editCateCancel() {
      this.$refs.editCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('编辑已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
