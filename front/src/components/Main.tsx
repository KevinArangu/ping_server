import { type FC, type ReactNode } from "react";

interface Props {
  children: ReactNode
}

const Main: FC<Props> = (props) => {
	return (
		<main className="h-screen overflow-y-auto bg-gray-800">
			<div className="container grid px-6 mx-auto">{props.children}</div>
		</main>
	)
};

export default Main