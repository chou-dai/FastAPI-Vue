<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { MyMemory } from '@/client';
import { MoreFilled, Timer } from '@element-plus/icons-vue';
import { parseDate } from "../util/timeUtil";
import TimeCount from './TimeCount.vue';
import BaseButton from "./BaseButton.vue";

@Options({
    components: {
        Timer,
        MoreFilled,
        TimeCount,
        BaseButton
    },
    props: {
        memory: Object,
        onClickEdit: Function,
        onClickDelete: Function
    },
})
export default class MyMemoryListItem extends Vue {
    public memory!: MyMemory;
    public onClickEdit!: () => void;
    public onClickDelete!: () => void;

    public parseDateByCreatedAt(date: Date) {
        return parseDate(date);
    }
}
</script>

<template>
    <div class="list-item-content">
        <TimeCount :createdAt="new Date(memory.createdAt.replace('Z',''))" />
        <div class="item-detail">
            <div class="item-date-wrapper">
                <div class="item-date">
                    <el-icon><Timer /></el-icon>
                    <span>{{ parseDateByCreatedAt(new Date(memory.createdAt.replace('Z',''))) }}</span>
                </div>
                <el-tag
                    v-if="memory.isPublic"
                    size="small"
                    effect="dark"
                    round
                >公開中</el-tag>
            </div>
            <h4 class="item-title">{{ memory.title }}</h4>
            <div class="item-description">{{ memory.description }}</div>
        </div>
    </div>
    <el-dropdown trigger="click">
        <el-button circle>
            <el-icon><MoreFilled /></el-icon>
        </el-button>
        <template #dropdown>
            <el-dropdown-menu class="dropdown-menu">
                <BaseButton
                    title="編集"
                    :onClick="onClickEdit"
                />
                <el-button
                    @click="onClickDelete"
                    color="#dc3545"
                >削除</el-button>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>