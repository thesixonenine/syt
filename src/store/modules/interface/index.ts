import {Department, HospitalDetail} from "@/api/hospital_type.ts";

export interface HospitalDetailState {
    hospitalDetail: HospitalDetail,
    departments: Department[]
}