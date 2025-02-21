import { Box, Container, Flex, Text } from '@chakra-ui/react';
import { IoAirplane, IoMoon } from 'react-icons/io5';
import { LuSun } from 'react-icons/lu';

export default function Navbar() {
  return (
    <Container maxW={'900px'}>
      <Box bg={'gray.400'} px={4} my={4} borderRadius={'5'}>
        <Flex h={16} alignItems={'center'} justifyContent={'space-between'}>
          {/* LEFT SIDE */}
          <Flex
            justifyContent={'center'}
            alignItems={'center'}
            gap={3}
            display={{ base: 'none', sm: 'flex' }}
          >
            <IoMoon width={50} height={50} />
            <Text fontSize={'40'}>+</Text>
            <LuSun width={40} height={40} />
            <Text fontSize={'40'}>=</Text>
            <IoAirplane width={50} height={50} />
          </Flex>

          {/* RIGHT SIDE */}
          <Flex alignItems={'center'} gap={3}>
            <Text fontSize={'lg'} fontWeight={500}>
              Daily Tasks
            </Text>
          </Flex>
        </Flex>
      </Box>
    </Container>
  );
}
