import request from "@/http/request.ts";
import {Response} from "@/api/common_type.ts";
import {LoginReq, LoginResp} from "@/api/login_type.ts";

enum Api {
    SMS_CODE = '/sms/send/{mobile}',
    LOGIN = '/user/login',
}

export const smsCodeReq = (mobile: string) => {
    return request.get<any, Response<string>>(Api.SMS_CODE.replace('{mobile}', mobile));
}

export const loginReq = (data: LoginReq) => {
    return request.post<any, Response<LoginResp>>(Api.LOGIN, data);
}
