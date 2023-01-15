<script lang="ts">
import { Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemories';

export default class MyMemory extends Vue {
    public myMemories = useMyMemoriesStore();
}
</script>

<template>
    <div class="container">
        <div v-for="memory in myMemories.memories" :key="memory.id">
            <div>{{memory.description}}</div>
            <div>{{memory.title}}</div>
        </div>
    </div>
</template>