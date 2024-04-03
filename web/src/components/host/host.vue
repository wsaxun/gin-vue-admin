<template>
    <div class="gva-table-box">
        <el-table ref="multipleTableRef" :data="tableData" style="width: 100%" height="300" row-key="id" 
            @select="handleSelectChange"
            @select-all="handleSelectionAll">
            <el-table-column type="selection" width="55" />
            <el-table-column property="hostname" label="主机名" />
            <el-table-column property="env" label="环境" />
            <el-table-column property="ip" label="ip" />
            <el-table-column property="description" label="描述" />
        </el-table>
    </div>
    <div class="gva-pagination">
        <el-pagination 
            :page-sizes="[1, 3, 5, 10]"
            :page-size=pageSize
            :current-page="currentPage"
            :total="total"
            :style="{'justify-content':'center'}"
            layout="total, sizes, prev, pager, next, jumper"  @size-change="handleSizeChange"
            @current-change="handleCurrentChange" />
    </div>
</template>

<style scoped>

</style>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElTable } from 'element-plus'
import { isNull } from '@/utils/checkParam';
import {getHost} from '@/api/host'

const multipleTableRef = ref(null)
const multipleSelection = ref([])
const currentPageSelection = ref([])
const currentPage = ref(1)
const pageSize =  ref(3)
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
    var flag = false
    var i = 0
    currentPageSelection.value.forEach(function(val, index){
        if (val.id === row.id) {
            flag = true
            i = index
        }
    })
    if (flag) {
        currentPageSelection.value.splice(i, 1)
        var n = null
        multipleSelection.value.forEach(function(val, index){
            if (val.id === row.id) {
                n = index
            }
            if (n !== null) {
                multipleSelection.value.splice(n, 1)
            }
        })
    } else {
        currentPageSelection.value.push(row)
        multipleSelection.value.push(row)
    }
}

// 全选
const handleSelectionAll = (rows) => {

    if (rows.length === 0 ) {
        currentPageSelection.value.splice(0)
        tableData.forEach(function(tableVal){
            var n = null
            multipleSelection.value.forEach(function(val, index){
                if (val.id === tableVal.id) {
                    n = index
                }
            })
            if (n !== null ) {
                multipleSelection.value.splice(n, 1)
            }
        })
        return
    }
    if (multipleSelection.value.length === 0){
        tableData.forEach(function(tableVal){
            currentPageSelection.value.push(tableVal)
            multipleSelection.value.push(tableVal)
        })
        return
    }

    currentPageSelection.value.splice(0)
    tableData.forEach(function(tableVal){
        currentPageSelection.value.push(tableVal)
        var n = null
        multipleSelection.value.forEach(function(val, index){
            if (val.id === tableVal.id) {
                n = index
                return
            }
        })
        if (n === null ) {
            multipleSelection.value.push(tableVal)
        }       
    })
    // console.log(currentPageSelection.value)
    // console.log(multipleSelection.value)
}

// 分页状态选中维护
const maintainSelection = () => {
    currentPageSelection.value.splice(0)
    tableData.forEach(function(tableItem){
        multipleSelection.value.forEach(function(multiItem){
            if (tableItem.id === multiItem.id){
                currentPageSelection.value.push(tableItem)
                return
            }
        })
    })
}
// 选择指定行
const toggleSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {
      multipleTableRef.value.toggleRowSelection(row, undefined);
    });
  } else {
    multipleTableRef.value.clearSelection()
  }
}

const resetPage = (val) => {
    currentPage.value = val
    currentPageSelection.value.splice(0)
    multipleSelection.value.splice(0)
    multipleTableRef.value.clearSelection()
}

const  fetchData = async () => {
    let res = await getHost({fields: "id,hostname,env,ip,description",page: currentPage.value, pageSize: pageSize.value})
    tableData.splice(0)
    if (res.code === 0 && !isNull(res.data.list)){
        res.data.list.forEach(ele => {
            tableData.push(ele)
        })
        total.value = res.data.total
        pageSize.value = res.data.pageSize
        currentPage.value = res.data.page
    }
    maintainSelection()
    toggleSelection(currentPageSelection.value)
}

fetchData()

defineExpose({
    multipleSelection,
    resetPage,
    fetchData
})

</script>