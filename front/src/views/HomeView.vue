<template>
  <el-container class="layout-container-demo" style="height: 100%">
    <el-aside width="200px">
      <h1 id="soriys-title">File Tranfer</h1>
      <el-menu id="soriys-menu" :router="true" :default-active="'/home/all'">
        <el-menu-item index="/home/all"> <el-icon>
            <Folder />
          </el-icon>
          <span>ALL Files</span></el-menu-item>
        <el-menu-item index="/home/recents"><el-icon>
            <Clock />
          </el-icon>
          <span>Recents</span></el-menu-item>
        <el-menu-item index="/home/trash"><el-icon>
            <Collection />
          </el-icon>
          <span>Trash</span></el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <HomeHeader></HomeHeader>
      </el-header>
      <el-main>
        <RouterView></RouterView>
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
// import { ref } from 'vue'
// import { Menu as IconMenu, Message, Setting, Search } from '@element-plus/icons-vue'
import HomeHeader from '@/components/HomeHeader.vue'

</script>

<style scoped>
.layout-container-demo .el-header {
  position: relative;
  /* background-color: var(--el-color-primary); */
  /* color: var(--el-color-white); */
}

.layout-container-demo .el-aside {
  color: var(--el-color-white);
  background: var(--el-color-primary);
}

.layout-container-demo .el-menu {
  border-right: none;
}

.layout-container-demo .el-main {
  padding: 0 20px 0 20px;
}

.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}

#soriys-title {
  /* text-align: center; */
  padding-left: var(--el-menu-base-level-padding);
  font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif
}

.el-menu-item.is-active {
  background-color: var(--el-menu-hover-bg-color);
}

.el-menu-item:hover {
  background-color: var(--el-menu-hover-bg-color);
  color: var(--el-color-primary);
}

.el-menu el-menu--vertical {
  --el-menu-vertical-height: 100px;
}

#soriys-menu {
  --el-menu-bg-color: transparent;
  --el-menu-text-color: var(--el-color-white);
}

</style>