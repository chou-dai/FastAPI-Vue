<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
    props: {
        title: String,
        disable: Boolean,
        onClick: Function
    }
})
export default class BaseButton extends Vue {
    public title!: string;
    public size: string = "default";
    public onClick!: () => void;
}
</script>


<template>
    <el-button
        :size="size"
        @click="onClick"
        type="primary"
    >
        {{ title }}
    </el-button>
</template>