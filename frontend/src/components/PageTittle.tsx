import { type FC, type ReactNode } from "react"
import ConectionsStatuses from "./ConnectionsStatuses"

interface Props {
    children: ReactNode
  }

const PageTitle: FC<Props> = ({ children }) => {
    return (
      <div className="my-6 flex flex-col space-y-2 justify-between sm:flex-row sm:space-y-0">
        <h1 className="truncate font-semibold text-white text-xl md:text-2xl xl:text-4xl">{children}</h1>
        <ConectionsStatuses />
      </div>
    )
  }
  
  export default PageTitle