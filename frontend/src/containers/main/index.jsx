import { Center, Flex, HStack } from '@chakra-ui/layout'
import React from 'react'
import ChannelsView from './ChannelsView'
import ThreadView from './ThreadView'

class Main extends React.Component {
    onAddChannel(channel) {
        console.log(channel)
    }
    render() {
        return (<Center> 
            <Flex>
                <ChannelsView onAdd={this.onAddChannel}/>
                <ThreadView />
            </Flex>
        </Center>)
    }
}

export default Main