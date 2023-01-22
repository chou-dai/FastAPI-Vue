<script lang="ts">
import { Vue, Options } from 'vue-class-component';
import { usePublicMemoriesStore } from '@/store/publicMemoriesStore';
import { PublicMemoryListItem } from '@/components';
import { useIsAuthStore } from '@/store/isAuthStore';

@Options({
    components: {
        PublicMemoryListItem
    }
})
export default class Home extends Vue {
    public publicMemoriesStore = usePublicMemoriesStore();
    public isAuthStore = useIsAuthStore();
}
</script>

<template>
    <div class="container">
        <div class="app-title-wrapper">
            <h1 class="app-title">経過時間計測アプリ</h1>
            <p>あれからどれだけ経過したかをカウントアップするアプリケーションです。</p>
            <p v-if="!isAuthStore.isAuth">ログインをしていただくとMy記録ページをご利用いただけます。</p>
        </div>
        <div class="timeline-wrapper">
            <h2 class="timeline-title">タイムライン</h2>
            <el-scrollbar class="memories-list timeline-memories">
                <div class="list-item"
                    v-for="memory in publicMemoriesStore.memories"
                    :key="memory.id"
                >
                    <PublicMemoryListItem :memory="memory"/>
                </div>
            </el-scrollbar>
        </div>
    </div>
</template>