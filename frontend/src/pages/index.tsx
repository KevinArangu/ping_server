import GridPings from "@/components/GridPings"
import GridStates from "@/components/GridStates"
import PageTitle from "@/components/PageTitle"
import ProgressCard from "@/components/ProgressCard"
import StateCard from "@/components/StateCard"
import useFetch from "@/hooks/useFetch"
import { PingResponse } from "@/models/ping"
import { type StatsResponse } from "@/models/stats"
import { endpoints } from "@/utils/endpoints"
import { FC } from "react"

const Home: FC = () => {
  const {data: stats, error: sError, isLoading: sIsloading} = useFetch<StatsResponse>(endpoints.getStats())
  const {data: ping, error: pError, isLoading: pIsLoading} = useFetch<PingResponse>(endpoints.getAPIPing())

  return (
    <>
      <PageTitle>Estadísticas de mi red</PageTitle>
      <div className="flex flex-col mt-12 space-y-4 md:space-y-6">
        <GridStates>
          <StateCard title="Estado del API" isOK={ping ? ping.ok : undefined} />
          <StateCard title="Estado del gateway" isOK={stats ? stats.is_local_conected : undefined} />
          <StateCard title="Estado remoto" isOK={stats ? stats.is_remote_conected : undefined} />
        </GridStates>
        <GridPings>
          <ProgressCard 
            title="Pings locales completados" 
            isGreen={true}
            percentMode="floor"
            total={stats ? Number.parseInt(stats.total_local_pings) : undefined} 
            count={stats ? Number.parseInt(stats.completed_local_pings) : undefined} 
          />
          <ProgressCard 
            title="Pings locales perdidos" 
            isGreen={false}
            percentMode="ceil"
            total={stats ? Number.parseInt(stats.total_local_pings) : undefined} 
            count={stats ? Number.parseInt(stats.error_local_pings) : undefined} 
          />
          <ProgressCard 
            title="Pings remotos completados" 
            isGreen={true}
            percentMode="floor"
            total={stats ? Number.parseInt(stats.total_remote_pings) : undefined} 
            count={stats ? Number.parseInt(stats.completed_remote_pings) : undefined} 
          />
          <ProgressCard 
            title="Pings remotos perdidos" 
            isGreen={false}
            percentMode="ceil"
            total={stats ? Number.parseInt(stats.total_remote_pings) : undefined} 
            count={stats ? Number.parseInt(stats.error_remote_pings) : undefined} 
          />
        </GridPings>
      </div>
    </>
  )
}

export default Home
