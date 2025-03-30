<template>
    <div>
        <div class="all-box">
        <!-- <div class="all-box-item" v-for="i in 10" :key="i">{{ i }}</div> -->
        <div v-for="f in tableData" :key="f.id" class="all-box-split">
            <div class="all-box-item">
                <!-- <h5>{{ f.name }}</h5> -->
                <div class="all-box-item-row"> <el-text class="mx-1" size="large" truncated>{{ f.name }}</el-text></div>
                <div class="all-box-item-row">
                    <el-text class="mx-1" type="success">{{ f.type?.trimStart() }}</el-text>

                </div>
                <div class="all-box-item-row">
                    <el-text class="mx-1" type="primary">{{ f.createTime }}</el-text>
                </div>
                <div style="flex: 1;"></div>
                <div class="all-box-item-row" style="text-align: right;">
                    <el-icon :size="30" class="h-pointer h-red" style="margin-right: 1rem" @click="handleDelete(f.id)">
                        <DeleteFilled />
                    </el-icon>
                    <el-icon :size="30" class="h-pointer h-blue" @click="fileService.download(f.id, f.name)">
                        <Download />
                    </el-icon>
                </div>
            </div>
        </div>
    </div>
        <div style="margin: 30px 20px 20px 20px;">
            <el-upload style="height: 100%;" class="bg-slide-blue" :headers="headers" drag action="/api/file" :on-success="handleSuccess"
                multiple>
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                    Drop file here or <em>click to upload</em>
                </div>
                <!-- <template #tip>
                    <div class="el-upload__tip">
                        files with a size less than 500Mb
                    </div>
                </template> -->
            </el-upload>
        </div>
    </div>
    
</template>

<script lang="ts" setup>

const items = [
    1, 2, 3, 4, 5
]

import { Download, UploadFilled } from '@element-plus/icons-vue'
import { onMounted, ref } from 'vue'
import { FileModel, fileService } from '@/service/FileService'
import { defaultTokenService } from '@/service/TokenService'

onMounted(() => {
    refresh()
})

const handleDelete = async (id: number) => {

    await fileService.remove(id)
    await refresh()
}

const handleSuccess = async (res: any) => {
    console.log("handleSuccess")
    console.log(res)

    await refresh()
}
const load = () => {

}
const refresh = async (pageSize?: number, pageIndex?: number) => {
    tableData.value = await fileService.getList(pageSize, pageIndex)
}

const tableData = ref<FileModel[]>([])
const headers = {
    "Authorization": defaultTokenService.get()
}
</script>

<style scoped>
.all-box {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
}
.all-box-split{
    width: 100px;
    margin: 10px 0px;
    padding: 20px;
    flex: 1 1 calc(20% - 20px);
    box-sizing: border-box;
    height: 200px;
}
.all-box-item {
    
    box-shadow: var(--el-box-shadow-light);
    padding: 20px;
    height: 100%;
    display: flex;
    flex-direction: column;
}



.h-pointer:hover {
    cursor: pointer;
}

.h-blue {
    color: rgb(0, 81, 255) !important;
}

.h-red {
    color: red !important;
}

.all-box-item-row {
    margin-bottom: 5px;
}
</style>