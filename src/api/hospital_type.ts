export interface BookingRule {
    cycle: number;
    releaseTime: string;
    stopTime: string;
    quitDay: number;
    quitTime: string;
    rule: string[];
}

export interface Param {
    hostypeString: string;
    fullAddress: string;
}

export interface Hospital {
    id: string;
    createTime: string;
    updateTime: string;
    isDeleted: number;
    param: Param;
    hoscode: string;
    hosname: string;
    hostype: string;
    provinceCode: string;
    cityCode: string;
    districtCode: string;
    address: string;
    logoData: string;
    intro: string;
    route: string;
    status: number;
    bookingRule?: any;
}

export interface HospitalDetail {
    bookingRule: BookingRule;
    hospital: Hospital;
}

export interface Department {
    "depcode": string,
    "depname": string,
    "children"?: Department[],
}

export interface BookingSchedule {
    workDate: string,
    workDateMd: string,
    dayOfWeek: string,
    docCount: number,
    reservedNumber: number,
    availableNumber: number,
    status: number,
}

export interface BookingBase {
    workDateString: string,
    releaseTime: string,
    bigname: string,
    stopTime: string,
    depname: string,
    hosname: string,
}

export interface Booking {
    total: number,
    bookingScheduleList: BookingSchedule[],
    baseMap: BookingBase,
}

export interface DoctorParam {
    dayOfWeek: string,
    depname: string,
    hosname: string,
}

export interface Doctor {
    id: string,
    createTime: string,
    updateTime: string,
    isDelete: number,
    param: DoctorParam,
    hoscode: string,
    depcode: string,
    title: string,
    docname: string,
    skill: string,
    workDate: string,
    workTime: number,
    reservedNumber: number,
    availableNumber: number,
    amount: number,
    status: number,
    hosScheduleId: string,
}
