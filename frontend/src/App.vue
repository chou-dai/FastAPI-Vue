<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { Header } from './components';
import { usePublicMemoriesStore } from './store/publicMemoriesStore';
import { useIsAuthStore } from './store/isAuthStore';
import { ref } from 'vue';
import { useIsLoadingStore } from './store/isLoadingStore';

@Options({
    components: {
        Header
    }
})
export default class App extends Vue {
    public isLoadingStore = useIsLoadingStore();
    // アプリ起動時のバックエンドfetch
    async created() {
        this.isLoadingStore.startLoading();
        await useIsAuthStore().fetchIsAuth();
        await usePublicMemoriesStore().fetchPublicMemories();
        this.isLoadingStore.endLoading();
    }
    public loading = ref(true)
}
</script>

<template>
    <Header/>
    <main>
        <router-view/>
    </main>
</template>