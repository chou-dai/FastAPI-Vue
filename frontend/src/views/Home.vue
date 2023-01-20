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
        <div>
            <h1>記録カウントアップタイマー</h1>
            <p>あれからどれだけ経過したかをカウントアップするアプリケーションです。</p>
            <p v-if="!isAuthStore.isAuth">ログインをしていただくとMy記録ページをご利用いただけます。</p>
        </div>
        <h2>タイムライン</h2>
        <el-scrollbar height="400px" class="memories-list">
            <div class="list-item"
                v-for="memory in publicMemoriesStore.memories"
                :key="memory.id"
            >
                <PublicMemoryListItem :memory="memory"/>
            </div>
        </el-scrollbar>
    </div>
</template>