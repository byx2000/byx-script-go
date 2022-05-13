import axios from "axios"

// 封装axios网络请求
export function request(config) {
    const instance = axios.create({
        baseURL: process.env.NODE_ENV === 'production'
            ? '/api/'
            : 'http://localhost:8080',
        timeout: 10000
    })

    instance.interceptors.response.use(res => {
        return res.data
    })

    return instance(config)
}

// 运行脚本
export function runScript(data) {
    return request({
        method: 'post',
        url: '/run',
        data
    })
}