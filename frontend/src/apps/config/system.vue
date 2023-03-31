<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { PrimaryTableCol, TableRowData } from "tdesign-vue-next"

import { NaApi } from "@/api"
import { ConfigItem } from "@/api/native/config"

import ConfigUpdate from "./config_update.vue"

@Component({
    components: { ConfigUpdate }
})
export default class VendorAlibabaUpdate extends Vue {
    public loading = true

    @Ref
    public updateModal!: ConfigUpdate

    // 初始化

    public created() {
        this.getConfigList()
    }

    // 列表

    public configList: ConfigItem[] = []

    async getConfigList() {
        const res = await NaApi.config.list({}).finally(() => this.loading = false)
        this.configList = res.Items
    }

    // 表格定义

    public tableColumns: PrimaryTableCol<TableRowData | ConfigItem>[] = [
        { colKey: 'Description', title: '参数名', ellipsis: true },
        { colKey: 'Value', title: '参数值', ellipsis: true },
        { colKey: 'Operation', title: '操作', width: "120px" },
    ]
}
</script>

<template>
    <t-space fixed direction="vertical">
        <t-breadcrumb>
            <t-breadcrumb-item to="/">
                首页
            </t-breadcrumb-item>
            <t-breadcrumb-item>
                系统参数
            </t-breadcrumb-item>
        </t-breadcrumb>

        <t-card :loading="loading" title="参数列表" hover-shadow header-bordered>
            <template #subtitle>
                记录总数: {{ configList?.length || 0 }}
            </template>
            <t-table :loading="loading" :data="configList" :columns="tableColumns" :hover="true" row-key="Id"
                cell-empty-content="--">
                <template #Operation="{ row }">
                    <t-link theme="primary" hover="color" @click="updateModal.open(row)">
                        修改
                    </t-link>
                </template>
            </t-table>
        </t-card>

        <ConfigUpdate ref="updateModal" @submit="getConfigList" />
    </t-space>
</template>