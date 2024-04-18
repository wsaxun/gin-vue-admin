<template>
    <div class="gva-table-box">
        <Code :propsDisabled="true" v-model:language="language" v-model:name="name" v-model:toolTypeID="toolTypeID"
            v-model:toolTypeValue="toolTypeValue" v-model:timeout="timeout" v-model:description="description"
            v-model:code="code" :toolTypes="allToolTypes">
        </Code>
        <div style="margin-top: 18px;">
            <el-form>
                <el-form-item label="参数">
                    <el-input v-model="args" clearable
                        placeholder="请输入参数，多个参数以空格分隔,参数不能带空格。如：arg1 arg2 arg3"></el-input>
                </el-form-item>
                <el-form-item label="主机" required>
                    <el-input v-model="hosts" clearable placeholder="选择or输入主机">
                        <template #prepend>
                            <el-button plain @click="search">选择主机</el-button>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="execTask">执行</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div>
            <el-dialog v-model="dialogTableVisible" title="主机列表">
                <Host ref="childRef">
                </Host>
                <template #footer>
                    <div class="dialog-footer">
                        <el-button type="primary" @click="confirm">确认</el-button>
                        <el-button @click="dialogTableVisible=false">取消</el-button>
                    </div>
                </template>
            </el-dialog>
        </div>
    </div>

</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import Code from '@/components/tools/code.vue'
import Host from '@/components/host/host.vue'
import { ElMessage } from 'element-plus'
import { isNull } from '@/utils/checkParam'
import { useRoute, useRouter } from 'vue-router'
import { getToolContent, getToolTypes, postToolType } from '@/api/tools'
import {piniaHosts} from '@/pinia/modules/hosts'

const language = ref(null)
const name = ref(null)
const timeout = ref(null)
const description = ref(null)
const code = ref(null)
const toolTypeID = ref(null)
const toolTypeValue = ref(null)
const allToolTypes = ref([])

const route = useRoute()
const router = useRouter()
const id = route.params.id

const hostStore = piniaHosts()

const hosts = ref("")
const args = ref("")

const childRef = ref()
const dialogTableVisible = ref(null)

const toolView = async (id) => {
    let res = await getToolContent(id)
    if (res.code === 0) {
        language.value = res.data.language
        name.value = res.data.name
        timeout.value = res.data.timeout
        description.value = res.data.description
        code.value = res.data.code
        toolTypeID.value = res.data.toolTypeID
    }
}
toolView(id)

const getAllToolType = () => {
    getToolTypes({ "fields": "id,name" }).then((res) => {
        if (res.code === 0) {
            res.data.forEach(function (v, _) {
                allToolTypes.value.push({ "value": v.name, "id": v.id })
            })
        }
    })
}
getAllToolType()

async function execTask() {
    if (isNull(hosts.value)) {
        alertMsg("error", "主机为空")
        return
    } 
}


const confirm = () => {
    dialogTableVisible.value=false
    if (hosts.value.length !== 0) {
        hosts.value = ""
    }
    if (!isNull(hostStore.multipleSelection)) {
        hostStore.multipleSelection.forEach((value) => {
        if (!isNull(value.hostname)){
            if (hosts.value === "") {
                hosts.value = value.hostname
            } else {
                hosts.value += `,${value.hostname}`
            }   
        }
    })
    }

}

const search = () => {
    dialogTableVisible.value = true
    if (! isNull(childRef.value)){
        childRef.value.fetchData()
    } 
}

const alertMsg = (level, value) => {
    ElMessage({
        message: value,
        grouping: true,
        type: level,
    })
}

onUnmounted(() => {
    hostStore.resetCurrentPage()
    hostStore.resetMult()
})
</script>