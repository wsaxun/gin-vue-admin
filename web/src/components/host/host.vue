<template>
    <div class="gva-table-box">
        <el-table ref="multipleTableRef" :data="tableData" style="width: 100%" height="300" row-key="id"
            @select="handleSelectChange" @select-all="handleSelectionAll">
            <el-table-column type="selection" width="55" />
            <el-table-column property="hostname" label="主机名" />
            <el-table-column property="env" label="环境" />
            <el-table-column property="ip" label="ip" />
            <el-table-column property="description" label="描述" />
        </el-table>
    </div>
    <div class="gva-pagination">
        <el-pagination :page-sizes="[1, 3, 5, 10]" :page-size=pageSize :current-page="currentPage" :total="total"
            :style="{ 'justify-content': 'center' }" layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>
</template>

<style scoped></style>

<script setup>
import { ref, reactive } from 'vue'
import { ElTable } from 'element-plus'
import { isNull } from '@/utils/checkParam';
import { getHost } from '@/api/host'
import { piniaHosts } from '@/pinia/modules/hosts'

const multipleTableRef = ref(null)
const hostStore = piniaHosts()
const currentPage = ref(1)
const pageSize = ref(3)
const tableData = reactive([])
const total = ref(0)

const handleSizeChange = (val) => {
    pageSize.value = val
    fetchData()
}

const handleCurrentChange = (val) => {
    currentPage.value = val
    fetchData()
}

// 单选
const handleSelectChange = (_, row) => {
    hostStore.handleSelectChange(row)
}

// 全选
const handleSelectionAll = (rows) => {
    hostStore.selectCurAll(rows, tableData)
}

// 分页状态选中维护
const maintainSelection = () => {
    hostStore.resetCurrentPage()
    tableData.forEach(function (tableItem) {
        if (hostStore.multipleSelection.length > 0) {
            hostStore.multipleSelection.forEach(function (multiItem) {
                if (tableItem.id === multiItem.id) {
                    hostStore.pushCurrentPage(tableItem)
                    return
                }
            })
        }
    })
}
// 选择指定行
const toggleSelection = () => {
    if (hostStore.currentPageSelection) {
        hostStore.currentPageSelection.forEach((row) => {
            multipleTableRef.value.toggleRowSelection(row, undefined);
        });
    } else {
        multipleTableRef.value.clearSelection()
    }
}

const resetPage = (val) => {
    currentPage.value = val
    hostStore.resetCurrentPage()
    hostStore.resetMult()
    multipleTableRef.value.clearSelection()
}

const fetchData = async () => {
    let res = await getHost({ fields: "id,hostname,env,ip,description", page: currentPage.value, pageSize: pageSize.value })
    tableData.splice(0)
    if (res.code === 0 && !isNull(res.data.list)) {
        res.data.list.forEach(ele => {
            tableData.push(ele)
        })
        total.value = res.data.total
        pageSize.value = res.data.pageSize
        currentPage.value = res.data.page
    }
    maintainSelection()
    toggleSelection()
}

fetchData()

defineExpose({
    resetPage,
    fetchData
})

</script>