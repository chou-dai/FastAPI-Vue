<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { MemoryRequest } from '@/client';
import BaseButton from "./BaseButton.vue";

@Options({
    components: {
        BaseButton
    },
    props: {
        title: String,
        inputState: Object,
        onClose: Function,
        onSubmit: Function
    }
})

export default class MemoryModal extends Vue {
    public title!: string;
    public inputState!: MemoryRequest;
    public onClose!: () => void;
    public onSubmit!: () => void;
}
</script>


<template>
    <div class="modal-outer">
        <div class="modal-backdrop" @click="onClose" />
        <div class="modal-inner">
            <h3 class="text-center">{{ title }}</h3>
            <div class="m-y-small">
                <span>タイトル<span class="required">※</span></span>
                <el-input
                    v-model="inputState.title"
                    maxlength="20"
                    show-word-limit
                    type="text"
                />
            </div>
            <div class="m-y-small">
                <span>説明</span>
                <el-input
                    v-model="inputState.description"
                    maxlength="50"
                    show-word-limit
                    type="textarea"
                />
            </div>
            <div class="m-y-small modal-is-public">
                <span>タイムラインに公開</span>
                <el-switch v-model="inputState.isPublic" />
            </div>
            <div class="m-t-middle modal-button-wrapper">
                <el-button
                    @click="onClose"
                >キャンセル</el-button>
                <el-button
                    :disabled="inputState.title === ''"
                    @click="onSubmit"
                    type="primary"
                >保存</el-button>
            </div>
        </div>
    </div>
</template>