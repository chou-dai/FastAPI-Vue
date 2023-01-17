import { PublicMemoryResponse } from "@/client";
import { memoryApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const usePublicMemoriesStore = defineStore("publicMemories", {
    state: (): PublicMemoryResponse => ({
        memories: []
    }),

    getters: {},

    actions: {
        async fetchPublicMemories() {
            const res = await memoryApi.apiMemoriesGet();
            this.memories = res.data.memories;
        }
    }
});