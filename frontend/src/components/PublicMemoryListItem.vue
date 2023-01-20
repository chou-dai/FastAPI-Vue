<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { PublicMemory } from '@/client';
import { Delete, Edit, Position, MoreFilled } from '@element-plus/icons-vue';
import TimeCount from '../components/TimeCount.vue';

@Options({
    components: {
        Delete,
        Edit,
        Position,
        MoreFilled,
        TimeCount
    },
    props: {
        memory: Object
    },
})
export default class PublicMemoryListItem extends Vue {
    public memory!: PublicMemory;
}
</script>


<template>
    <div class="list-item-content">
        <TimeCount :createdAt="new Date(memory.createdAt.replace('Z',''))" />
        <div>
            <div>{{ memory.title }}</div>
            <div>{{ memory.description }}</div>
            <div>{{ memory.userName }}</div>
        </div>
    </div>
</template>