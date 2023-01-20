<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { PublicMemory } from '@/client';
import { UserFilled, Timer } from '@element-plus/icons-vue';
import { parseDate } from "../util/timeUtil";
import TimeCount from './TimeCount.vue';

@Options({
    components: {
        UserFilled,
        Timer,
        TimeCount
    },
    props: {
        memory: Object
    }
})
export default class PublicMemoryListItem extends Vue {
    public memory!: PublicMemory;

    public parseDateByCreatedAt(date: Date) {
        return parseDate(date);
    }
}
</script>


<template>
    <div class="list-item-content">
        <TimeCount :createdAt="new Date(memory.createdAt.replace('Z',''))" />
        <div>
            <h4>{{ memory.title }}</h4>
            <div>{{ memory.description }}</div>
            <div class="flex">
                <el-icon><UserFilled /></el-icon>
                <span>{{ memory.userName }}</span>
            </div>
            <div class="flex">
                <el-icon><Timer /></el-icon>
                <span>{{ parseDateByCreatedAt(new Date(memory.createdAt.replace('Z',''))) }}</span>
            </div>
        </div>
    </div>
</template>