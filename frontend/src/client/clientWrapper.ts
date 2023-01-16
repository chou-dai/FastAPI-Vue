import {
    MemoryApi, UserApi,
} from "./api";
import { Configuration } from "./configuration";

const protocol = location.protocol;
const hostname = location.hostname;
const basePath = protocol + "//" + hostname;

const option = {};

export const memoryApi = new MemoryApi({baseOptions:option, basePath:basePath} as Configuration);
export const userApi = new UserApi({baseOptions:option, basePath:basePath} as Configuration);