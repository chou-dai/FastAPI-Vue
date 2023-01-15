import { defineStore } from "pinia";

export interface OpenModalType {
    isOpenedCreateMemoryModal: boolean;
    isOpenedUpdateMemoryModal: boolean;
}

export const useOpenModalStore = defineStore("openModal", {

    state: (): OpenModalType => ({
        isOpenedCreateMemoryModal: false,
        isOpenedUpdateMemoryModal: false,
    }),

    getters: {},

    actions: {
        openCreateMemoryModal() {
            this.isOpenedCreateMemoryModal = true;
        },
        closeCreateMemoryModal() {
            this.isOpenedCreateMemoryModal = false;
        },
        openUpdateMemoryModal() {
            this.isOpenedUpdateMemoryModal = true;
        },
        closeUpdateMemoryModal() {
            this.isOpenedUpdateMemoryModal = false;
        }
    }
});