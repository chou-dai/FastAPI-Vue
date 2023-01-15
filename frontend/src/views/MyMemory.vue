<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import MemoryModal from "../components/MemoryModal.vue";
import { MemoriesGetRequest } from '@/client';
import { memoryApi } from '@/client/clientWrapper';

@Options({
    components: {
        MemoryModal
    }
})
export default class MyMemory extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
    public isOpenedModal: boolean = false;
    public inputState: MemoriesGetRequest = {title: ""};

    public openModal() {
        this.isOpenedModal = true;
    }
    public onClose() {
        this.isOpenedModal = false;
    }
    public changeInputState(e: any) {
        this.inputState = {
            ...this.inputState,
            [e.target.name]: e.target.value
        }
    }
    public async submitNewMemory() {
        this.onClose();
        await memoryApi.memoriesPost(this.inputState);
        this.myMemoriesStore.fetchMyMemories();
    }
}
</script>

<template>
    <div class="container">
        <button @click="openModal">button</button>
        <div v-for="memory in myMemoriesStore.memories" :key="memory.id">
            <div>●:{{memory.title}}</div>
            <div>○:{{memory.description}}</div>
        </div>
    </div>
    <MemoryModal
        v-if="isOpenedModal"
        title="aaa"
        :inputState="inputState"
        :changeInputState="changeInputState"
        :onClose="onClose"
        :onSubmit="submitNewMemory"
    />
</template>