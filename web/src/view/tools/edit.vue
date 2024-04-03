<template>
    <div class="gva-table-box">
        <Code :propsDisabled="false" v-model:language="language" v-model:name="name" v-model:toolTypeID="toolTypeID"
            v-model:toolTypeValue="toolTypeValue" v-model:timeout="timeout" v-model:description="description"
            v-model:code="code" :toolTypes="allToolTypes">
        </Code>
        <div style="margin-top: 18px;">
            <el-form>
                <el-form-item>
                    <el-button type="primary" @click="saveData">保存</el-button>
                    <el-button @click="resetData">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>


</template>

<script setup>
import { ref } from 'vue'
import Code from '@/components/tools/code.vue'
import { ElMessage } from 'element-plus'
import { isNull } from '@/utils/checkParam'
import { useRoute, useRouter } from 'vue-router'
import { getToolContent, putToolContent, getToolTypes, postToolType } from '@/api/tools'

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

function resetData() {
    language.value = "python"
    name.value = null
    timeout.value = 3
    description.value = null
    code.value = null
}

async function saveData() {
    var flag = true
    allToolTypes.value.forEach((item) => {
        if (item.value === toolTypeValue.value) {
            flag = false
        }
    })
    let res
    if (flag) {
        res = await postToolType({ "name": toolTypeValue.value })
        if (res.code === 0) {
            toolTypeID.value = res.data.id
        } else {
            alertMsg("error", "创建脚本分类失败")
            return
        }
    }
    let data = {
        "name": name.value,
        "toolTypeID": toolTypeID.value,
        "language": language.value,
        "description": description.value,
        "code": code.value
    }

    res = await putToolContent(id, data)
    if (res.code === 0) {
        router.push({ name: "tool" })
        alertMsg("success", "修改成功")
    } else {
        alertMsg("error", res.msg)
    }
}

const alertMsg = (level, value) => {
    ElMessage({
        message: value,
        grouping: true,
        type: level,
    })
}
</script>