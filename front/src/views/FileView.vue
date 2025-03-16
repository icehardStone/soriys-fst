<template>
    <div id="file-container">
        <div>
            <el-upload class="bg-slide-blue" :headers="headers" drag action="/api/file" :on-success="handleSuccess"
                multiple>
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                    Drop file here or <em>click to upload</em>
                </div>
                <template #tip>
                    <div class="el-upload__tip">
                        files with a size less than 500Mb
                    </div>
                </template>
            </el-upload>
        </div>
        <div id="file-container-data-grid">
            <el-table :data="tableData" :row-key="'id'" style="width: 100%">
                <el-table-column prop="name" label="Name" width="220" />
                <el-table-column prop="type" label="type" width="120" />
                <el-table-column prop="createTime" label="createTime" width="220" />
                <el-table-column prop="size" label="size" width="600" />
                <!-- <el-table-column prop="zip" label="Zip" width="120" /> -->
                <el-table-column fixed="right" label="Operations" min-width="120">
                    <template #default="scope">
                        <el-icon :size="30" class="h-pointer h-red" style="margin-right: 1rem"
                            @click="handleDelete(scope.row.id)">
                            <DeleteFilled />
                        </el-icon>
                        <el-icon :size="30" class="h-pointer h-blue"
                            @click="fileService.download(scope.row.id, scope.row.name)">
                            <Download />
                        </el-icon>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <div v-infinite-scroll="load" id="file-container-data-list" style="overflow: auto">
            <div v-for="f in tableData" :key="f.id" class="list-item">
                <el-card>
                    <div class="list-item-container">
                        <div class="list-item-container-basic">
                            <h4>{{ f.name }}</h4>
                            <p>
                                <el-text class="mx-1" type="success">{{ f.type }}</el-text>
                                <el-text class="mx-1" type="primary">{{ f.createTime }}</el-text>
                            </p>
                        </div>
                        <div class="list-item-container-operate">
                            <el-icon :size="30" class="h-pointer h-red" style="margin-right: 1rem"
                                @click="handleDelete(f.id)">
                                <DeleteFilled />
                            </el-icon>
                            <el-icon :size="30" class="h-pointer h-blue" @click="fileService.download(f.id, f.name)">
                                <Download />
                            </el-icon>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
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

<style lang="less">
.h-pointer:hover {
    cursor: pointer;
}

.h-blue {
    color: rgb(0, 81, 255) !important;
}

.h-red {
    color: red !important;
}

#file-container {
    padding: 20px;
    overflow: auto;

    &-data-list {
        .list-item {
            margin-bottom: 20px;

            &-container {
                display: flex;

                &-basic {
                    flex: 1;
                }

                &-operate {
                    display: flex;
                    align-content: center;
                    justify-content: center;
                }
            }
        }
    }
}

.mx-1 {
    margin-right: 20px !important;
}

@media (min-width:760px) {
    #file-container {
        &-data-grid {
            display: block;
        }

        &-data-list {
            display: none;
        }
    }
}

@media (max-width:760px) {
    #file-container {
        &-data-grid {
            display: none;
        }

        &-data-list {
            display: block;
        }
    }
}
</style>
