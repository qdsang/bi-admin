package request

import "bi-admin/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}