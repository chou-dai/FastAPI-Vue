<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { MyMemory } from '@/client';
import { Delete, Edit } from '@element-plus/icons-vue';

@Options({
    components: {
        Delete,
        Edit
    },
    props: {
        memory: Object,
        onClickEdit: Function,
        onClickDelete: Function
    },
})
export default class MemoryListItem extends Vue {
    public memory!: MyMemory;
    public onClickEdit!: () => void;
    public onClickDelete!: () => void;
}
</script>


<template>
    <div>
        <div>●:{{ memory.title }}</div>
        <div>○:{{ memory.description }}</div>
        <div>○:{{ memory.createdAt }}</div>
        <div>○:{{ memory.isPublic }}</div>
        <el-button-group>
            <el-button type="primary" @click="onClickEdit">
                <el-icon>
                    <Edit />
                </el-icon>
            </el-button>
            <el-button type="danger" @click="onClickDelete">
                <el-icon>
                    <Delete />
                </el-icon>
            </el-button>
        </el-button-group>
    </div>
</template>