<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import Header from './components/Header.vue';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import { usePublicMemoriesStore } from './store/publicMemoriesStore';
import { useIsAuthStore } from './store/isAuthStore';

@Options({
    components: {
        Header,
    },
})
export default class App extends Vue {
    // 初回レンダリング時のバックエンドfetch
    async created() {
        await useIsAuthStore().fetchIsAuth();
        await usePublicMemoriesStore().fetchPublicMemories();
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