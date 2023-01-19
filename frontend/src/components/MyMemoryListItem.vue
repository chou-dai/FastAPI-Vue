<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { MyMemory } from '@/client';
import { Delete, Edit, Position, MoreFilled } from '@element-plus/icons-vue';

@Options({
    components: {
        Delete,
        Edit,
        Position,
        MoreFilled
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
}
</script>


<template>
    <div class="list-item-content">
        <div>{{ memory.createdAt }}</div>
        <div>
            <div>{{ memory.title }}</div>
            <el-tag v-if="memory.isPublic" size="small" effect="dark" round>
                公開中
            </el-tag>
            <div>{{ memory.description }}</div>
        </div>
    </div>
    <div class="list-item-button">
        <el-dropdown trigger="click">
            <el-button circle>
                <el-icon><MoreFilled /></el-icon>
            </el-button>
            <template #dropdown>
                <el-dropdown-menu>
                    <el-button-group>
                        <el-button type="primary" @click="onClickEdit">
                            編集
                        </el-button>
                        <el-button type="danger" @click="onClickDelete">
                            削除
                        </el-button>
                    </el-button-group>
                </el-dropdown-menu>
            </template>
          </el-dropdown>
    </div>
</template>