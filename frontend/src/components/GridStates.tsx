import { FC, ReactNode } from "react";

interface Props {
    children: ReactNode;
}

const GridStates: FC<Props> = ({ children }) => {
    return(<section className="grid md:grid-flow-col md:auto-cols-auto gap-4">
        {children}
    </section>)
}

export default GridStates