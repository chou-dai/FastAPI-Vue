import { MyMemoryResponse } from "@/client";
import { authApi, memoryApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const useIsAuthStore = defineStore("isAuth", {

    state: (): {isAuth: boolean} => ({
        isAuth: false
    }),

    getters: {},

    actions: {
        async fetchIsAuth() {
            await authApi.apiAuthGet()
                .then(res => {
                    this.isAuth = (res.status === 200);
                })
                .catch(() => {
                    console.log("Not Authenticated User");
                    this.isAuth = false;
                })
        }
    }
});