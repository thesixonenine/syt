import request from "@/http/request.ts";
import {Response} from "@/api/common_type.ts";

enum Api {
    SMS_CODE = '/sms/send/{mobile}',
}

export const smsCodeReq = (mobile: string) => {
    return request.get<any, Response<string>>(Api.SMS_CODE.replace('{mobile}', mobile))
}