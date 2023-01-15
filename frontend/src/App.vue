<template>
    <Header/>
    <main>
        <router-view/>
    </main>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import Header from './components/Header.vue';

@Options({
    components: {
        Header,
    },
})
export default class App extends Vue {}
</script>

<style>
</style>
