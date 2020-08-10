<template>
  <div class="page-container">
    <el-button type="primary" icon="el-icon-plus" size="small" class="add-task-btn" @click="handleNew()">
      {{ $t('task.addTask') }}
    </el-button>
    <el-table :data="list" border highlight-current-row stripe class="task-table">
      <el-table-column :label="$t('task.name')" prop="name" align="center" width="200px" />
      <el-table-column :label="$t('task.desc')" prop="desc" />
      <el-table-column :label="$t('task.type')" prop="type" />
      <el-table-column :label="$t('task.status')" prop="status" />
      <el-table-column :label="$t('task.cron')" prop="cron" />
      <el-table-column :label="$t('task.status')" prop="status" width="155px">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <!-- <el-table-column label="status" prop="status" /> -->
      <el-table-column :label="$t('common.operation')" align="center" width="290px">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)">
            {{ $t('common.edit') }}
          </el-button>
          <el-button type="danger" size="mini" @click="handleDelete(scope.row)">
            {{ $t('common.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="$t('task.addTask')" :visible.sync="dialogVisible" width="600px">
      <el-form label-width="100px" size="small">
        <el-form-item :label="$t('task.name') + ':' ">
          <el-input v-model="formData.name" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.desc') + ':' ">
          <el-input v-model="formData.desc" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.type') + ':' ">
          <el-input v-model="formData.type" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.cron') + ':' ">
          <el-input v-model="formData.cron" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.status') + ':' ">
          <el-input v-model="formData.status" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.startDate') + ':' ">
          <el-input v-model="formData.startDate" class="task-input" placeholder="" />
        </el-form-item>
        <el-form-item :label="$t('task.endDate') + ':' ">
          <el-input v-model="formData.endDate" class="task-input" placeholder="" />
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
import { addTask, taskList, updateTask, deleteTask } from '@/api/task'
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
      task_id: undefined
    }
  },
  created() {
    this.getList()
  },
  methods: {
    parseTime,
    async getList() {
      const { data } = await taskList()
      this.list = data.list
    },
    handleNew() {
      this.formData = {}
      this.dialogVisible = true
    },
    handleEdit(row) {
      this.formData = { ...row }
      this.dialogVisible = true
    },
    handleDelete(row) {
      this.$confirm(this.$t('task.deleteConfirm'), this.$t('common.confirm')).then(async() => {
        await deleteTask({ task_id: row.task_id })
        this.$message.success(this.$t('common.deleteSuccess'))
        this.$message.warning('But You cant really delete it, because we need this demo.')
        this.getList()
      })
    },
    async handleSubmit() {
      if (this.formData.task_id) {
        await updateTask(this.formData)
      } else {
        await addTask(this.formData)
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
.add-task-btn {
  margin-bottom: 24px;
}
.task-input {
  width: 250px;
}
</style>
