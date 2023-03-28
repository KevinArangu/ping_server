import { BaseBar, Box, ColorBar, Count, Percent, Statistics, Title } from "@/styles/components/ProgressCard";
import { FC, useCallback, useMemo } from "react"

    
type PercentMode = "floor" | "ceil"; 
interface Props {
    title: string;
    percentMode: PercentMode;
    isGreen: boolean;
    total?: number;
    count?: number;
}

const ProgressCard: FC<Props> = ({ title, total, count, isGreen, percentMode }) => {
    const getMath = useCallback(
        (percentMode: PercentMode, total: number, count: number) => {
          if (percentMode === "floor") {
              return Math.floor((count * 100) / total)
          } else if (percentMode === "ceil") {
              return Math.ceil((count * 100) / total)
          }},[])
    const percent = useMemo(() => total && count ? getMath(percentMode, total, count) : 0, [count, total, getMath, percentMode])
    
    return (  
        <Box>
            <Title>{title}</Title>
            <Statistics>
                <Percent
                    $isGreen={isGreen} 
                    $existValues={!!total && !! count}
                >
                    {percent}%
                </Percent>
                {(!!total && !! count) && (
                    <Count>{count} de {total}</Count>
                )}
            </Statistics>

            <BaseBar>
                <ColorBar 
                    $isGreen={isGreen}
                    style={{ width: percent + "%" }}
                />
            </BaseBar>
        </Box>
    )
}

export default ProgressCard