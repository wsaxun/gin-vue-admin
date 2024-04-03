<template>
    <el-row :gutter="20">
        <el-col :span="3">
            <div class="gva-table-box">
                <el-menu @select="handleSelect" :default-active="defaultIndex">
                    <el-menu-item :index="item.id" :key="item.id" v-for="item in items" @click="toolView(item.id)">{{
        item.name }}</el-menu-item>
                </el-menu>
            </div>
        </el-col>
        <el-col :span="21">
            <el-row :gutter="20">
                <el-col :span="24">
                    <div class="gva-table-box rowContainer" style="margin-bottom: 20px;">
                        <el-input class="toolWidth" v-model="serachInput" laceholder="Please Input" clearable
                            @clear="toolView(toolTypeID)" @keyup.enter="serach">
                            <template #append>
                                <el-button :icon="Search" @click="serach" />
                            </template>
                        </el-input>
                        <el-button type="primary" @click="addTool">新建</el-button>
                    </div>
                </el-col>
            </el-row>
            <el-row :gutter="20">
                <el-col :span="24">
                    <div class="flex flex-wrap gap-4 gva-table-box">
                        <el-card shadow="hover" :key="tool.id" @click="editTool(tool.id)" v-for="tool in tools"
                            class="card_item toolWidth">
                            <div  style="display: flex; align-items: center; justify-content: space-between;">
                                <span v-if="tool.language === 'shell'">
                                    <img src="/src/assets/Icon-Shell.png"/>
                                </span>
                                <span v-else>
                                    <img src="/src/assets/Icon-Python.png"/>
                                </span>
                                <span>{{ tool.name }}</span>
                            </div>
                            <div class="rowContainer" style="margin: 20px auto 0 auto;">
                                <el-button type="primary" :icon="VideoPlay" @click.stop="exec(tool.id)">执行</el-button>
                                <el-popconfirm title="确认删除?" @confirm.stop="delToolContent(tool.id)">
                                    <template #reference>
                                        <el-button type="primary" :icon="Delete" @click.stop="">删除</el-button>
                                    </template>
                                </el-popconfirm>
                            </div>
                        </el-card>
                    </div>
                </el-col>
            </el-row>
        </el-col>
    </el-row>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import { getToolTypes, getToolContents, deleteToolContent, getToolContentByName } from '@/api/tools'
import { isNull } from '@/utils/checkParam'
import { piniaTools } from "@/pinia/modules/devops"
import { Delete, VideoPlay } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const router = useRouter()
const items = reactive([])
const defaultIndex = ref(null)
const tools = reactive([])

const serachInput = ref("")

const toolTypeID = ref(0)
const toolStore = piniaTools()

const toolView = async (id, name) => {
    const res = await getToolContents({ "toolTypeID": id, "fields": "name,id,language", "name": name })
    toolTypeID.value = id
    tools.splice(0)
    if (res.code === 0 && !isNull(res.data)) {
        res.data.forEach(function (v, _) {
            tools.push(v)
        })
    }
    toolStore.setToolTypeID(id)
}

const viewData = async () => {
    const res = await getToolTypes({ "fields": "id,name" })
    if (res.code === 0 && !isNull(res.data)) {
        res.data.forEach(function (v, _) {
            items.push(v)
        })
    } else {
        return
    }
    if (toolTypeID.value === 0 && toolStore.toolTypeID === 0) {
        defaultIndex.value = items[0].id
    } else {
        defaultIndex.value = toolStore.toolTypeID
    }

    toolView(defaultIndex.value)
}
viewData()

function addTool() {
    router.push({ name: "add" })
}

function editTool(id) {
    router.push({ name: "edit", params: { id: id } })
}

function exec(id) {
    router.push({ name: "exec", params: { id: id } })
}

function serach() {
    if (isNull(serachInput.value)) {
        return
    }
    getToolContentByName({ "toolTypeID": toolTypeID.value, "fields": "name,id", "name": serachInput.value }).then(res => {
        if (res.code === 0 && !isNull(res.data)) {
            tools.splice(0)
            res.data.forEach(function (v, _) {
                tools.push(v)
            })
        } else {
            tools.splice(0)
        }
    }).catch(error => {
        ElMessage({
            message: "查询异常",
            grouping: true,
            type: "error",
        })
    })
}

function delToolContent(id) {
    deleteToolContent(id).then(res => {
        if (res.code === 0) {
            const index = tools.findIndex(tool => tool.id === id)
            if (index !== -1) {
                tools.splice(index, 1)
            }
        } else {
            ElMessage({
                message: "删除失败",
                grouping: true,
                type: "error",
            })
        }
    }).catch(error => {
        ElMessage({
            message: "删除失败",
            grouping: true,
            type: "error",
        })
    })
}

</script>

<style scoped>
.toolWidth {
    width: 215px
}

.rowContainer {
    display: flex;
    justify-content: space-between;
    margin: 0 auto;
}

.el-menu {
    border-right: none
}
</style>