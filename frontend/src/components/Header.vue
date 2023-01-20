<script lang="ts">
import { UserRequest } from "@/client";
import { userApi } from "@/client/clientWrapper";
import { useOpenModalStore } from "@/store/openModalStore";
import { Options, Vue } from "vue-class-component";
import UserModal from "../components/UserModal.vue";
import Navigation from "../components/Navigation.vue";
import { useIsAuthStore } from '@/store/isAuthStore';
import { useLoginUserStore } from "@/store/loginUserStore";
import { allFetchApi } from "../util/fetchApiWrapper";
import { ElMessage } from 'element-plus';

@Options({
    components: {
        UserModal,
        Navigation
    }
})
export default class Header extends Vue {
    public openModalStore = useOpenModalStore();
    public isAuthStore = useIsAuthStore();
    public isInputError = false;
    private initialState: UserRequest = {
        name: "",
        password: ""
    };
    public inputState = Object.assign({}, this.initialState);
    public loginUserInfo = useLoginUserStore();

    private successMessage() {
        ElMessage({
            message: '処理が正常に行われました。',
            type: 'success',
        })
    }
    // サインアップモーダルを開く
    public openSignupModal() {
        this.openModalStore.openSignupModal();
        this.isInputError = false;
        this.inputState = Object.assign({}, this.initialState);
    }
    // ログインモーダルを開く
    public openLoginModal() {
        this.openModalStore.openLoginModal();
        this.isInputError = false;
        this.inputState = Object.assign({}, this.initialState);
    }
    // ユーザー名変更モーダルを開く
    public openUpdateUserNameModal() {
        this.openModalStore.openUpdateUserModal();
        this.isInputError = false;
        this.inputState = Object.assign({}, this.initialState);
    }
    // サインアップAPIを送信
    public async submitSignup() {
        await userApi.apiSignupPost(this.inputState)
            .then(res => {
                if (res.status === 200) {
                    this.openModalStore.closeSignupModal()
                    location.reload();
                }
            })
            .catch(() => {
                this.isInputError = true;
            });
    }
    // ログインAPIを送信
    public async submitLogin() {
        await userApi.apiLoginPost(this.inputState)
            .then(res => {
                if (res.status === 200) {
                    this.openModalStore.closeLoginModal()
                    location.reload();
                }
            })
            .catch(() => {
                this.isInputError = true;
            });
    }
    // ユーザー名変更APIを送信
    public async submitUpdateUser() {
        await userApi.apiAuthUsersPut(this.inputState)
            .then(async(res) => {
                if (res.status === 200) {
                    this.successMessage();
                    this.openModalStore.closeUpdateUserModal()
                    await allFetchApi()
                }
            })
            .catch(() => {
                this.isInputError = true;
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
            <el-button
                v-if="!isAuthStore.isAuth"
                type="primary"
                @click="openLoginModal"
            >ログイン</el-button>
            <el-button
                v-if="!isAuthStore.isAuth"
                type="primary"
                @click="openSignupModal"
            >サインアップ</el-button>
            <el-button
                v-if="isAuthStore.isAuth"
                type="primary"
                @click="openUpdateUserNameModal"
            >ユーザー名変更</el-button>
            <el-button
                v-if="isAuthStore.isAuth"
                type="primary"
                @click="submitLogout"
            >ログアウト</el-button>
        </div>
        
    </header>
    <UserModal
        v-if="openModalStore.isOpenedSignupModal"
        title="サインアップ"
        :inputState="inputState"
        :onClose="openModalStore.closeSignupModal"
        :onSubmit="submitSignup"
        :isError="isInputError"
        errorMessage="このユーザー名はすでに使用されています。"
    />
    <UserModal
        v-if="openModalStore.isOpenedLoginModal"
        title="ログイン"
        :inputState="inputState"
        :onClose="openModalStore.closeLoginModal"
        :onSubmit="submitLogin"
        :isError="isInputError"
        errorMessage="ユーザー名もしくはパスワードに誤りがあります。"
    />
    <UserModal
        v-if="openModalStore.isOpenedUpdateUserModal"
        title="ユーザー名変更"
        :inputState="inputState"
        :onClose="openModalStore.closeUpdateUserModal"
        :onSubmit="submitUpdateUser"
        :isError="isInputError"
        errorMessage="パスワードに誤りがあります。"
    />
</template>