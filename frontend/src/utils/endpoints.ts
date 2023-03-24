import { API_URL } from "@/configs";

export const endpoints = {
    getStats: () => `${API_URL}/stats`,
    getAPIPing: () => `${API_URL}/ping`,
}