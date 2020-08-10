<template>
  <div class="page-container">
    <el-button type="primary" icon="el-icon-plus" size="small" class="add-project-btn" @click="handleNew()">
      {{ $t('project.addTask') }}
    </el-button>
    <el-table :data="list" border highlight-current-row stripe class="project-table">
      <el-table-column :label="$t('project.name')" prop="name" align="center" width="200px" />
      <el-table-column :label="$t('project.desc')" prop="desc" />
      <el-table-column :label="$t('project.type')" prop="type" />
      <el-table-column :label="$t('project.status')" prop="status" />
      <el-table-column :label="$t('project.created_at')" prop="created_at" width="155px">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <!-- <el-table-column label="status" prop="status" /> -->
      <el-table-column :label="$t('common.operation')" align="center" width="290px">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleDashboard(scope.row)">
            {{ $t('common.dashboard') }}
          </el-button>
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)">
            {{ $t('common.edit') }}
          </el-button>
          <el-button type="danger" size="mini" @click="handleDelete(scope.row)">
            {{ $t('common.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="$t('project.addTask')" :visible.sync="dialogVisible" width="600px">
      <el-form label-width="100px" size="small">
        <el-form-item :label="$t('project.name') + ':' ">
          <el-input v-model="formData.name" class="project-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('project.desc') + ':' ">
          <el-input v-model="formData.desc" class="project-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('project.type') + ':' ">
          <el-input v-model="formData.type" class="project-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('project.status') + ':' ">
          <el-input v-model="formData.status" class="project-input" placeholder="" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialogVisible = false">{{ $t('common.cancel') }}</el-button>
        <el-button size="small" type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { addProject, projectList, updateProject, deleteProject } from '@/api/project'
import { parseTime } from '@/utils'

export default {
  data() {
    return {
      dialogVisible: false,
      list: [],
      tables: [],
      formData: {
        host: '',
        port: '',
        username: '',
        password: '',
        database: ''
      },
      project_id: undefined
    }
  },
  created() {
    this.getList()
  },
  methods: {
    parseTime,
    async getList() {
      const { data } = await projectList()
      this.list = data.list
    },
    handleNew() {
      this.formData = {}
      this.dialogVisible = true
    },
    handleDashboard(row) {
      this.$router.push(`/${row.project_id}/dashboard`).catch(_ => {})
    },
    handleEdit(row) {
      this.formData = { ...row }
      this.dialogVisible = true
    },
    handleDelete(row) {
      this.$confirm(this.$t('project.deleteConfirm'), this.$t('common.confirm')).then(async() => {
        await deleteProject({ project_id: row.project_id })
        this.$message.success(this.$t('common.deleteSuccess'))
        this.$message.warning('But You cant really delete it, because we need this demo.')
        this.getList()
      })
    },
    async handleSubmit() {
      if (this.formData.project_id) {
        await updateProject(this.formData)
      } else {
        await addProject(this.formData)
      }
      this.getList()
      this.dialogVisible = false
    }
  }
}
</script>

<style lang="scss" scoped>
.page-container {
  padding: 24px;
}
.add-project-btn {
  margin-bottom: 24px;
}
.project-input {
  width: 250px;
}
</style>
