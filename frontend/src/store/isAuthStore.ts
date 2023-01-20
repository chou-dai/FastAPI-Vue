import { authApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";
import { authUserFetchApi } from "@/util/fetchApiWrapper";


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
                    await authUserFetchApi();
                })
                .catch(() => {
                    console.log("Not Authenticated User");
                    this.isAuth = false;
                })
        }
    }
});