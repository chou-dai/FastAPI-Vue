import { authApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const useIsLoadingStore = defineStore("isLoading", {

    state: (): {isLoading: boolean} => ({
        isLoading: false
    }),

    getters: {},

    actions: {
        startLoading() {
            this.isLoading = true;
        },
        endLoading() {
            this.isLoading = false;
        }
    }
});