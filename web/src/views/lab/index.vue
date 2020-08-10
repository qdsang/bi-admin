<template>
  <div v-loading="loading" class="container">
    <el-card body-style="padding: 0px;" class="lab-list" shadow="never">
      <div slot="header">
        <span>{{ $t('common.lab') }}</span>
        <i class="el-icon-plus" @click="addLab" />
      </div>
      <ul>
        <draggable v-model="labList" :group="{name: 'lab',pull: true}" class="draggable-wrapper" @change="handleOrderChange">
          <li v-for="item in labList" :key="item.lab_id" :class="{'lab-list-item': true, 'high-light-lab': currentLab.lab_id === item.lab_id}">
            <span @click="switchDb(item)">
              <i class="el-icon-document" />
              <span>{{ item.name }}</span>
            </span>
            <el-dropdown szie="mini" trigger="click" @command="handleCommand">
              <span class="el-dropdown-link">
                <i class="el-icon-more" />
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item
                  :command="{
                    type: 'edit',
                    target: item
                  }"
                >
                  {{ $t('common.edit') }}
                </el-dropdown-item>
                <el-dropdown-item
                  :command="{
                    type: 'delete',
                    target: item
                  }"
                >
                  {{ $t('common.delete') }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          <!-- {{ item.desc }} -->
          </li>
        </draggable>
      </ul>
    </el-card>
    <labItem :lab="currentLab" />
    <el-dialog :title="$t('lab.addOrEditLab')" width="750px" :visible.sync="editDialogVisible">
      <el-form label-width="160px">
        <el-form-item :label="$t('lab.labName')">
          <el-input v-model="dbObj.name" size="small" style="width: 450px;" :placeholder="$t('lab.labNamePlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('lab.labDesc')">
          <el-input v-model="dbObj.desc" type="textarea" :rows="5" size="small" style="width: 450px;" :placeholder="$t('lab.labDescPlaceholder')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" size="small" @click="editDialogVisible = false"> {{ $t('common.cancel') }}</el-button>
        <el-button type="primary" size="small" @click="handleSubmit"> {{ $t('common.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import draggable from 'vuedraggable'
import labItem from './labItem'
import { addLab, updateLab, labList, deleteLab, dbOrder } from '@/api/lab'

export default {
  components: { labItem, draggable },
  data() {
    return {
      labList: [],
      currentLab: undefined,
      editDialogVisible: false,
      dbObj: {},
      loading: false,
      isCollapse: false
    }
  },
  computed: {
    projectPath() {
      const projectId = this.$route.params.projectid
      if (projectId) {
        return '/' + projectId
      } else {
        return ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      labList().then(resp => {
        this.loading = false
        this.labList = []
        this.labList = this.labList.concat(resp.data.labs)
        const lab = this.labList.find(item => item.lab_id === this.$route.query.id)
        if (lab) {
          this.currentLab = lab
        } else {
          this.currentLab = this.labList[0]
        }
        if (this.currentLab) {
          this.$router.push(`${this.projectPath}/lab?id=${this.currentLab.lab_id}`).catch(_ => {})
        }
      })
    },
    switchDb(db) {
      if (db.lab_id === this.currentLab.lab_id) {
        this.getList()
        return
      }
      // this.$confirm('确定要离开当前页面吗?系统可能不会保存您所做的更改。', '提示').then(() => {
      this.currentLab = db
      this.$router.push(`${this.projectPath}/lab?id=${this.currentLab.lab_id}`)
      // })
    },
    addLab() {
      this.dbObj = {}
      this.editDialogVisible = true
    },
    editLab(db) {
      this.dbObj = Object.assign({}, db)
      this.editDialogVisible = true
    },
    handleCommand(cmd) {
      if (cmd.type === 'edit') {
        this.editLab(cmd.target)
      } else {
        this.deleteLab(cmd.target)
      }
    },
    handleSubmit() {
      if (this.dbObj.lab_id) {
        updateLab(this.dbObj).then(resp => {
          this.getList()
          this.editDialogVisible = false
        })
      } else {
        addLab(this.dbObj).then(resp => {
          this.getList()
          this.editDialogVisible = false
        })
      }
    },
    handleOrderChange(evt) {
      const data = {
        order: this.labList.map(item => item.lab_id)
      }
      dbOrder(data)
    },
    deleteLab(db) {
      this.$confirm(this.$t('lab.deleteConfirm', db.name), this.$t('common.confirm')).then(() => {
        deleteLab({ lab_id: db.lab_id }).then(() => {
          this.getList()
          this.$message({
            type: 'success',
            message: this.$t('common.deleteSuccess')
          })
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  min-height: calc(100vh - 62px);
  align-items: stretch;
  .lab-list {
    width: 250px;
    min-height: 100%;
    padding: 20px 10px;
    /deep/ .el-card__header {
      div {
        display: flex;
        justify-content: space-between;
        font-size: 14px;
        color: #606266;
        i {
          cursor: pointer;
        }
      }
      padding: 5px 0px;
    }
    .lab-list-item {
      display: flex;
      justify-content: space-between;
      line-height: 35px;
      font-size: 14px;
      cursor: pointer;
      color: #606266;
      i {
        margin-right: 10px;
        line-height: 35px;
      }
    }
    .high-light-lab {
      color: #205cd8;
    }
  }
  .lab-wrapper {
    width: 100%;
  }
}
</style>
