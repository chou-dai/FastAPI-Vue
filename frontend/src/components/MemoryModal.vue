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
            <el-row class="m-y-small" justify="space-between">
                <span>タイムラインに公開</span>
                <el-switch v-model="inputState.isPublic" />
            </el-row>
            <el-row class="m-t-middle" justify="space-between">
                <BaseButton
                    title="キャンセル"
                    :onClick="onClose"
                />
                <el-button
                    :disabled="inputState.title === ''"
                    @click="onSubmit"
                >保存</el-button>
            </el-row>
        </div>
    </div>
</template>