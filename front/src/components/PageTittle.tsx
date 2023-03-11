import { type FC, type ReactNode } from "react"

interface Props {
    children: ReactNode
  }

const PageTitle: FC<Props> = (props) => {
    return (
      <h1 className="truncate my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">{props.children}</h1>
    )
  }
  
  export default PageTitle