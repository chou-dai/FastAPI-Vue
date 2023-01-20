import { LoginUserResponse, MyMemoryResponse } from "@/client";
import { userApi } from "@/client/clientWrapper";
import { defineStore } from "pinia";


export const useLoginUserStore = defineStore("loginUser", {

    state: (): LoginUserResponse => ({
        id: 0,
        name: "",
    }),

    getters: {},

    actions: {
        async fetchLoginUserStore() {
            const res = await userApi.apiAuthUsersGet();
            this.id = res.data.id;
            this.name = res.data.name;
            
        }
    }
});