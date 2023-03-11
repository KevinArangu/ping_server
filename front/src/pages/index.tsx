import PageTitle from "@/components/PageTittle"
import useFetch from "@/hooks/useFetch"
import { endpoints } from "@/utils/endpoints"
import { FC } from "react"

const Home: FC = () => {
  const {data, error, isLoading} = useFetch(endpoints.getPokemon("ditto"))

  return (
    <>
      <PageTitle>Estadísticas de la red</PageTitle>
      {isLoading && <p>loading</p>}
      {data && <p>{data.name}</p>}
    </>
  )
}

export default Home
