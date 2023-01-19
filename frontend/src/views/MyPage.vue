<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import { usePublicMemoriesStore } from '@/store/publicMemoriesStore';
import { MemoryModal, MemoryListItem } from "../components";
import { MemoryRequest, MyMemory as Memory } from '@/client';
import { memoryApi } from '@/client/clientWrapper';
import { useOpenModalStore } from '@/store/openModalStore';
import { useIsAuthStore } from '@/store/isAuthStore';

@Options({
    components: {
        MemoryModal,
        MemoryListItem
    }
})
export default class MyPage extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
    public publicMemoriesStore = usePublicMemoriesStore();
    public openModalStore = useOpenModalStore();
    public isAuthStore = useIsAuthStore();
    private initialState: MemoryRequest = {
        title: "",
        description: "",
        isPublic: false
    };
    public inputState = Object.assign({}, this.initialState);

    // メモリー作成モーダルを開く
    public openCreateModal() {
        this.openModalStore.openCreateMemoryModal();
        this.inputState = Object.assign({}, this.initialState);
    }
    // メモリー更新モーダルを開く
    public openUpdateModal(memory: Memory) {
        this.openModalStore.openUpdateMemoryModal();
        this.inputState = Object.assign({},memory);
    }
    // メモリー作成APIをPOST
    public async submitCreateMemory() {
        this.openModalStore.closeCreateMemoryModal()
        await memoryApi.apiAuthMemoriesPost(this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
        await this.publicMemoriesStore.fetchPublicMemories();
    }
    // メモリー更新APIをPUT
    public async submitUpdateMemory() {
        this.openModalStore.closeUpdateMemoryModal()
        await memoryApi.apiAuthMemoriesIdPut(this.inputState.id!, this.inputState);
        await this.myMemoriesStore.fetchMyMemories();
        await this.publicMemoriesStore.fetchPublicMemories();
    }
    // メモリー削除APIをDELETE
    public async deleteMemory(memory: Memory) {
        if(confirm(memory.title+"を本当に削除しますか？")) {
            await memoryApi.apiAuthMemoriesIdDelete(memory.id);
            await this.myMemoriesStore.fetchMyMemories();
            await this.publicMemoriesStore.fetchPublicMemories();
        }
    }
}
</script>

<template>
    <div class="container">
        <div v-if="isAuthStore.isAuth">
            <el-button type="primary" @click="openCreateModal">新規作成</el-button>
            <div v-for="memory in myMemoriesStore.memories" :key="memory.id">
                <MemoryListItem
                    :memory="memory"
                    :onClickEdit="() => openUpdateModal(memory)"
                    :onClickDelete="() => deleteMemory(memory)"
                />
            </div>
            <MemoryModal
                v-if="openModalStore.isOpenedCreateMemoryModal"
                title="記録の作成"
                :inputState="inputState"
                :onClose="openModalStore.closeCreateMemoryModal"
                :onSubmit="submitCreateMemory"
            />
            <MemoryModal
                v-if="openModalStore.isOpenedUpdateMemoryModal"
                title="記録の変更"
                :inputState="inputState"
                :onClose="openModalStore.closeUpdateMemoryModal"
                :onSubmit="submitUpdateMemory"
            />
        </div>
        <el-empty
            v-if="!isAuthStore.isAuth"
            :image-size="300"
            description="こちらのページをご利用いただくためにはログインが必要です。"
        />
    </div>
</template>