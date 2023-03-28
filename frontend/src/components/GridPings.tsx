import { FC, ReactNode } from "react";

interface Props {
    children: ReactNode;
}

const GridLayout: FC<Props> = ({ children }) => {
    return(<section className="grid md:grid-cols-2 xl:grid-cols-4 gap-4">
        {children}
    </section>)
}

export default GridLayout