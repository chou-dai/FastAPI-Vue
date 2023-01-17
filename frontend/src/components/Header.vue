<script lang="ts">
import { UserRequest } from "@/client";
import { userApi } from "@/client/clientWrapper";
import { useOpenModalStore } from "@/store/openModalStore";
import { Options, Vue } from "vue-class-component";
import UserModal from "../components/UserModal.vue";
import Navigation from "../components/Navigation.vue";
import { useIsAuthStore } from '@/store/isAuthStore';

@Options({
    components: {
        UserModal,
        Navigation
    }
})
export default class Header extends Vue {

    public openModalStore = useOpenModalStore();
    private initialInputState: UserRequest = {
        name: "",
        password: ""
    };
    public inputState = this.initialInputState;
    public isAuthStore = useIsAuthStore();

    // サインアップ・ログインモーダルの文字入力の状態更新
    public changeInputState(e: any) {
        this.inputState = {
            ...this.inputState,
            [e.target.name]: e.target.value
        }
    }
    // サインアップモーダルを開く
    public openSignupModal() {
        this.openModalStore.openSignupModal();
        this.inputState = this.initialInputState;
    }
    // ログインモーダルを開く
    public openLoginModal() {
        this.openModalStore.openLoginModal();
        this.inputState = this.initialInputState;
    }
    // サインアップAPIを送信
    public async submitSignup() {
        this.openModalStore.closeSignupModal()
        await userApi.apiSignupPost(this.inputState)
            .then(res => {
                if (res.status === 200) {
                    location.reload();
                }
            });
    }
    // ログインAPIを送信
    public async submitLogin() {
        this.openModalStore.closeLoginModal()
        await userApi.apiLoginPost(this.inputState)
            .then(res => {
                if (res.status === 200) {
                    location.reload();
                }
            });
    }
    // ログアウトAPIを送信
    public async submitLogout() {
        await userApi.apiAuthLogoutPost()
            .then(res => {
                if (res.status === 200) {
                    location.reload();
                }
            });
    }
}
</script>

<template>
    <header class="accent-color">
        <Navigation />
        <div>
            <el-button type="primary" v-if="!isAuthStore.isAuth" @click="openLoginModal">ログイン</el-button>
            <el-button type="primary" v-if="!isAuthStore.isAuth" @click="openSignupModal">サインアップ</el-button>
            <el-button type="success" v-if="isAuthStore.isAuth" @click="submitLogout">ログアウト</el-button>
        </div>
        
    </header>
    <UserModal
        v-if="openModalStore.isOpenedSignupModal"
        title="サインアップ"
        :inputState="inputState"
        :changeInputState="changeInputState"
        :onClose="openModalStore.closeSignupModal"
        :onSubmit="submitSignup"
    />
    <UserModal
        v-if="openModalStore.isOpenedLoginModal"
        title="ログイン"
        :inputState="inputState"
        :changeInputState="changeInputState"
        :onClose="openModalStore.closeLoginModal"
        :onSubmit="submitLogin"
    />
</template>