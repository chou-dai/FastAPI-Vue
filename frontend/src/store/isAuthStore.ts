import { authApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";
import { useMyMemoriesStore } from "./myMemoriesStore"


export const useIsAuthStore = defineStore("isAuth", {

    state: (): {isAuth: boolean} => ({
        isAuth: false
    }),

    getters: {},

    actions: {
        async fetchIsAuth() {
            await authApi.apiAuthGet()
                .then(async(res) => {
                    console.log("Authenticated User");
                    this.isAuth = (res.status === 200);
                    // MyMemoriesをfetch
                    await useMyMemoriesStore().fetchMyMemories();
                })
                .catch(() => {
                    console.log("Not Authenticated User");
                    this.isAuth = false;
                })
        }
    }
});