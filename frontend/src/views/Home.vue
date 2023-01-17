<script lang="ts">
import { Vue } from 'vue-class-component';
import { usePublicMemoriesStore } from '@/store/publicMemoriesStore';

export default class Home extends Vue {
    public publicMemoriesStore = usePublicMemoriesStore();
}
</script>

<template>
    <div class="container">
        <div v-for="memory in publicMemoriesStore.memories" :key="memory.id">
            <div>●:{{ memory.title }}</div>
            <div>○:{{ memory.description }}</div>
            <div>○:{{ memory.userName }}</div>
        </div>
    </div>
</template>