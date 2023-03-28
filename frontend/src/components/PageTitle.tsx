import { type FC, type ReactNode } from "react"
import { Container, Text } from "@/styles/components/PageTitle"

interface Props {
    children: ReactNode
  }

const PageTitle: FC<Props> = ({ children }) => {
    return (
      <Container>
        <Text>{children}</Text>
      </Container>
    )
  }
  
  export default PageTitle