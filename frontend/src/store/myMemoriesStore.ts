import { MyMemoryResponse } from "@/client";
import { memoryApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const useMyMemoriesStore = defineStore("myMemories", {

    state: (): MyMemoryResponse => ({
        memories: []
    }),

    getters: {},

    actions: {
        async fetchMyMemories() {
            const res = await memoryApi.apiAuthMemoriesGet();
            this.memories = res.data.memories;
        }
    }
});