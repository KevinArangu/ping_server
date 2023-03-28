import { Container, Main } from "@/styles/components/Main";
import { type FC, type ReactNode } from "react";

interface Props {
  children: ReactNode
}

const MainComponent: FC<Props> = ({ children }) => {
	return (
		<Main>
			<Container>{children}</Container>
		</Main>
	)
};

export default MainComponent