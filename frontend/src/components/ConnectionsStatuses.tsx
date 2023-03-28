import ConectionStatus from "@/components/ConectionStatus"
import useFetch from "@/hooks/useFetch"
import { PingResponse } from "@/models/ping"
import { StatsResponse } from "@/models/stats"
import { Container } from "@/styles/components/ConnectionsStatuses"
import { endpoints } from "@/utils/endpoints"
import { FC } from "react"

const ConectionsStatuses: FC = () => {
    const {data: stats, error: sError, isLoading: sIsloading} = useFetch<StatsResponse>(endpoints.getStats())
    const {data: ping, error: pError, isLoading: pIsLoading} = useFetch<PingResponse>(endpoints.getAPIPing())

    return (
        <Container>
            {ping && <ConectionStatus ok={ping.ok}>API</ConectionStatus>}

            {stats && (
            <>
                <ConectionStatus ok={stats.is_local_conected}>Local</ConectionStatus>
                <ConectionStatus ok={stats.is_remote_conected}>Remoto</ConectionStatus>
            </>
            )}
        </Container>
    )
}

export default ConectionsStatuses