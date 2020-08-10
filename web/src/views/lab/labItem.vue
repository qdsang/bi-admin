<template>
  <div class="lab-wrapper">
    <div class="tool-bar">
      <div>
        <span class="db-name">{{ lab.name }}</span>
        <span>{{ lab.desc }}</span>
      </div>
      <div>
        <el-button type="primary" size="mini" @click="handleExeSql">
          {{ $t('lab.exec') }}
        </el-button>
        <el-button type="primary" size="mini" @click="handleSqlFormat">
          {{ $t('lab.format') }}
        </el-button>
        <!-- <el-form-item :label="$t('common.dataSource')+':'"> -->
        <el-select v-model="selectedBase" size="mini" filterable :placeholder="$t('dataSource.sourcePlaceholder')" style="width:200px;" clearable @change="handleBaseChange">
          <el-option v-for="item in baseList" :key="item.source_id" :label="item.base_alias || item.database" :value="item.source_id" />
        </el-select>
        <!-- </el-form-item> -->
        <!-- <el-button type="primary" size="mini" @click="handleLinkChart">
          {{ $t('lab.addChart') }}
        </el-button> -->
      </div>
    </div>
    <sql-edit ref="sqlEditor" :value="sql" :options="sqlEditOption" @change="handleSqlEditChange" />
    <component :is="currentType.componentName" :data="chartData" :schema="schema" :chart-style="chartStyle" class="visualize-window" />
  </div>
</template>
<script>
import sqlPrettier from 'sql-prettier'
import exeSql from '@/api/exeSql'
import { updateLab } from '@/api/lab'
import { sourceList } from '@/api/source'
import sqlEdit from '@/components/sqlEdit'
import DataTable from '@/widgets/DataTable'
import chartTypeList from '@/utils/chartTypeList'

export default {
  components: { sqlEdit, DataTable },
  props: {
    lab: {
      required: false,
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data() {
    const self = this
    return {
      sqlEditOption: {
        mode: 'sql',
        // theme: '3024-day',
        lineNumbers: true,
        // lineWrapping: false,
        autoCloseBrackets: true,
        matchBrackets: true,
        foldGutter: true,
        extraKeys: {
          'Cmd-Enter': () => { self.handleExeSql() },
          'Ctrl-Enter': () => { self.handleExeSql() },
          'Ctrl-Space': 'autocomplete'
        }
      },
      sql: '',
      baseList: [],
      loading: false,
      selectedBase: null,
      data: [],
      schema: [],
      chartTypeList,
      chartType: 'table',
      chartStyle: {}
    }
  },
  computed: {
    editor() {
      return this.$refs.sqlEditor.editor
    },
    chartData() {
      return this.currentType.dataTransfer ? this.currentType.dataTransfer(this.data, this.schema) : this.data
    },
    currentType() {
      return chartTypeList.find(item => item.type === this.chartType)
    }
  },
  watch: {
    'lab.lab_id': {
      immediate: true,
      handler(value) {
        if (!value) return
        this.getList(value)
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    async getList() {
      this.selectedBase = this.lab.source_id
      this.sql = this.lab.sql || 'select 1'
      const { data } = await sourceList()
      this.baseList = data.list
    },
    handleBaseChange() {
      this.handleLabChange()
    },
    handleSqlEditChange(code, k2) {
      // console.log(code, k2)
      this.sql = code
    },
    handleSqlFormat() {
      this.sql = sqlPrettier.format(this.sql)
    },
    handleExeSql() {
      console.log(this.selectedBase)
      this.exeSql(this.sql)
    },
    handleLinkChart() {

    },
    handleDelete(chart) {
      this.$confirm(this.$t('lab.removeChartConfirm'), this.$t('common.confirm'), {
        type: 'warning'
      }).then(() => {
        // this.lab.content.layout = layout
        // Promise.all([updateLab(this.lab)]).then(resp => {
        //   this.getList()
        //   this.$message({
        //     type: 'success',
        //     message: this.$t('common.deleteSuccess')
        //   })
        // })
      })
    },
    handleLabChange() {
      this.lab.sql = this.sql
      this.lab.source_id = this.selectedBase
      updateLab(this.lab)
    },
    handleResize(i, newH, newW, newHPx, newWPx) {
      this.$refs[`chartInstance${i}`][0].$children[0].$emit('resized')
    },
    exeSql(sqlSentence) {
      const source_id = this.selectedBase
      if (!sqlSentence || !source_id) {
        this.$message.warning(this.$t('lab.chartQueryException'))
        return
      }
      this.loading = true
      exeSql().fetch({ source_id: source_id, sql: sqlSentence }).then(resp => {
        if (resp.data.length > 0) {
          this.data = resp.data
          const schema = []
          const dataOne = resp.data[0]
          for (const key in dataOne) {
            schema.push({ name: key, Column: key, Type: 'text' })
          }
          // Column: "count"
          // Type: "int(11)"
          // availableFunc: Array(6)
          // calculFunc: "sum"
          // name: "count(合计)"
          this.schema = schema
        } else {
          this.data = this.schema = []
        }
        this.handleLabChange()
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.tool-bar {
  display: flex;
  justify-content: space-between;
  border-top: none;
  height: 45px;
  line-height: 45px;
  color:#303133;
  padding: 0 10px;
  position: relative;
  .db-name {
    font-size: 1.2em;
    font-weight: 600;
    color: #909399;
    margin-left: 0;
  }
  span {
    color: #C0C4CC;
    font-size: 0.8em;
    margin-left: 10px;
  }
}
.visualize-card {
  /deep/ .el-card__header {
    padding: 0;
    .operation-bar {
      font-size: 14px;
      display: flex;
      justify-content: space-between;
      height: 35px;
      padding: 0 10px;
      line-height: 35px;
      z-index: 9;
      i {
        margin-right: 10px;
        color: #409EFF;
        cursor: pointer;
      }
    }
  }
}
.welcome-container {
  text-align: center;
  height: 500px;
  color:#C0C4CC;
  /deep/ .el-button {
    margin-top: 200px;
    margin-bottom: 25px;
  }
}
</style>
