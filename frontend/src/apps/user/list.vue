<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { PrimaryTableCol, TableRowData } from "tdesign-vue-next"

import { NaApi } from "@/api"
import { UserItem } from "@/api/native/user"

import UserUpdate from "./update.vue"

@Component({
    components: { UserUpdate }
})
export default class UserList extends Vue {
    public loading = true

    @Ref
    public updateModal!: UserUpdate

    // 初始化

    public created() {
        this.getUserList()
    }

    // 列表

    public userList: UserItem[] = []

    async getUserList() {
        const res = await NaApi.user.list({}).finally(() => {
            this.loading = false
        })
        this.userList = res.Items
    }

    // 表格定义

    public tableColumns: PrimaryTableCol<TableRowData | UserItem>[] = [
        { colKey: 'Username', title: '用户名', ellipsis: true },
        { colKey: 'Email', title: '邮箱', ellipsis: true },
        { colKey: 'Level', title: '用户组', ellipsis: true },
        { colKey: 'ArtworkQuota', title: '绘图配额', ellipsis: true },
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
                用户管理
            </t-breadcrumb-item>
        </t-breadcrumb>

        <t-card :loading="loading" title="用户列表" hover-shadow header-bordered>
            <template #subtitle>
                记录总数: {{ userList?.length || 0 }}
            </template>
            <t-table :loading="loading" :data="userList" :columns="tableColumns" :hover="true" row-key="Id"
                cell-empty-content="--">
                <template #Operation="{ row }">
                    <t-link theme="primary" hover="color" @click="updateModal.open(row)">
                        修改
                    </t-link>
                </template>
            </t-table>
        </t-card>

        <UserUpdate ref="updateModal" @submit="getUserList" />
    </t-space>
</template>