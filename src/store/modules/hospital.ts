import {defineStore} from "pinia";
import {Response} from "@/api/common_type.ts";
import {Department, HospitalDetail} from "@/api/hospital_type.ts";
import {cardDepartReq, cardDetailReq} from "@/api/hospital.ts";
import {HospitalDetailState} from "@/store/modules/interface";

const useDetailStore = defineStore('HospitalDetail', {
    state: (): HospitalDetailState => {
        return {
            hospitalDetail: ({} as HospitalDetail),
            departments: ([] as Department[]),
        }
    },
    actions: {
        async getHospitalDetail(hoscode: string) {
            const result: Response<HospitalDetail> = await cardDetailReq(hoscode);
            if (result.code == 200) {
                console.log('cardDetail:\n', result.data);
                this.hospitalDetail = result.data;
            }
        },
        async getDepartment(hoscode: string) {
            const result: Response<Department[]> = await cardDepartReq(hoscode);
            if (result.code == 200) {
                console.log('departments:\n', result.data);
                this.departments = result.data;
            }
        }
    },
    getters: {},
});
export default useDetailStore;
