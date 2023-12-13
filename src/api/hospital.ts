import request from "@/http/request";
import {Response} from "@/api/common_type.ts";
import {Department, HospitalDetail} from "@/api/hospital_type.ts";

enum Api {
    hospital_detail = '/hosp/hospital/{hoscode}',
    hospital_depart = '/hosp/hospital/department/{hoscode}',
}

export const cardDetailReq = (hoscode: string) => {
    return request.get<any, Response<HospitalDetail>>(Api.hospital_detail.replace('{hoscode}', hoscode))
}
export const cardDepartReq = (hoscode: string) => {
    return request.get<any, Response<Department[]>>(Api.hospital_depart.replace('{hoscode}', hoscode))
}
