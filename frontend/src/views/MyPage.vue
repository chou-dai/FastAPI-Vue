<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import MemoryModal from "../components/MemoryModal.vue";
import { MemoryRequest, MyMemory as Memory } from '@/client';
import { memoryApi } from '@/client/clientWrapper';
import { useOpenModalStore } from '@/store/openModalStore';
import { useIsAuthStore } from '@/store/isAuthStore';

@Options({
    components: {
        MemoryModal
    }
})
export default class MyPage extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
    public openModalStore = useOpenModalStore();
    private initialInputState: MemoryRequest = {};
    public inputState = this.initialInputState;
    public isAuthStore = useIsAuthStore();

    // モーダルの入力値の状態管理
    public changeInputState(e: any) {
        this.inputState = {
            ...this.inputState,
            [e.target.name]: e.target.value
        }
    }
    // メモリー作成モーダルを開く
    public openCreateModal() {
        this.openModalStore.openCreateMemoryModal();
        this.inputState = this.initialInputState;
    }
    // メモリー更新モーダルを開く
    public openUpdateModal(memory: Memory) {
        this.openModalStore.openUpdateMemoryModal();
        this.inputState = memory;
    }
    // メモリー作成APIをPOST
    public async submitCreateMemory() {
        this.openModalStore.closeCreateMemoryModal()
        await memoryApi.apiAuthMemoriesPost(this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
    }
    // メモリー更新APIをPUT
    public async submitUpdateMemory() {
        this.openModalStore.closeUpdateMemoryModal()
        await memoryApi.apiAuthMemoriesIdPut(this.inputState.id!, this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
    }
    // メモリー削除APIをDELETE
    public async deleteMemory(memory: Memory) {
        if(confirm(memory.title+"を本当に削除しますか？")) {
            await memoryApi.apiAuthMemoriesIdDelete(memory.id);
            await this.myMemoriesStore.fetchMyMemories();
        }
    }
}
</script>

<template>
    <div class="container">
        <div v-if="isAuthStore.isAuth">
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
        <el-empty
            v-if="!isAuthStore.isAuth"
            :image-size="300"
            description="こちらのページをご利用いただくためにはログインが必要です。" />
    </div>
</template>