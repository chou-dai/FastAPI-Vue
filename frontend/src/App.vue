<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import Header from './components/Header.vue';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';

@Options({
    components: {
        Header,
    },
})
export default class App extends Vue {
    // 初回レンダリング時のバックエンドへのfetch
    async created() {
        await useMyMemoriesStore().fetchMyMemories();
    }
}
</script>

<template>
    <Header/>
    <main>
        <router-view/>
    </main>
</template>