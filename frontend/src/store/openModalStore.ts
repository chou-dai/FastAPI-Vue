import { defineStore } from "pinia";

export interface OpenModalType {
    isOpenedCreateMemoryModal: boolean;
    isOpenedUpdateMemoryModal: boolean;
    isOpenedSignupModal: boolean;
    isOpenedLoginModal: boolean;
}

export const useOpenModalStore = defineStore("openModal", {

    state: (): OpenModalType => ({
        isOpenedCreateMemoryModal: false,
        isOpenedUpdateMemoryModal: false,
        isOpenedSignupModal: false,
        isOpenedLoginModal: false,
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
        },
        openSignupModal() {
            this.isOpenedSignupModal = true;
        },
        closeSignupModal() {
            this.isOpenedSignupModal = false;
        },
        openLoginModal() {
            this.isOpenedLoginModal = true;
        },
        closeLoginModal() {
            this.isOpenedLoginModal = false;
        }
    }
});