import {defineStore} from "pinia";
import {loginReq, smsCodeReq} from "@/api/login.ts";
import {Response} from "@/api/common_type.ts";
import {LoginReq, LoginResp} from "@/api/login_type.ts";

const userStore = defineStore('User', {
    state: () => {
        return {
            visible: false,
            code: '',
            info: (JSON.parse(localStorage.getItem("UserInfo") as string) || {}) as LoginResp,
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
        async login(req: LoginReq) {
            let result: Response<LoginResp> = await loginReq(req);
            if (result.code == 200) {
                this.info = result.data;
                // 持久化存储
                localStorage.setItem("UserInfo", JSON.stringify(this.info));
                return "ok";
            } else {
                return Promise.reject(new Error(result.message));
            }
        },
    },
    getters: {},
});

export default userStore;
