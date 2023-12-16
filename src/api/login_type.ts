export interface LoginReq {
    phone: string,
    code: string
}

export interface LoginResp {
    name: string,
    token: string
}
