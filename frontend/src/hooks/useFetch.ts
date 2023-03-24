import fetcher from "@/utils/fetcher"
import useSWR  from "swr"

export default function useFetch<T>(url:string) {
    return useSWR<T>(url, fetcher, { refreshInterval: 10000 })
}