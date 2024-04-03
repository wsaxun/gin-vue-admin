export const isNull = (...params) => {
        return params.some(param => param === null || param === "" || param === undefined);
}