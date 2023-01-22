<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { useMyMemoriesStore } from '@/store/myMemoriesStore';
import { MemoryModal, MyMemoryListItem, BaseButton } from "../components";
import { MemoryRequest, MyMemory as Memory } from '@/client';
import { memoryApi } from '@/client/clientWrapper';
import { useOpenModalStore } from '@/store/openModalStore';
import { useIsAuthStore } from '@/store/isAuthStore';
import { ElMessage } from 'element-plus';
import { authErrorMessage, serverErrorMessage } from '../util/errorHandler';
import { memoriesFetchApi } from '@/util/fetchApiWrapper';
import { useIsLoadingStore } from '@/store/isLoadingStore';

@Options({
    components: {
        MemoryModal,
        MyMemoryListItem,
        BaseButton
    }
})
export default class MyPage extends Vue {
    public myMemoriesStore = useMyMemoriesStore();
    public openModalStore = useOpenModalStore();
    public isAuthStore = useIsAuthStore();
    public isLoadingStore = useIsLoadingStore();
    private initialState: MemoryRequest = {
        title: "",
        description: "",
        isPublic: false
    };
    public inputState = Object.assign({}, this.initialState);

    // 処理成功時のメッセージ
    private successMessage() {
        ElMessage({
            message: '処理が正常に行われました。',
            type: 'success',
        })
    }
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
        await memoryApi.apiAuthMemoriesPost(this.inputState)
            .then(() => {
                this.successMessage();
                location.reload();
            })
            .catch((err) => {
                if(err.response.status == 401) authErrorMessage();
                else serverErrorMessage();
            });
    }
    // メモリー更新APIをPUT
    public async submitUpdateMemory() {
        this.openModalStore.closeUpdateMemoryModal()
        await memoryApi.apiAuthMemoriesIdPut(this.inputState.id!, this.inputState)
            .then(async() => {
                this.successMessage();
                memoriesFetchApi();
            })
            .catch((err) => {
                if(err.response.status == 401) authErrorMessage();
                else serverErrorMessage();
            });
    }
    // メモリー削除APIをDELETE
    public async deleteMemory(memory: Memory) {
        if(confirm(memory.title+"を本当に削除しますか？")) {
            await memoryApi.apiAuthMemoriesIdDelete(memory.id)
                .then(async() => {
                    this.successMessage();
                    memoriesFetchApi();
                })
                .catch((err) => {
                    if(err.response.status == 401) authErrorMessage();
                    else serverErrorMessage();
                });
        }
    }
}
</script>

<template>
    <div class="container my-page">
        <div v-if="isAuthStore.isAuth">
            <div class="title-button-wrapper">
                <div class="title-wrapper">
                    <h2 class="my-page-title">My記録</h2>
                    <p>記録を作成した時点からカウントアップが始まります。</p>
                </div>
                <BaseButton
                    title="新規登録"
                    :onClick="openCreateModal"
                    size="large"
                    class="new-memory-button"
                />
            </div>
            <div v-if="myMemoriesStore.memories.length > 0" class="memories-list">
                <div class="list-item"
                    v-for="memory in myMemoriesStore.memories"
                    :key="memory.id"
                >
                    <MyMemoryListItem
                        :memory="memory"
                        :onClickEdit="() => openUpdateModal(memory)"
                        :onClickDelete="() => deleteMemory(memory)"
                    />
                </div>
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
                title="記録の編集"
                :inputState="inputState"
                :onClose="openModalStore.closeUpdateMemoryModal"
                :onSubmit="submitUpdateMemory"
            />
        </div>
        <el-empty
            v-if="!isAuthStore.isAuth && !isLoadingStore.isLoading"
            :image-size="300"
            description="こちらのページをご利用いただくためにはログインが必要です。"
        />
    </div>
</template>