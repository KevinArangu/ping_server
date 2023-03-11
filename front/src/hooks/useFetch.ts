import fetcher from "@/utils/fetcher"
import useSWR  from "swr"

export default function useFetch(url:string) {
    return useSWR(url, fetcher, { refreshInterval: 60000 })
}