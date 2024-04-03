import service from '@/utils/request'

// 脚本分类

// 查看脚本分类
export const getToolTypes = (data) => {
    return service({
        url: '/devops/tool/toolTypes',
        method: 'get',
        params: data
    })
}

// 删除脚本分类byid
export const deleteToolType = () => {
    return service({
        url: '/devops/tool/toolType',
        method: 'delete',
    })
}

// 更新脚本分类
export const putToolType = (data) => {
    return service({
        url: '/devops/tool/toolType',
        method: 'put',
        data: data
    })
}

// 新增脚本分类
export const postToolType = (data) => {
    return service({
        url: '/devops/tool/toolType',
        method: 'post',
        data: data
    })
}

// 查看脚本分类by id
export const getToolType = (data) => {
    return service({
        url: '/devops/tool/toolType',
        method: 'get',
        params: data
    })
}


// 脚本内容
export const getToolContents = (data) => {
    return service({
        url: '/devops/tool/toolContent',
        method: 'get',
        params: data
    })
}

// 新增加脚本内容
export const postToolContent = (data) => {
    return service({
        url: '/devops/tool/toolContent',
        method: 'post',
        data: data
    })
}

// 获取脚本内容
export const getToolContent = (id) => {
    return service({
        url: '/devops/tool/toolContent/'+id,
        method: 'get',
    })
}

// 获取脚本内容by name
export const getToolContentByName = (data) => {
    return service({
        url: '/devops/tool/toolContent/fuzzy',
        method: 'get',
        params: data
    })
}


// 更新脚本内容
export const putToolContent = (id, data) => {
    return service({
        url: '/devops/tool/toolContent/'+ id,
        method: 'put',
        data: data
    })
}

// 删除脚本内容
export const deleteToolContent = (id) => {
    return service({
        url: '/devops/tool/toolContent/' + id,
        method: 'delete'
    })
}