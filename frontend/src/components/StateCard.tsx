import { BaseBar, Box, ColorBar, Count, Percent, Statistics, Title } from "@/styles/components/ProgressCard";
import { FC, useCallback, useMemo } from "react"

    
type PercentMode = "floor" | "ceil"; 
interface Props {
    title: string;
    isOK: boolean;
}

const StateCard: FC<Props> = ({ title, isOK }) => {    
    return (  
        <Box>
            <Title>{title}</Title>
            <div className="flex space-x-2 items-center">
                <div className="h-5 w-5 bg-green-500 rounded-full"></div>
                <p className="text-green-500">{ isOK ? "Conectado" : "Desconectado" }</p>
            </div>
        </Box>
    )
}

export default StateCard