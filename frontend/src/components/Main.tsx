import { type FC, type ReactNode } from "react";

interface Props {
  children: ReactNode
}

const Main: FC<Props> = ({ children }) => {
	return (
		<main className="h-screen overflow-y-auto bg-gray-700 select-none">
			<div className="container grid px-6 mx-auto">{children}</div>
		</main>
	)
};

export default Main