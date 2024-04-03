import service from '@/utils/request'

// 主机

// 查看主机
export const getHost = (data) => {
    return service({
        url: '/devops/cmdb/host',
        method: 'get',
        params: data
    })
}

// 查看主机by id
export const getHostByID = (id, data) => {
    return service({
        url: '/devops/cmdb/host/' + id,
        method: 'get',
        params: data
    })
}


// 删除主机
export const deleteHost = (id) => {
    return service({
        url: '/devops/cmdb/host/' + id,
        method: 'delete',
    })
}

// 更新主机
export const putHost = (id, data) => {
    return service({
        url: '/devops/cmdb/host/' + id,
        method: 'put',
        data: data
    })
}

// 新增脚本分类
export const postHost = (data) => {
    return service({
        url: '/devops/cmdb/host',
        method: 'post',
        data: data
    })
}

