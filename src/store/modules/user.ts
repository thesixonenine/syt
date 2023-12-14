import {defineStore} from "pinia";

const userStore = defineStore('User', {
    state: () => {
        return {
            visible: false,
        }
    },
    actions: {},
    getters: {},
});

export default userStore;
