import { Box, Heading, Text, Image, Flex } from '@chakra-ui/react';
import Avatar from "../media/avatar.jpg"
import React from 'react';

const postView = ({
    avatar="https://noc.social/system/accounts/avatars/000/000/795/original/b660d8c590596f59.jpg?1591672357",
    avatarAlt="post",
    title = "hello",
    msg = "this is post description",
    parseMsg = false,
    username = "username"
}) => {
    return (
        <Flex padding="3">
            <Image 
                padding="3"
                borderRadius="full"
                boxSize="70px" 
                src={avatar} 
                alt={avatarAlt}
                fallbackSrc={Avatar}
            />
            <Box flex="1" align="left">
                <Heading fontSize="lg">{title}</Heading>
                <Text fontSize="md" color="gray"> @{username}</Text>
                <Text mt={2}>
                    {!parseMsg && msg}
                    {parseMsg && <div dangerouslySetInnerHTML={{__html:msg}}/>}
                </Text>
            </Box>
        </Flex>
    )
}

export default postView