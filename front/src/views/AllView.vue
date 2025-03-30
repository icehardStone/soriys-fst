<template>
    <div class="all">
        <div class="all-title">
            <h1>display all files !</h1>
        </div>
        <div class="all-content">
            <el-scrollbar class="all-content-box">
                <RouterView></RouterView>
            </el-scrollbar>
        </div>
    </div>
</template>

<style scoped>
.all {
    height: 100%;
    display: flex;
    flex-direction: column;
    max-height: 100%;
    overflow: auto;
}

.all-content {
    flex: 1;
    overflow: auto;
}

.all-content-box {
    height: 100%;
    max-height: 100%;
    overflow: auto;
}

.all-title{
    border-bottom: 1px solid rgb(246, 241, 241);
    margin-bottom: 10px;
}
</style>