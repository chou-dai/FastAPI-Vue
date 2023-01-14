<template>
    <div class="hello">
        <button @click="fetchMemoryData">送信</button>
        <h1>{{ msg }}</h1>
        <ul>
            <li v-for="memory in memories" :key="memory.id">
                {{ memory.title }}
            </li>
        </ul>
    </div>
</template>

<script lang="ts">
import { Memory } from "@/client";
import { memoryApi } from "@/client/clientWrapper";
import { Options, Vue } from "vue-class-component";

@Options({
    props: {
        msg: String,
    },
})

export default class HelloWorld extends Vue {
    public msg!: string;
    public memories: Array<Memory> = [];
    public async fetchMemoryData() {
        const res = await memoryApi.memoriesGet();
        this.memories = res.data.memories;
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
    margin: 40px 0 0;
}
ul {
    list-style-type: none;
    padding: 0;
}
li {
    display: inline-block;
    margin: 0 10px;
}
a {
    color: #42b983;
}
</style>
