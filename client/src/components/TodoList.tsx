import { Flex, Spinner, Stack, Text } from '@chakra-ui/react';
import { useQuery } from '@tanstack/react-query';
import { IoIosGlobe } from 'react-icons/io';
import TodoItem from './TodoItem';

export type Todo = {
  _id: string;
  body: string;
  completed: boolean;
};
const TodoList = () => {
  const { data: todos, isLoading } = useQuery<Todo[]>({
    queryKey: ['todos'],
    queryFn: async () => {
      try {
        const res = await fetch('http://localhost:4000/api/todos');
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.message || 'Something went wrong!');
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
  });
  return (
    <>
      <Text
        fontSize={'4xl'}
        textTransform={'uppercase'}
        fontWeight={'bold'}
        textAlign={'center'}
        my={2}
        color={'blue.400'}
      >
        Today's Tasks
      </Text>
      {isLoading && (
        <Flex justifyContent={'center'} my={4}>
          <Spinner size={'xl'} />
        </Flex>
      )}
      {!isLoading && todos?.length === 0 && (
        <Stack alignItems={'center'} gap='3'>
          <Text fontSize={'xl'} textAlign={'center'} color={'white'}>
            All tasks completed! 🤞
          </Text>
          <IoIosGlobe width={70} height={70} />
        </Stack>
      )}
      <Stack gap={3}>
        {todos?.map((todo) => (
          <TodoItem key={todo._id} todo={todo} />
        ))}
      </Stack>
    </>
  );
};
export default TodoList;
