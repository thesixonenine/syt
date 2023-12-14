import {defineStore} from "pinia";
import {smsCodeReq} from "@/api/login.ts";
import {Response} from "@/api/common_type.ts";

const userStore = defineStore('User', {
    state: () => {
        return {
            visible: false,
            code: '',
        }
    },
    actions: {
        async getCode(mobile: string) {
            let result: Response<string> = await smsCodeReq(mobile)
            if (result.code == 200) {
                this.code = result.data;
                console.log('sms code:', this.code);
                if (this.code == null) {
                    this.code = '123456';
                }
                return this.code;
            } else {
                return Promise.reject(new Error(result.message));
            }
        },
    },
    getters: {},
});

export default userStore;
