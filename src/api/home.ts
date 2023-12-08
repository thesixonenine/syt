import request from "@/http/request";

enum Api {
    card_list = "/hosp/hospital/{pageNo}/{pageSize}"
}

export const cardListReq = (page: number, pageSize: number) => {
    return request.get(Api.card_list.replace('{pageNo}', String(page))
        .replace('{pageSize}', String(pageSize)));
}
