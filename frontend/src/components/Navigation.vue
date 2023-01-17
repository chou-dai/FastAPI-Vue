<script lang="ts">
import { Vue } from "vue-class-component";
import { ref } from "vue";

export default class Header extends Vue {
    public activeIndex = ref(location.pathname === "/" ? "home" : "mypage")
}
</script>

<template>
    <div style="display:flex">
        <div>logo</div>
        <el-menu
            :default-active="activeIndex"
            background-color="#694F94"
            text-color="#fff"
            active-text-color="#fff"
            mode="horizontal"
            :ellipsis="false"
        >
            <el-menu-item index="home">
                <router-link to="/">ホーム</router-link>
            </el-menu-item>
            <el-menu-item index="mypage">
                <router-link to="/mypage">My記録</router-link>
            </el-menu-item>
        </el-menu>
    </div>
</template>