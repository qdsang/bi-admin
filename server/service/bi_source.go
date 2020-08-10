package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func GetSourceList(info request.BiSourceListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiSource{})
	var itemList []model.BiSource
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Name != "" {
	// 	db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	// }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func SourceCreate(item *model.BiSource) (err error) {
	// if (!global.GVA_DB.First(&model.BiSource{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	item.SourceID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func SourceUpdate(item *model.BiSource) (err error) {
	// if (!global.GVA_DB.First(&model.BiSource{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	err = global.GVA_DB.Where("source_id = ?", item.SourceID).First(&model.BiSource{}).Updates(&item).Error
	return err
}

func SourceDelete(item *model.BiSource) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiSource{}, "source_id = ?", item.SourceID).Error
	return err
}

func GetSourceById(sourceID string) (err error, item model.BiSource) {
	err = global.GVA_DB.Where("source_id = ?", sourceID).First(&item).Error
	return
}

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func GetSourceTables(id string) (err error, list interface{}, total int) {
	sql := "SELECT * FROM information_schema.tables where TABLE_SCHEMA not in (\"information_schema\", \"performance_schema\", \"mysql\", \"sys\")"
	err, sqlList, sqlTotal := GetSourceSql(id, sql)
	if err != nil {
		return err, list, 0
	}

	tables := make([]map[string]interface{}, 0)

	for i := 0; i < sqlTotal; i++ {
		m := sqlList[i].(map[string]interface{})

		tableCatalog := m["TABLE_CATALOG"]
		tableSchema := m["TABLE_SCHEMA"]
		tableName := m["TABLE_NAME"]
		tableComment := m["TABLE_COMMENT"]
		tableType := m["TABLE_TYPE"]

		// created_at: "2020-07-31T03:45:21.542Z"
		// creator: 1
		// database: "3d19dc5d-4027-491b-ba6a-018913d67441"
		// status: 1
		// table: "order_analysis"
		// table_alias: "订单分析表"
		// updated_at: "2020-07-31T03:45:21.538Z"
		tables = append(tables, map[string]interface{}{
			"tableCatalog": tableCatalog,
			"tableSchema":  tableSchema,
			"tableName":    tableName,
			"tableComment": tableComment,
			"tableType":    tableType,

			"table":       tableName,
			"table_alias": tableComment,
		})
		// fmt.Println(tableCatalog, tableSchema, tableName, tableComment, tableType)
		// fmt.Println(m)
	}

	// fmt.Println(tables)
	return err, tables, len(tables)
}

func GetSourceSql(id string, sql string) (err error, list []interface{}, total int) {
	err, db := GetDb(id)
	if err != nil {
		return err, list, 0
	}
	defer db.Close()
	if db == nil {
		global.GVA_LOG.Error("db异常", db)
	}
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		global.GVA_LOG.Error("rows异常", err)
		return err, list, 0
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	rowList := make([]interface{}, 0)

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			global.GVA_LOG.Error("rows Scan异常", err)
			return err, list, 0
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			switch (*val).(type) {
			case []byte:
				m[colName] = string((*val).([]byte))
			default:
				m[colName] = *val
			}
			// m[colName] = reflect.ValueOf(val)
			// reflect.ValueOf(val)
		}
		rowList = append(rowList, m)
	}
	return err, rowList, len(rowList)
}

func GetDb(id string) (err error, db *gorm.DB) {
	err, sourceItem := GetSourceById(id)
	if err != nil {
		return err, db
	}

	driver := "mysql"
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	source := ""
	if sourceItem.Username != "" {
		source = sourceItem.Username + ":" + sourceItem.Password + "@" + source
	}
	if sourceItem.Port == 0 {
		sourceItem.Port = 3306
	}
	source = source + "(" + sourceItem.Host + ":" + strconv.Itoa(sourceItem.Port) + ")"
	source = source + "/" + sourceItem.Database + "?charset=utf8&parseTime=True&loc=Local" // + sourceItem.Config
	// global.GVA_LOG.Info("启动", err)
	// 123456
	if db, err := gorm.Open(driver, source); err != nil {
		global.GVA_LOG.Error(driver+"启动异常", err)
		// os.Exit(0)
		return err, db
	} else {
		db.DB().SetMaxIdleConns(1)
		db.DB().SetMaxOpenConns(10)
		global.GVA_LOG.Info("启动成功:" + source)
		return nil, db
	}
}

// type Result struct {
// 	Name   string
// 	Gender string
// }
// // 由于需要返回多条数据所以需要使用切片去接收，如果有且仅有一条数据可以不使用切片
// list := []Result{}
// //调用原生sql语句
// db.Raw("SELECT name, gender FROM user_infos WHERE name=?", "连少").Find(&list)
// // 打印结果
// fmt.Println(list)
