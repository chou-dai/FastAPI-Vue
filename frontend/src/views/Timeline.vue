<script lang="ts">
import { Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';

export default class Timeline extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
}
</script>

<template>
    <div class="container">
        <div v-for="memory in myMemoriesStore.memories" :key="memory.id">
            <div>●:{{ memory.title }}</div>
            <div>○:{{ memory.description }}</div>
        </div>
    </div>
</template>