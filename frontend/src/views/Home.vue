<script lang="ts">
import { Vue, Options } from 'vue-class-component';
import { usePublicMemoriesStore } from '@/store/publicMemoriesStore';
import { PublicMemoryListItem } from '@/components';

@Options({
    components: {
        PublicMemoryListItem
    }
})
export default class Home extends Vue {
    public publicMemoriesStore = usePublicMemoriesStore();
}
</script>

<template>
    <div class="container">
        <div>
            <h1>記録カウントアップ アプリ</h1>
            <p>あれから何時間経過したかをカウントアップするアプリケーションです。</p>
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