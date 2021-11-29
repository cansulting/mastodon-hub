import { Center, Flex, HStack } from '@chakra-ui/layout'
import React from 'react'
import ChannelsView from './ChannelsView'
import ThreadView from './ThreadView'
import { EboxEvent as ActionCenter} from 'elabox-foundation'
import { retrieveTimeline, addChannel, retrieveInitData, removeChannel } from '../../actions/actions'

class Main extends React.Component {
    state = { 
        actionCenter: new ActionCenter('http://' + window.location.hostname),
        posts:[],
        channels: [/*"elastos.social","noc.social","mastodon.online"*/],
        err: null
    }
    onError(err) {
        this.setState({err:err})
    }
    async componentDidMount() {
        await this.state.actionCenter.waitUntilConnected()
        // load user data
        const res = await retrieveInitData(this.state.actionCenter)
        const channels = res.message.channels
        this.setState({channels: [...this.state.channels, ...channels]})
        await this.retrieveChannelTimeline(this.state.channels[0])
    }
    async retrieveChannelTimeline(channel) {
        this.onError(null)
        const res = await retrieveTimeline(this.state.actionCenter, channel)
        if (res.code === 200) {
            this.setState({posts:res.message})
        } else {
            this.onError(res.message)
        }
    }
    async onAddChannel(channel) {
        console.log(channel)
        await addChannel(this.state.actionCenter, channel)
        this.setState({channels: [...this.state.channels, channel]})
    }
    async onRemoveChannel(channelI) {
        const channel = this.state.channels[channelI]
        await removeChannel(this.state.actionCenter, channel)
        var channels = [...this.state.channels]
        channels.splice(channelI, 1)
        this.setState({channels: channels})
    }
    onSelectChannel(index) {
        this.retrieveChannelTimeline(this.state.channels[index])
    }
    render() {
        return (
        <Center> 
            <Flex>
                <ChannelsView 
                    onAdd={this.onAddChannel.bind(this)}
                    onRemoved={this.onRemoveChannel.bind(this)}
                    onSelect={this.onSelectChannel.bind(this)}
                    channels={this.state.channels}
                />
                <ThreadView posts={this.state.posts} errorMsg={this.state.err}/>
            </Flex>
        </Center>)
    }
}

export default Main