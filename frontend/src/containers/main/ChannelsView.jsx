import { Button, IconButton } from "@chakra-ui/button";
import { FiPlus } from "react-icons/fi"
import { Text, Flex, Spacer } from "@chakra-ui/layout";
import { AlertDialog, AlertDialogBody, AlertDialogContent, AlertDialogFooter, AlertDialogHeader, AlertDialogOverlay } from "@chakra-ui/modal";
import { useState } from "react";
import { Input, InputGroup, InputLeftAddon, InputRightAddon } from "@chakra-ui/input";

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

const ChannelsView = ({onAdd=(channel) => {}}) => {
    return (
        <Flex w="xs" bgColor="gray.500">
            <ChannelToolbar onAdd={onAdd}/>
        </Flex>
    )
}

export default ChannelsView;