import { FC, ReactNode, useMemo } from "react";

interface Props {
    ok: boolean;
    children: ReactNode;
}

const ConectionStatus: FC<Props> = ({ ok, children }) => {
    const textColor = useMemo(() => ok ? "text-green-600" : "text-red-500", [ok])
    const bgColor = useMemo(() => ok ? "bg-green-600" : "bg-red-500", [ok])

    return (
        <div className="flex items-center space-x-1 bg-white rounded-full px-2 py-0.5 shadow-md border-2 border-gray-400 cursor-default text-xs md:text-sm">
            <div className={`h-4 w-4 rounded-full ${bgColor}`} />
            <p className={`font-bold ${textColor}`}>
                {children}
            </p>  
        </div>
    )
}

export default ConectionStatus