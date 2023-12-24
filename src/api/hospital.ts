import request from "@/http/request";
import {Response} from "@/api/common_type.ts";
import {Booking, Department, Doctor, HospitalDetail} from "@/api/hospital_type.ts";

enum Api {
    hospital_detail = '/hosp/hospital/{hoscode}',
    hospital_depart = '/hosp/hospital/department/{hoscode}',
    hospital_depart_booking = '/hosp/hospital/auth/getBookingScheduleRule/{page}/{limit}/{hoscode}/{depcode}',
    doctor_detail = '/hosp/hospital/auth/findScheduleList/{hoscode}/{depcode}/{workDate}',
}

export const cardDetailReq = (hoscode: string) => {
    return request.get<any, Response<HospitalDetail>>(Api.hospital_detail.replace('{hoscode}', hoscode))
}
export const cardDepartReq = (hoscode: string) => {
    return request.get<any, Response<Department[]>>(Api.hospital_depart.replace('{hoscode}', hoscode))
}
export const cardDepartBookingReq = (page: number, limit: number, hoscode: string, depcode: string) => {
    return request.get<any, Response<Booking>>(Api.hospital_depart_booking
        .replace('{page}', String(page))
        .replace('{limit}', String(limit))
        .replace('{hoscode}', hoscode)
        .replace('{depcode}', depcode))
}
export const doctorDetailReq = (hoscode: string, depcode: string, workDate: string) => {
    return request.get<any, Response<Doctor[]>>(Api.doctor_detail
        .replace('{hoscode}', hoscode)
        .replace('{depcode}', depcode)
        .replace('{workDate}', workDate))
}
