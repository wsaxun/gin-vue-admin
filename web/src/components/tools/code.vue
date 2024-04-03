<script setup>
import { ref, computed, watch, onMounted} from 'vue'
import { Codemirror } from 'vue-codemirror'
import { StreamLanguage } from "@codemirror/language"
import { shell } from "@codemirror/legacy-modes/mode/shell"
import { python } from "@codemirror/legacy-modes/mode/python"
import { oneDark } from '@codemirror/theme-one-dark'

const props = defineProps({
    propsDisabled: { type: Boolean, default: true },
    toolTypes: {type: Array}
})


const language = defineModel("language")
const name = defineModel("name")
const timeout = defineModel("timeout")
const description = defineModel("description")
const code = defineModel("code")
const toolTypeID = defineModel("toolTypeID")
const toolTypeValue = defineModel("toolTypeValue")

const lang = computed(() => language.value === 'shell' ? shell : python)
const extensions = ref([StreamLanguage.define(lang.value), oneDark])
const style = { height: '400px', width: '100%' }

function changeLang() {
    extensions.value = [StreamLanguage.define(lang.value), oneDark]
}


const restaurants = ref([])
const querySearch = (queryString, cb) => {
    const results = queryString ? restaurants.value.filter(createFilter(queryString)) : restaurants.value
    cb(results)
}
const createFilter = (queryString) => {
    return (restaurant) => {
        return (
            restaurant.value.indexOf(queryString) === 0
        )
    }
}

const loadAll = () => {
    let results = []
    props.toolTypes.forEach(function (v, _) {
        if (v.id === toolTypeID.value){
            toolTypeValue.value = v.value
        }
        results.push(v)
    })
    restaurants.value = results
}

watch(props.toolTypes, (value) => {
    if (value.length > 0) {
        loadAll()
    }
})
const handleSelect = (item) => {
    toolTypeID.value = item.id
}
</script>

<template>
    <el-form label-position="left" label-width="auto" :disabled="props.propsDisabled">
        <el-form-item label="脚本名" required style="max-width: 600px">
            <el-input v-model="name" />
        </el-form-item>
        <el-form-item label="语言类型">
            <el-radio-group v-model="language" @change="changeLang">
                <el-radio label="python" />
                <el-radio label="shell" />
            </el-radio-group>
        </el-form-item>
        <el-form-item label="执行超时(分)">
            <el-input-number v-model="timeout" />
        </el-form-item>
        <el-form-item label="分类" required style="max-width: 600px">
            <el-autocomplete v-model="toolTypeValue" :fetch-suggestions="querySearch" clearable
                placeholder="Please Input" @select="handleSelect" />
        </el-form-item>
        <el-form-item label="描述" style="max-width: 600px">
            <el-input v-model="description" />
        </el-form-item>
        <el-form-item>
            <Codemirror :disabled="props.propsDisabled" v-model="code" placeholder="暂无数据..." :style="style"
                autofocus="false" indent-with-tab="true" tabSize="4" fullScreen="true" :extensions="extensions">
            </Codemirror>
        </el-form-item>
    </el-form>
</template>