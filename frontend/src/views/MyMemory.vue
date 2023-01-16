<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import MemoryModal from "../components/MemoryModal.vue";
import { MemoriesGetRequest, Memory } from '@/client';
import { memoryApi } from '@/client/clientWrapper';
import { useOpenModalStore } from '@/store/openModalStore';

@Options({
    components: {
        MemoryModal
    }
})
export default class MyMemory extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
    public openModalStore = useOpenModalStore();
    public isOpenedModal: boolean = false;
    private initialInputState: MemoriesGetRequest = {title: ""};
    public inputState = this.initialInputState;

    public changeInputState(e: any) {
        this.inputState = {
            ...this.inputState,
            [e.target.name]: e.target.value
        }
    }
    public openCreateModal() {
        this.openModalStore.openCreateMemoryModal();
        this.inputState = this.initialInputState;
    }
    public openUpdateModal(memory: Memory) {
        this.openModalStore.openUpdateMemoryModal();
        this.inputState = memory;
    }
    public async submitCreateMemory() {
        this.openModalStore.closeCreateMemoryModal()
        await memoryApi.memoriesPost(this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
    }
    public async submitUpdateMemory() {
        this.openModalStore.closeUpdateMemoryModal()
        await memoryApi.memoriesIdPut(this.inputState.id!, this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
    }
    public async deleteMemory(memory: Memory) {
        if(confirm(memory.title+"を本当に削除しますか？")) {
            await memoryApi.memoriesIdDelete(memory.id);
            await this.myMemoriesStore.fetchMyMemories();
        }
    }
}
</script>

<template>
    <div class="container">
        <button @click="openCreateModal">新規作成</button>
        <div v-for="memory in myMemoriesStore.memories" :key="memory.id">
            <div>●:{{ memory.title }}</div>
            <div>○:{{ memory.description }}</div>
            <div>
                <button @click="openUpdateModal(memory)">変更</button>
                <button @click="deleteMemory(memory)">削除</button>
            </div>
        </div>
        <MemoryModal
            v-if="openModalStore.isOpenedCreateMemoryModal"
            title="記録の作成"
            :inputState="inputState"
            :changeInputState="changeInputState"
            :onClose="openModalStore.closeCreateMemoryModal"
            :onSubmit="submitCreateMemory"
        />
        <MemoryModal
            v-if="openModalStore.isOpenedUpdateMemoryModal"
            title="記録の変更"
            :inputState="inputState"
            :changeInputState="changeInputState"
            :onClose="openModalStore.closeUpdateMemoryModal"
            :onSubmit="submitUpdateMemory"
        />
    </div>
</template>