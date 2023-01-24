<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { UserRequest } from '@/client';

@Options({
    props: {
        title: String,
        inputState: Object,
        onClose: Function,
        onSubmit: Function,
        isError: Boolean,
        errorMessage: String
    },
})

export default class UserModal extends Vue {
    public title!: string;
    public inputState!: UserRequest;
    public onClose!: () => void;
    public onSubmit!: () => void;
    public isError!: boolean;
    public errorMessage!: string;
}
</script>


<template>
    <div class="modal-outer">
        <div class="modal-backdrop" @click="onClose" />
        <div class="modal-inner">
            <h3 class="text-center">{{ title }}</h3>
            <div class="error-message">
                <span v-if="isError">{{ errorMessage }}</span>
            </div>
            <el-input
                v-model="inputState.name"
                placeholder="ユーザー名"
                maxlength="10"
                show-word-limit
                type="text"
                class="m-y-small"
            />
            <el-input
                v-model="inputState.password"
                placeholder="パスワード"
                maxlength="30"
                type="url"
                class="m-y-small"
            />
            <el-button
                class="m-t-middle"
                :disabled="inputState.password === '' || inputState.name === ''"
                @click="onSubmit"
                type="primary"
            >送信</el-button>
        </div>
    </div>
</template>