<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { getYears, getDays, getHoursStr, getMinutesStr, getSecondsStr } from "../util/timeUtil";

@Options({
    props: {
        createdAt: Date
    }
})
export default class TimeCount extends Vue {
    public createdAt!: Date;
    public diff = 0;
    
    private setTimeDiff() {
        const now = Date.now();
        this.diff = now - this.createdAt.getTime();
    }
    created() {
        this.setTimeDiff();
        setInterval(() => {
            this.setTimeDiff();
        }, 1000);
    }

    public createDisplayDays(time: number) {
        const years = getYears(time);
        const days = getDays(time);
        if (years > 0) return `${years}年${days}日`;
        if (days > 0) return `${days}日`;
        return false
    }
    public createDisplayTime(time: number) {
        const hours = getHoursStr(time);
        const minutes = getMinutesStr(time);
        const seconds = getSecondsStr(time);
        return `${hours}:${minutes}:${seconds}`
    }
}
</script>

<template>
    <h3 class="count-timer">
        <div v-if="createDisplayDays(diff)">
            {{ createDisplayDays(diff) }}
        </div>
        <div>
            {{ createDisplayTime(diff) }}
        </div>
    </h3>
</template>