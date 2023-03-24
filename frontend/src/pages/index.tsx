import ConectionStatuses from "@/components/ConnectionsStatuses"
import PageTitle from "@/components/PageTittle"
import useFetch from "@/hooks/useFetch"
import { type StatsResponse } from "@/models/stats"
import { endpoints } from "@/utils/endpoints"
import { FC } from "react"

const Home: FC = () => {
  const {data: stats, error: sError, isLoading: sIsloading} = useFetch<StatsResponse>(endpoints.getStats())

  return (
    <>
      <PageTitle>Estadísticas de mi red</PageTitle>

      {sIsloading && <p>loading</p>}
      {stats && <div className="font-medium text-lg">
        <p>Total pings: {stats.total_pings}</p> 
        <div className="w-full my-2"></div>
        <p className="text-white">Total local pings: {stats.total_local_pings}</p> 
        <p className="text-green-600">Completed local pings: {stats.completed_local_pings}</p>
        <p className="text-red-600">Error local pings: {stats.error_local_pings}</p>
        <div className="w-full my-2"></div>
        <p className="text-white">Total remote pings: {stats.total_remote_pings}</p> 
        <p className="text-green-600">Completed remote pings: {stats.completed_remote_pings}</p>
        <p className="text-red-600">Error remote pings: {stats.error_remote_pings}</p>
      </div>}
    </>
  )
}

export default Home
