<template>
    <div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search v-model="queryParam.username" placeholder="输入用户名查找" enter-button allowClear @search="getUserList" />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible = true">新增</a-button>
                </a-col>
            </a-row>
            <a-table 
                rowKey="username" 
                :columns="columns"
                :pagination="pagination" 
                :dataSource="userlist"
                bordered
                @change="handleTableChange"
            >
                <span slot="role" slot-scope="role">{{role == 1 ? '管理员':'订阅者'}}</span>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" icon="edit" style="margin-right: 15px;" @click="editUser(data.ID)">编辑</a-button>
                        <a-button type="danger" icon="delete" style="margin-right: 15px;" @click="deleteUser(data.ID)">删除</a-button>
                        <a-button type="info" icon="info" @click="deleteUser(data.ID)">重置</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>

        <!-- 新增用户的模块 -->
        <a-modal
            closable
            title="新增用户"
            width="60%"
            :visible="addUserVisible"
            @ok="addUserOk"
            @cancel="addUserCancel"
            destroyOnClose
        >
            <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
                <a-form-model-item has-feedback label="用户名" prop="username">
                    <a-input v-model="newUser.username" ></a-input>
                </a-form-model-item>
                <a-form-model-item has-feedback label="密码" prop="password">
                    <a-input-password v-model="newUser.password" ></a-input-password>
                </a-form-model-item>
                <a-form-model-item has-feedback label="确认密码" prop="checkpasswd">
                    <a-input-password v-model="newUser.checkpasswd" ></a-input-password>
                </a-form-model-item>
                <a-form-model-item label="是否为管理员" prop="role">
                    <a-select defaultValue="2"  @change="adminChange">
                        <a-select-option key="1" value="1">是</a-select-option>
                        <a-select-option key="2" value="2">否</a-select-option>
                    </a-select>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑用户的模块 -->
        <a-modal
        closable
        title="编辑用户"
        width="60%"
        :visible="editUserVisible"
        @ok="editUserOk"
        @cancel="editUserCancel"
    >
        <a-form-model :model="userInfo" :rules="userRules" ref="editUserRef">
            <a-form-model-item has-feedback label="用户名" prop="username">
                <a-input v-model="userInfo.username" ></a-input>
            </a-form-model-item>
            <a-form-model-item label="是否为管理员" prop="role">
                <a-switch checked-children="是" un-checked-children="否" @change="adminChangeSwitch"/>
            </a-form-model-item>
        </a-form-model>
    </a-modal>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        width: '10%',
        key:'id',
        align:'center'
    },
    {
        title: '用户名',
        dataIndex: 'username',
        width: '20%',
        key:'username',
        align:'center'
    },
    {
        title: '角色',
        dataIndex: 'role',
        width: '20%',
        key:'role',
        align:'center',
        scopedSlots: { customRender: 'role' }
    },
    {
        title: '操作',
        width: '30%',
        key:'action',
        align:'center',
        scopedSlots: { customRender: 'action' }
    }
]

export default {
    data () {
        return {
            pagination:{
                pageSizeOptions:['5', '10', '20'],
                defaultCurrent: 1,
                defaultPageSize: 5,
                total:0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
            },
            userlist: [],
            columns,
            queryParam: {
                username:'',
                pagesize: 5,
                pagenum: 1,
            },
            userInfo: {
                id: 0,
                username: '',
                password: '',
                checkpasswd: '',
                role: 2
            },
            newUser: {
                id: 0,
                username: '',
                password: '',
                checkpasswd: '',
                role: 2
            },
            userRules: {
                username: [
                    {
                        required: true,
                        message:'请输入用户名',
                        trigger: 'blur'
                    },
                    {
                        min: 4,
                        max: 12,
                        message:'用户名位数应在4到12位之间',
                        trigger: 'blur'
                    }
                ],
                password: [
                    { 
                        validator: (rule, value, callback) => {
                            if (this.userInfo.password == '') {
                                callback(new Error('请输入密码'))
                            }
                            if ([ ... this.userInfo.password].length < 6 || [ ... this.userInfo.password].length > 20) {
                                callback(new Error('密码应在6到20位之间'))
                            }
                            else{
                                callback()
                            }
                        }, 
                        trigger: 'blur' 
                    }
                ],
                checkpasswd: [
                    { validator: (rule, value, callback) => {
                        if (this.userInfo.checkpasswd == '') {
                            callback(new Error('请验证密码'))
                        }
                        if (this.userInfo.checkpasswd  !== this.userInfo.password ) {
                            callback(new Error('密码不一致，请重新输入'))
                        }
                        else{
                            callback()
                        }
                    }, trigger: 'blur' }
                ]
            },
            addUserRules: {
                username: [
                    {
                        required: true,
                        message:'请输入用户名',
                        trigger: 'blur'
                    },
                    {
                        min: 4,
                        max: 12,
                        message:'用户名位数应在4到12位之间',
                        trigger: 'blur'
                    }
                ],
                password: [
                    { 
                        validator: (rule, value, callback) => {
                            if (this.newUser.password == '') {
                                callback(new Error('请输入密码'))
                            }
                            if ([ ... this.newUser.password].length < 6 || [ ... this.newUser.password].length > 20) {
                                callback(new Error('密码应在6到20位之间'))
                            }
                            else{
                                callback()
                            }
                        }, 
                        trigger: 'blur' 
                    }
                ],
                checkpasswd: [
                    { validator: (rule, value, callback) => {
                        if (this.newUser.checkpasswd == '') {
                            callback(new Error('请验证密码'))
                        }
                        if (this.newUser.checkpasswd  !== this.newUser.password ) {
                            callback(new Error('密码不一致，请重新输入'))
                        }
                        else{
                            callback()
                        }
                    }, trigger: 'blur' }
                ]
            },
            visible: false,
            addUserVisible: false,
            editUserVisible: false
        }
    },
    created () {
        this.getUserList()
    },
    methods: {
        async getUserList () {
            const { data: res } = await this.$http.get('users', {
                params:{
                    username: this.queryParam.username,
                    pagesize: this.pagination.pagesize,
                    pagenum: this.pagination.pagenum,
                },
            })
            // console.log(res)
            if (res.status != 200) return this.$message.error(res.$message)
            this.userlist = res.data
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
            this.getUserList()
        },
        // 删除用户
        deleteUser (ID) {
            this.$confirm({
                title: '提示：二次确认',
                content:'确认删除该用户吗？一旦删除无法恢复',
                 onOk: async () => {
                    // return new Promise((resolve, reject) => {
                    //     setTimeout(Math.random() > 0.5 ? resolve : reject, 1000);
                    // }).catch(() => console.log('Oops errors!'));
                    const res = await this.$http.delete(`user/${ID}`)
                    if (res.status != 200) return this.$message.error(res.$message)
                    this.$message.success('删除成功')
                    this.getUserList()
                },
                onCancel: () => {
                    this.$message.info('已经取消删除')
                }
            })
        },
        // 新增用户
        addUserOk () {
            this.$refs.addUserRef.validate(async (valid) => {
                if (!valid) return this.$message.error('参数不符合要求，请重新输入')
                const { data:res } = await this.$http.post('user/add', {
                    username: this.newUser.username,
                    password: this.newUser.password,
                    role: this.newUser.role
                })
                console.log(res)
                // if (res.status != 200) return this.$message.error(res.message)
                this.addUserVisible = false
                this.$message.success('添加用户成功')
                this.getUserList()
            })
        },
        addUserCancel () {
            this.$refs.addUserRef.resetFields()
            this.addUserVisible = false
            this.$message.info('添加已取消')
        },
        adminChange (value) {
            // 强制转换 字符串 为 数字
            this.userInfo.role = Number(value)
        },

        adminChangeSwitch (checked) {
            if (checked) {
                this.userInfo.role = 1
            }else{
                this.userInfo.role = 2
            }
        },

        // 编辑用户
        async editUser (id) {
            this.editUserVisible = true
            const { data: res } = await this.$http.get(`user/${id}`)
            this.userInfo = res.data
            this.userInfo.id = id
        },
        editUserOk () {
            this.$refs.editUserRef.validate(async (valid) => {
                if (!valid) return this.$message.error('参数不符合要求，请重新输入')
                const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
                    username: this.userInfo.username,
                    role: this.userInfo.role
                })
                if (res.status != 200) return this.$message.error(res.message)
                this.editUserVisible = false
                this.$message.success('更新用户信息成功')
                this.getUserList()
            })
        },
        editUserCancel () {
            this.$refs.editUserRef.resetFields()
            this.editUserVisible = false
            this.$message.info('编辑已取消')
        }
    }
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>
