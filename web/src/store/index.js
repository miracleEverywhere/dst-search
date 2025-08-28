import {defineStore} from "pinia";

export const useStore = defineStore('store', {
    state: () => {
        return {
            data: []
        }
    },
    actions: {
        clean() {
            this.data = []
        }
    }
})