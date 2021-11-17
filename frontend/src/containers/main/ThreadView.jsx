import { Box, VStack } from "@chakra-ui/layout"
import PreviewPost from "../../components/PreviewPost"

const ThreadView = ({posts = [{},{},{},{},{},{},{},{},]}) => {
    return <Box w="2xl" padding="1" bgColor="yellow.500">
        { posts.map((post,i) => <PreviewPost  key={i+"post"} />)}
    </Box>
}

export default ThreadView