import { defineStore } from 'pinia'
import { ref } from 'vue'

export const piniaHosts = defineStore("hosts",() => {
    const multipleSelection = ref([])
    const currentPageSelection = ref([])
    function handleSelectChange (row) {
        var flag = false
        var i = null
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
            })
            if (n !== null) {
                multipleSelection.value.splice(n, 1)
            }
        } else {
            currentPageSelection.value.push(row)
            multipleSelection.value.push(row)
        }
    }
    function selectCurAll(rows, apiData){
        if (rows.length === 0 ) {
            resetCurrentPage()
            apiData.forEach(function(tableVal){
                var n = null
                multipleSelection.value.forEach(function(val, index){
                    if (val.id === tableVal.id) {
                        n = index
                    }
                })
                if (n !== null ) {
                    spliceMult(n, 1)
                }
            })
            return
        }
        if (multipleSelection.value.length === 0){
            apiData.forEach(function(tableVal){
                pushCurrentPage(tableVal)
                pushMultPage(tableVal)
            })
            return
        }
    
        resetCurrentPage()
        apiData.forEach(function(tableVal){
            pushCurrentPage(tableVal)
            var n = null
            if (multipleSelection.value.length > 0 ) {
                multipleSelection.value.forEach(function(val, index){
                    if (val.id === tableVal.id) {
                        n = index
                        return
                    }
                })
            }
    
            if (n === null ) {
                pushMultPage(tableVal)
            }       
        })
    }
    function pushCurrentPage(tableItem) {
        currentPageSelection.value.push(tableItem)
    }
    function pushMultPage(tableItem) {
        multipleSelection.value.push(tableItem)
    }
    function resetCurrentPage() {
        currentPageSelection.value.splice(0)
    }
    function resetMult(){
        multipleSelection.value.splice(0)
    }
    function spliceMult(n) {
        multipleSelection.value.splice(n, 1)
    }
    return {
        multipleSelection,
        currentPageSelection,
        resetCurrentPage,
        resetMult,
        pushCurrentPage,
        pushMultPage,
        handleSelectChange,
        spliceMult,
        selectCurAll
    }
})