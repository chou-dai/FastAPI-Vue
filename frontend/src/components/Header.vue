<script lang="ts">
import { UserRequest } from "@/client";
import { userApi } from "@/client/clientWrapper";
import { useOpenModalStore } from "@/store/openModalStore";
import { Options, Vue } from "vue-class-component";
import UserModal from "../components/UserModal.vue";

@Options({
    components: {
        UserModal
    }
})
export default class Header extends Vue {
    public openModalStore = useOpenModalStore();
    private initialInputState: UserRequest = {
        name: "",
        password: ""
    };
    public inputState = this.initialInputState;

    public changeInputState(e: any) {
        this.inputState = {
            ...this.inputState,
            [e.target.name]: e.target.value
        }
    }
    public openSignupModal() {
        this.openModalStore.openSignupModal();
        this.inputState = this.initialInputState;
    }
    public openLoginModal() {
        this.openModalStore.openLoginModal();
        this.inputState = this.initialInputState;
    }
    public async submitSignup() {
        this.openModalStore.closeSignupModal()
        userApi.apiSignupPost(this.inputState);
        
    }
    public async submitLogin() {
        this.openModalStore.closeLoginModal()
        await userApi.apiLoginPost(this.inputState);
    }
    public async submitLogout() {
        await userApi.apiAuthLogoutPost();
    }
}
</script>

<template>
    <header class="accent-color">
        <div class="navigation">
            logo
            <router-link to="/" class="nav-link">My記録</router-link>
            <router-link to="/timeline" class="nav-link">タイムライン</router-link>
        </div>
        <div>
            <button @click="openLoginModal">ログイン</button>
            <button @click="openSignupModal">サインアップ</button>
            <button @click="submitLogout">ログアウト</button>
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