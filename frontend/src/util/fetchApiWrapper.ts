import { useLoginUserStore } from "@/store/loginUserStore";
import { useMyMemoriesStore } from "@/store/myMemoriesStore"
import { usePublicMemoriesStore } from "@/store/publicMemoriesStore";

export const authUserFetchApi = async() => {
    await useLoginUserStore().fetchLoginUserStore();
    await useMyMemoriesStore().fetchMyMemories();
};

export const allFetchApi = async() => {
    await useLoginUserStore().fetchLoginUserStore();
    await useMyMemoriesStore().fetchMyMemories();
    await usePublicMemoriesStore().fetchPublicMemories();
};

export const memoriesFetchApi = async() => {
    await useMyMemoriesStore().fetchMyMemories();
    await usePublicMemoriesStore().fetchPublicMemories();
}