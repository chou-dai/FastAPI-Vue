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
    <div>
        <el-button
            :size="size"
            @click="onClick"
        >
            {{ title }}
        </el-button>
    </div>
</template>