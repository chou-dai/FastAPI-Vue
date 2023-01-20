<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { Header } from './components';
import { usePublicMemoriesStore } from './store/publicMemoriesStore';
import { useIsAuthStore } from './store/isAuthStore';

@Options({
    components: {
        Header
    }
})
export default class App extends Vue {
    // アプリ起動時のバックエンドfetch
    async created() {
        await useIsAuthStore().fetchIsAuth();
        await usePublicMemoriesStore().fetchPublicMemories();
    }
}
</script>

<template>
    <Header/>
    <main>
        <router-view/>
    </main>
</template>