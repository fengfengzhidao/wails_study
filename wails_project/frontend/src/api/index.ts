import axios from "axios";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import type {Ref} from "vue";
import {ref} from "vue";

export const useAxios = axios.create({
    baseURL: "",
    timeout: 6000,
})

export interface baseResponse<T> {
    code: number
    msg: string
    data: T
}

useAxios.interceptors.request.use((config) => {
    const store = useStore()
    // config.headers.set("token", store.userInfo.token)
    return config
})

useAxios.interceptors.response.use((res) => {
    return res.data
}, (error) => {
    Message.error(error.message)
    console.log(error)
})

export interface baseParams {
    page?: number
    limit?: number
    sort?: string
    key?: string
}

export interface baseListResponse<T> {
    count: number
    list: T[]
}

export interface optionsType {
    label: string
    value: string | number
}

export function defaultDeleteApi(url: string, idList: number[]): Promise<baseResponse<string>> {
    return useAxios.delete(url, {data: {idList}})
}

export function getOptions(ref: Ref<optionsType[]>, func: () => Promise<baseResponse<optionsType[]>>) {
    func().then((res) => {
        ref.value = res.data
    })
}


export function genOptions(func: () => Promise<baseResponse<optionsType[]>>): Ref<optionsType[]> {
    const r = ref<optionsType[]>([])
    func().then((res) => {
        r.value = res.data
    })
    return r
}

interface cacheType {
    time: number
    res: Promise<baseResponse<optionsType[]>>
}

const  cacheData:Record<string, cacheType> = {}



export function genOptionsCache(func: () => Promise<baseResponse<optionsType[]>>): Ref<optionsType[]> {
    // 判断有没有请求这个地址
    const key = func.toString()
    let val = cacheData[key]
    if (!val){
        const res = func()
        val = {
            time: new Date().getTime(),
            res: res,
        }
        cacheData[key] = val
    }
    const nowTime = new Date().getTime()
    if (nowTime - val.time > 5000){
        const res = func()
        val = {
            time: new Date().getTime(),
            res: res,
        }
        cacheData[key] = val
    }
    const r = ref<optionsType[]>([])
    val.res.then((res) => {
        r.value = res.data
    })
    return r
}