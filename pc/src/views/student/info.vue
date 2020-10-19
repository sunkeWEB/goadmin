<template>
  <BasicLayout>
    <template #wrapper>
      <el-card>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysuser:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysdicttype:export']"
              type="warning"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >导出</el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="studentData"
          border
        >
          <el-table-column prop="student_id" label="学号" width="120" align="center"></el-table-column>
          <el-table-column prop="name" label="姓名" width="140" align="center"></el-table-column>
          <el-table-column prop="avatar" label="照片" width="100" align="center">
            <template slot-scope="scope">
              <div @click="openAvatarFn(scope.row)" style="cursor: pointer" title="点击查看照片">
                <el-avatar :src="scope.row.avatar"></el-avatar>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="sex" label="性别" width="100" align="center">
            <template slot-scope="scope">
              <el-tag effect="plain" size="medium">{{scope.row.sex==0?"男":"女"}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="phone" label="电话" align="center"></el-table-column>
          <el-table-column label="年级班级" align="center">
            <template slot-scope="scope">
              {{scope.row.grade}}年级{{scope.row.class_by}}班
            </template>
          </el-table-column>
          <el-table-column label="状态" width="100" align="center">
            <template slot-scope="scope">
              {{scope.row.status===1?"在校":"离校"}}
            </template>
          </el-table-column>
          <el-table-column label="入学时间" align="center">
            <template slot-scope="scope">
              {{scope.row.createdAt}}
            </template>
          </el-table-column>
          <el-table-column prop="contact_by" label="紧急联系人" width="140" align="center"></el-table-column>
          <el-table-column prop="contact_phone" label="紧急联系电话" align="center"></el-table-column>
          <el-table-column prop="contact_relation" label="关系" width="140" align="center"></el-table-column>
          <el-table-column prop="address" label="家庭地址" width="160" align="center"></el-table-column>
          <el-table-column
            label="操作"
            align="center"
            width="220"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                v-permisaction="['system:sysuser:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改</el-button>
              <el-button
                v-if="scope.row.userId !== 1"
                v-permisaction="['system:sysuser:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="params.pageIndex"
          :limit.sync="params.pageSize"
          @pagination="paginationChange"
        />
      </el-card>

      <!-- 添加或修改参数配置对话框 -->
      <el-dialog :title="title" :visible.sync="open" width="600px">
        <el-form ref="form" :model="form" :rules="rules" label-width="80px">
          <el-row>
            <el-col :span="12">
              <el-form-item label="学生姓名" prop="name">
                <el-input v-model="form.name" placeholder="请输入学生姓名" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="用户性别" prop="sex">
                <el-select v-model="form.sex" placeholder="请选择" :change="sexChange">
                  <el-option
                    v-for="dict in sexOptions"
                    :key="dict.dictValue"
                    :label="dict.dictLabel"
                    :value="dict.dictValue"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="手机号码" prop="phone">
                <el-input v-model="form.phone" placeholder="请输入手机号码" maxlength="11" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="联系人" prop="contact_by">
                <el-input v-model="form.contact_by" placeholder="请输入紧急联系人姓名" maxlength="50" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="联系电话" prop="contact_phone">
                <el-input v-model="form.contact_phone" placeholder="请输入紧急联系人电话" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="联系关系" prop="contact_relation">
                <el-input v-model="form.contact_relation" placeholder="请输入紧急联系人关系" type="password" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item prop="grade" label="学生年级">
                <el-select v-model="form.grade" placeholder="请选择学生年级">
                  <el-option
                    v-for="dict in gradeOptions"
                    :key="dict.label"
                    :label="dict.label"
                    :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item prop="class_by" label="学生班级">
                <el-select v-model="form.class_by" placeholder="请选择学生班级">
                  <el-option
                    v-for="dict in classByOptions"
                    :key="dict.label"
                    :label="dict.label"
                    :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="24">
              <el-form-item label="居住地址">
                <el-input v-model="form.address" type="textarea" placeholder="请输入居住地址" />
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="学生头像">
                <el-upload
                  class="avatar-uploader"
                  :action="url"
                  :data="{type:'1'}"
                  :show-file-list="false"
                  :on-success="handleAvatarSuccess"
                  :before-upload="beforeAvatarUpload"
                >
                  <img v-if="form.avatar" :src="form.avatar" class="avatar">
                  <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
              </el-form-item>
            </el-col>

          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="updateSubmitForm" v-if="form.student_id">修 改</el-button>
          <el-button type="primary" @click="submitForm" v-else>确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </el-dialog>
      <!--  头像放大查看    -->
      <el-dialog :title="title" :visible.sync="avatarOpen" width="300px">
        <img :src="avatarUrl" style="width: 260px;height: 240px" />
      </el-dialog>
    </template>
  </BasicLayout>
</template>

<script>
import { addStudent, deleteStudent, getStudent, updateStudent } from '@/api/student/student'

export default {
  data() {
    return {
      url: process.env.VUE_APP_BASE_API + '/api/v1/public/uploadFile',
      loading: false,
      studentData: [],
      total: 0,
      params: {
        pageSize: 10,
        pageIndex: 1
      },
      form: {
        name: '',
        student_id: '',
        sex: '',
        address: '',
        contact_by: '',
        contact_phone: '',
        contact_relation: '',
        phone: '',
        grade: '',
        class_by: '',
        avatar: ''
      },
      avatarOpen: false,
      avatarUrl: '',
      queryParams: {},
      open: false,
      title: '',
      sexOptions: [],
      rules: {
        name: [
          { required: true, message: '请输入学生姓名', trigger: 'blur' },
          { min: 2, max: 30, message: '长度在 2 到 30 个字符', trigger: 'blur' }
        ],
        sex: [
          { required: true, message: '请选择性别', trigger: 'blur' }
        ],
        grade: [
          { required: true, message: '请选择年级', trigger: 'blur' }
        ],
        class_by: [
          { required: true, message: '请选择班级', trigger: 'blur' }
        ],
        contact_phone: [
          { required: true, message: '请输入紧急联系电话', trigger: 'blur' }
        ],
        contact_by: [
          { required: true, message: '请输入紧急联系人', trigger: 'blur' }
        ]
      },
      gradeOptions: [
        { label: '高三', value: 12 },
        { label: '高二', value: 11 },
        { label: '高一', value: 10 }
      ],
      classByOptions: [
        { label: '一班', value: 1 },
        { label: '二班', value: 2 },
        { label: '三班', value: 3 }
      ]
    }
  },
  created() {
    this.getDate()
    this.getDicts('sys_user_sex').then(response => {
      this.sexOptions = response.data
    })
  },
  methods: {
    async getDate() {
      this.loading = true
      const ret = await getStudent(this.params)
      if (ret.code === 200) {
        this.studentData = ret.data.list
        this.total = ret.data.count
      }
      this.loading = false
    },
    handleAdd() {
      this.reset()

      this.open = true
      this.title = '添加学生信息'
    },
    handleExport() {
      window.open('http://127.0.0.1:8000/api/v1/export/学生列表')
    },
    async handleDelete(row) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async() => {
        const ret = await deleteStudent({ student_id: row.student_id })
        if (ret.code === 200) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          this.getDate()
        } else {
          this.$message({
            type: 'warn',
            message: ret.msg || '删除成功!'
          })
        }
      })
    },
    // 表单重置
    reset() {
      this.form = {
        name: '',
        sex: '',
        address: '',
        contact_by: '',
        contact_phone: '',
        contact_relation: '',
        phone: '',
        grade: '',
        class_by: '',
        avatar: ''
      }
      this.resetForm('form')
    },
    cancel() {
      this.open = false
    },
    handleAvatarSuccess(e) {
      this.form.avatar = e.data.full_path
    },
    beforeAvatarUpload(e) {

    },
    sexChange(e, a) {
      console.log('==sexChange==', e, a)
    },
    async submitForm() {
      this.$refs['form'].validate(async valid => {
        if (valid) {
          const ret = await addStudent(this.form)
          if (ret.code === 200) {
            this.params.pageIndex = 1
            this.msgSuccess('添加成功')
            this.open = false
            this.getDate()
          }
        }
      })
    },
    async updateSubmitForm() {
      this.$refs['form'].validate(async valid => {
        if (valid) {
          const ret = await updateStudent(this.form)
          if (ret.code === 200) {
            this.params.pageIndex = 1
            this.msgSuccess('修改成功')
            this.open = false
            this.getDate()
          }
        }
      })
    },
    openAvatarFn(row) {
      this.avatarOpen = true
      this.title = row.name + '的照片'
      this.avatarUrl = row.avatar
    },
    paginationChange(pages) {
      const { page, limit } = pages
      this.params.pageIndex = page
      this.params.pageSize = limit
      this.getDate()
    },
    handleUpdate(row) {
      this.title = '修改学生信息'
      this.form = row
      this.open = true
    }
  },
  computed: {
    sexTextMap() {
      return this.form.sex === 0 ? '男' : '女'
    }
  }
}
</script>

<style scoped>
.avatar-uploader .el-upload {
  border: 1px dashed red;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 60px;
  height: 60px;
  line-height: 40px;
  text-align: center;
}
.avatar {
  width: 60px;
  height: 60px;
  display: block;
}
</style>
