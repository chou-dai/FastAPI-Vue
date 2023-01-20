import {format} from "date-fns";

export const parseDate = (date: Date) => {
    return format(date, "yyyy/MM/dd");
};


const ms = 1000;
const s = 60;
const m = 60;
const h = 24;
const d = 365;

export const getSecondsStr = (time: number) => {
    const seconds =  Math.floor((time / ms) % s);
    return seconds.toString().padStart( 2, '0');
};

export const getMinutesStr = (time: number) => {
    const minutes = Math.floor((time / (s * ms)) % m);
    return minutes.toString().padStart( 2, '0');
};

export const getHoursStr = (time: number) => {
    const hours = Math.floor((time / (m * s * ms)) % h);
    return hours.toString().padStart( 2, '0');
};

export const getDays = (time: number) => {
    return Math.floor((time / (h * m * s * ms)) % d);
};

export const getYears = (time: number) => {
    return Math.floor(time / (d * h * m * s * ms));
};