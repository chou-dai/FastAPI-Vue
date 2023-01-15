import { Memory, MemoryResponse } from "@/client";
import { memoryApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const useMyMemoriesStore = defineStore("myMemories", {

    state: (): MemoryResponse => ({
        memories: []
    }),

    getters: {},

    actions: {
        async fetchMyMemories() {
            const res = await memoryApi.memoriesGet();
            this.memories = res.data.memories;
        }
    }
});