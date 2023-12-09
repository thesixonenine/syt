import request from "@/http/request";
import type {Hospital} from "@/api/home_type.ts";
import {Response} from "@/api/common_type.ts";
import {HospitalDict} from "@/api/home_type.ts";

enum Api {
    card_list = "/hosp/hospital/{pageNo}/{pageSize}",
    card_level = '/cmn/dict/findByDictCode/{dictCode}',
}

export const cardListReq = (page: number, pageSize: number) => {
    return request.get<any, Response<Hospital>>(Api.card_list.replace('{pageNo}', String(page))
        .replace('{pageSize}', String(pageSize)));
};

export const cardLevelReq = (dictCode: string) => {
    return request.get<any, Response<HospitalDict[]>>(Api.card_level.replace('{dictCode}', dictCode));
};
