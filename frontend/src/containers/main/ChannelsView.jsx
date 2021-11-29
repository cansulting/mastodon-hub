import { Button, IconButton, ButtonGroup} from "@chakra-ui/button";
import { Menu, MenuList, MenuItem, MenuButton } from "@chakra-ui/react"
import { FiPlus } from "react-icons/fi"
import { AiOutlineMore } from "react-icons/ai"
import { Text, Flex, Spacer, VStack, HStack } from "@chakra-ui/layout";
import { AlertDialog, AlertDialogBody, AlertDialogContent, AlertDialogFooter, AlertDialogHeader, AlertDialogOverlay } from "@chakra-ui/modal";
import { useState } from "react";
import { Input, InputGroup, InputLeftAddon } from "@chakra-ui/input";

const ChannelToolbar = ({onAdd=(channel) => {}}) => {
    const [isOpen, setOpen] = useState(false)
    const onClose = (val) => {
        onAdd(val)
        setOpen(false)
    }
    return (<Flex>
        <Text> Channels</Text>
        <Spacer/>
        <IconButton 
            size="sm" 
            icon={<FiPlus size="25px"/>} 
            onClick={() => setOpen(true)}/>
        <AddChannelDialog isOpen={isOpen} onClose={onClose}/>
    </Flex>)
}

const AddChannelDialog = ({
    isOpen = false, onClose = (val)=>{}}) => {
    const [value, setValue] = useState("")
    const onAdd = () => onClose(value)
    return (
        <AlertDialog
            isOpen={isOpen}
            onClose={onClose}
            closeOnEsc>
            <AlertDialogOverlay>
                <AlertDialogContent>
                    <AlertDialogHeader fontSize="lg" fontWeight="bold" hide>
                        Add Channel
                    </AlertDialogHeader>
                    <AlertDialogBody>
                        <InputGroup>
                            <InputLeftAddon children="https://"/>
                            <Input placeholder="place mastodon host here"
                                onChange={(event) => setValue(event.target.value)}
                            />
                        </InputGroup>
                    </AlertDialogBody>
                    <AlertDialogFooter>
                        <Button onClick={onAdd}>Add Channel</Button>
                    </AlertDialogFooter>
                </AlertDialogContent>
            </AlertDialogOverlay>
        </AlertDialog>
    )
}

// the id 
// the label displayed for this channel
// callback when this channel was selected
// callback when this channel will be removed
const ChannelItem = ({
    id=1,               
    label="channel",    
    onClick=(id)=>{},   
    onRemoved=(id)=>{},  
    isActive=false,
}) => {
    const [showDropdown, setShowDropdown] = useState(false)
    const [focus, setFocus] = useState(false)
    const onRemoveClicked = () => onRemoved(id)
    return (
        <HStack onMouseLeave={() => setShowDropdown(false)} onMouseEnter={() => setShowDropdown(true)}>
            <ButtonGroup size="sm" isAttached variant={!isActive && !showDropdown ? "ghost" : "outline"}>
                <Button fontSize="xl" 
                    mr="-px"
                    onClick={() => onClick(id)}>
                    {label}
                </Button>
                { (showDropdown || focus) && 
                    <Menu onOpen={() => setFocus(true)} onClose={() => setFocus(false)}>
                        <MenuButton as={IconButton} icon={<AiOutlineMore />}/>
                        <MenuList>
                            <MenuItem onClick={onRemoveClicked}>
                                Remove
                            </MenuItem>
                        </MenuList>
                    </Menu>        
                }
            </ButtonGroup>
        </HStack>
    )
}

// display the channel panel
// @active current selected channel
// @channels list of channels
// @onSelect callback when channel was selected
// @onAdd callback when channel was added
// @onRemove callback when channel was removed
const ChannelsView = ({
    active=0,                                                       
    channels=["channel1"],                                          
    onSelect=(index)=>{},                                            
    onAdd=(channel) => {},
    onRemoved=(index)=>{} }) => {
    const [currentActive, setActive] = useState(active)
    const _onSelect=(index) => {
        setActive(index)
        onSelect(index)
    }
    return (
        <VStack w="xs" >
            <ChannelToolbar onAdd={onAdd}/>
            { channels.map( (channel, index) => <ChannelItem 
                key={'channel' + index}
                id={index}
                label={channel}
                onClick={_onSelect}
                isActive={currentActive===index}
                onRemoved={onRemoved}
                />)
            }
        </VStack>
    )
}

export default ChannelsView;