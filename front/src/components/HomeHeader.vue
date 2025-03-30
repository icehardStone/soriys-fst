<template>
    <div id="soriys-toolbar">
        <div id="soriys-toolbar-search">
            <el-input v-model="searchVal" style="width: 240px" placeholder="Please input" />
        </div>
        <ul id="soriys-toolbar-menus">
            <!-- <div></div> -->
            <li>
                <el-avatar :size="20" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
            </li>
            <li>
                <el-badge :value="3" class="item">
                    <el-icon>
                        <Bell />
                    </el-icon>
                </el-badge>
            </li>
            <li>
                <el-icon>
                    <QuestionFilled />
                </el-icon>
            </li>
        </ul>
    </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
// import { Menu as IconMenu, Message, Setting, Search } from '@element-plus/icons-vue'

const searchVal = ref('')

</script>

<style scoped>
#soriys-toolbar {
    display: flex;
    height: 100%;
    align-items: center;
}

#soriys-toolbar-menus {
    flex: 1;

    display: flex;
    flex-direction: row-reverse;
    align-items: center;

    list-style: none;
}

#soriys-toolbar-menus  li {
    margin-right: 20px;
}

.el-badge {
    /* margin-top: 10px; */
    text-align: center;
    vertical-align: middle;
}
</style>