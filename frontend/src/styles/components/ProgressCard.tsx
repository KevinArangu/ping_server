import tw from "tailwind-styled-components";

interface PercentProps {
    $existValues: boolean;
    $isGreen: boolean;
}
interface ColorBarProps {
    $isGreen: boolean;
}

export const Box = tw.div`flex flex-col p-4 bg-white shadow-lg rounded-2xl min-w-[9rem] dark:bg-gray-800`
export const Title = tw.p`ml-2 text-gray-700 text-md font-medium dark:text-gray-50 truncate`
export const Statistics = tw.div`flex justify-between items-center`
export const Percent = tw.p<PercentProps>`
    my-4 text-4xl font-bold text-left
    ${ ({$existValues}) => $existValues ? "" : "!text-gray-700 bg-gray-700 animate-pulse" }
    ${ ({$isGreen}) => $isGreen ? "text-green-500" : "text-red-500" }
`
export const Count = tw.p`text-gray-600 font-medium hidden md:block`
export const BaseBar = tw.div`relative h-2 bg-gray-200 rounded w-full`
export const ColorBar = tw.div<ColorBarProps>`
    absolute top-0 left-0 h-2 rounded
    ${ ({ $isGreen }) => $isGreen ? "bg-green-500" : "bg-red-500" }
`
