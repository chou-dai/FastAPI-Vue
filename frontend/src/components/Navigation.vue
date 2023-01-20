<script lang="ts">
import { Vue } from "vue-class-component";
import { ref } from "vue";

export default class Header extends Vue {
    public currentPath = ref(location.pathname === "/" ? "home" : "mypage")
}
</script>

<template>
    <div class="navigation">
        <div>logo</div>
        <el-menu
            mode="horizontal"
            :default-active="currentPath"
            :ellipsis="false"
        >
            <el-menu-item index="home">
                <router-link to="/" class="nav-link">ホーム</router-link>
            </el-menu-item>
            <el-menu-item index="mypage">
                <router-link to="/mypage" class="nav-link">My記録</router-link>
            </el-menu-item>
        </el-menu>
    </div>
</template>