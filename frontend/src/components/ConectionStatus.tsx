import { Box, Circle, Text } from "@/styles/components/ConectionStatus";
import { FC, ReactNode, useMemo } from "react";

interface Props {
    ok: boolean;
    children: ReactNode;
}

const ConectionStatus: FC<Props> = ({ ok, children }) => {
    return (
        <Box>
            <Circle $isOK={ok} />
            <Text $isOK={ok}>
                { children }
            </Text>  
        </Box>
    )
}

export default ConectionStatus