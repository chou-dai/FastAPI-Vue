<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { UserRequest } from '@/client';

@Options({
    props: {
        title: String,
        inputState: Object,
        changeInputState: Function,
        onClose: Function,
        onSubmit: Function
    },
})

export default class UserModal extends Vue {
    public title!: string;
    public inputState!: UserRequest;
    public changeInputState!: () => void;
    public onClose!: () => void;
    public onSubmit!: () => void;
}
</script>


<template>
    <div class="modal-outer">
        <div class="modal-inner">
            <h1>{{ title }}</h1>
            名前
            <input
                name="name"
                :value="inputState.name"
                @input="changeInputState"
            />
            パスワード
            <input
                name="password"
                :value="inputState.password"
                @input="changeInputState"
            />
            <div>
                <button @click="onClose">キャンセル</button>
                <button @click="onSubmit">送信</button>
            </div>
        </div>
    </div>
</template>