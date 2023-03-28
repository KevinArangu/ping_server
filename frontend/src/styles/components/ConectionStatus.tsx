import tw from "tailwind-styled-components"

interface CircleProps {
    $isOK: boolean;
}
interface TextProps {
    $isOK: boolean;
}

export const Box = tw.div`
    flex items-center px-2 py-0.5 
    space-x-1 bg-white rounded-full
    shadow-md border-2 border-gray-400 cursor-default 
    text-xs md:text-sm
`
export const Circle = tw.div<CircleProps>`
    h-4 w-4 rounded-full
    ${ ({ $isOK }) => $isOK ? "bg-green-600" : "bg-red-500" }
`
export const Text = tw.div<TextProps>`
    font-bold
    ${ ({ $isOK }) => $isOK ? "text-green-600" : "text-red-500" }
`