import { Box, VStack, Text } from "@chakra-ui/layout"
import PreviewPost from "../../components/PreviewPost"

const ThreadView = ({posts = [{},{},{},{},{},{},{},{},], errorMsg}) => {
    if (!posts || posts.length == 0) {
        return (<Box w="2xl" minH="100vh" padding="1" >
            <Text>Select a channel</Text>
        </Box>)
    }
    if ( errorMsg ) {
        return (<Box w="2xl" minH="100vh" padding="1" >
            <Text>{errorMsg}</Text>
        </Box>)
    }
    return <Box w="2xl" minH="100vh" padding="1" >
        { posts.map((post,i) => 
            <PreviewPost  key={i+"post"} 
                avatar={post.account.avatar}
                title={post.account.display_name !== "" ? 
                    post.account.display_name : post.account.acct}
                username={post.account.username}
                msg={post.content}
                parseMsg={true}
            />)}
    </Box>
}

export default ThreadView