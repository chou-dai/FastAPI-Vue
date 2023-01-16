<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { MemoryRequest } from '@/client';

@Options({
    props: {
        title: String,
        inputState: Object,
        changeInputState: Function,
        onClose: Function,
        onSubmit: Function
    },
})

export default class MemoryModal extends Vue {
    public title!: string;
    public inputState!: MemoryRequest;
    public changeInputState!: () => void;
    public onClose!: () => void;
    public onSubmit!: () => void;
}
</script>


<template>
    <div class="modal-outer">
        <div class="modal-inner">
            <h2>{{ title }}</h2>
            タイトル
            <input
                name="title"
                :value="inputState.title"
                @input="changeInputState"
            />
            説明
            <input
                name="description"
                :value="inputState.description"
                @input="changeInputState"
            />
            <div>
                <button @click="onClose">キャンセル</button>
                <button @click="onSubmit">送信</button>
            </div>
        </div>
    </div>
</template>