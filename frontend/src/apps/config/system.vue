<script lang="ts">
import { markRaw } from "vue"
import { Ref, Component, Vue } from "vue-facing-decorator"

import { Input, PrimaryTableCol, TableRowData, PrimaryTableRowValidateContext, PrimaryTable } from "tdesign-vue-next"

import { NaApi } from "@/api"
import { ConfigItem } from "@/api/native/config"

@Component
export default class VendorAlibabaUpdate extends Vue {

    // 初始化

    public created() {
        this.getConfigList()
    }

    // 列表

    public configList: ConfigItem[] = []

    async getConfigList() {
        const res = await NaApi.config.list({})
        this.configList = res.Items
    }

    // 保存数据

    public formSubmit(row: ConfigItem) {
        const form = this.editableRowData[row.Id]
        const query = { Id: row.Id, Value: form.Value }
        NaApi.config.update(query).then(() => {
            Object.assign(row, form)
            this.onCancel(row)
        })
    }

    // 表格编辑

    public editableRows: number[] = []
    public editableRowData: Record<number, ConfigItem> = {}

    public onEdit(row: ConfigItem) {
        if (!this.editableRows.includes(row.Id)) {
            this.editableRows.push(row.Id)
        }
        this.editableRowData[row.Id] = { ...row }
    }

    public onCancel(row: ConfigItem) {
        delete this.editableRowData[row.Id]
        const idx = this.editableRows.findIndex(t => t === row.Id)
        idx > -1 && this.editableRows.splice(idx, 1)
    }

    public onSubmit(row: ConfigItem) {
        this.tableRef.validateRowData(row.Id).then((params: PrimaryTableRowValidateContext<TableRowData>) => {
            if (params.result.length == 0) {
                this.formSubmit(row)
            }
        })
    }

    public onRowValidate(params: PrimaryTableRowValidateContext<TableRowData>) {
        if (params.trigger != 'self') {
            return
        }
        const { row, value } = params.result[0]
        const k = params.result[0].col.colKey || ""
        Object.assign(this.editableRowData[row.Id], { [k]: value })
    }

    // 表格定义

    @Ref
    public tableRef!: any

    public tableColumns: PrimaryTableCol<TableRowData | ConfigItem>[] = [
        { colKey: 'Description', title: '配置项', ellipsis: true },
        {
            colKey: 'Value',
            title: '配置值',
            edit: {
                component: markRaw(Input),
                props: {
                    clearable: true,
                    autofocus: true,
                },
                rules: [
                    { required: true, message: '不能为空' },
                    { max: 64, message: '字符数量不能超过 64' },
                ],
                validateTrigger: 'change',
                defaultEditable: false,
                showEditIcon: false,
            },
        },
        { colKey: 'Operation', title: '操作', width: "120px" },
    ]
}
</script>

<template>
    <t-table ref="tableRef" :data="configList" :columns="tableColumns" :editable-row-keys="editableRows" row-key="Id"
        cell-empty-content="--" @row-validate="onRowValidate">
        <template #Operation="{ row }">
            <template v-if="editableRows.includes(row.Id)">
                <t-link theme="primary" hover="color" @click="onSubmit(row)">
                    保存
                </t-link>
                <t-link theme="default" hover="color" @click="onCancel(row)">
                    取消
                </t-link>
            </template>
            <template v-else>
                <t-link theme="primary" hover="color" @click="onEdit(row)">
                    编辑
                </t-link>
            </template>
        </template>
    </t-table>
</template>