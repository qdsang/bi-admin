<template>
  <div class="page-container">
    <!-- <el-button type="primary" icon="el-icon-plus" size="small" class="add-source-btn" @click="handleNew()">
      {{ $t('dataSource.addDataSource') }}
    </el-button> -->
    <el-button type="primary" size="small" class="add-source-btn" @click="saveTableConfig()">
      {{ $t('common.confirm') }}
    </el-button>
    <el-table :data="tables">
      <el-table-column :label="$t('dataSource.tableName')" prop="table">
        <template slot-scope="scope">
          <el-input v-model="scope.row.table_alias" size="mini" placeholder="" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('dataSource.table_schema')" prop="tableSchema" />
      <el-table-column :label="$t('dataSource.table')" prop="table" />
      <el-table-column :label="$t('dataSource.linked')">
        <template slot-scope="scope">
          <el-switch v-model="scope.row.status" :active-text="$t('common.open')" :inactive-text="$t('common.closed')" />
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { tablesByBase, saveTableConfig } from '@/api/source'
import { parseTime } from '@/utils'

export default {
  data() {
    return {
      dialogVisible: false,
      tables: [],
      formData: {
        host: '',
        port: '',
        username: '',
        password: '',
        database: ''
      },
      showTableDialog: false,
      source_id: undefined
    }
  },
  created() {
  },
  mounted() {
    this.getList()
  },
  methods: {
    parseTime,
    async getList() {
      this.source_id = this.$route.params.id
      this.handleAddTable({ source_id: this.source_id })
    },
    // handleNew() {
    //   this.formData = {}
    //   this.dialogVisible = true
    // },
    // handleEdit(row) {
    //   this.formData = { ...row }
    //   this.dialogVisible = true
    // },
    async handleAddTable(row) {
      this.tables = []
      this.source_id = row.source_id
      const { data } = await tablesByBase(row.source_id)
      this.tables = data.list.map(item => {
        item.status = item.status === 1
        return item
      })
    },
    async saveTableConfig() {
      const tables = this.tables.map(item => {
        item.status = item.status ? 1 : 0
        return item
      })
      await saveTableConfig({ source_id: this.source_id, tables })
      this.showTableDialog = false
      this.$message.success('Success!')
    }
  }
}
</script>

<style lang="scss" scoped>
.page-container {
  padding: 24px;
}
.add-source-btn {
  margin-bottom: 24px;
}
.source-input {
  width: 250px;
}
</style>
