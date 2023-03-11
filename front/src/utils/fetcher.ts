export default async function fetcher(url: string, init: RequestInit) {
    return fetch(url, init).then(res => res.json())
}