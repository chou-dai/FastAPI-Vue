<template>
    <header class="accent-color">
        <div class="navigation">
            logo
            <router-link to="/" class="nav-link">My記録</router-link>
            <router-link to="/timeline" class="nav-link">タイムライン</router-link>
        </div>
        <div>
            button
        </div>
    </header>
</template>

<script lang="ts">
import { Vue } from "vue-class-component";

export default class Header extends Vue {}
</script>
