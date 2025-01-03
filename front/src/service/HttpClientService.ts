import axios from "axios";
import { defaultTokenService } from "./TokenService";

const httpClient = axios.create()

httpClient.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    config.headers.Authorization = defaultTokenService.get()
    // debugger;
    return config;
}, function (error) {
    return Promise.reject(error);
});

// 添加响应拦截器
httpClient.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
}, function (error) {

    // debugger
    // 超出 2xx 范围的状态码都会触发该函数。
    if (error.response.status == 401) {
        defaultTokenService.remove()
        window.location.href = "/"
    }
    return Promise.reject(error);
});

export {
    httpClient
}