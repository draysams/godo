import { Button, Flex, Input, Spinner } from '@chakra-ui/react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { IoMdAdd } from 'react-icons/io';
import { BASE_URL } from '../App';

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState('');
  const [isPending, setIsPending] = useState(false);

  const queryClient = useQueryClient();
  const { mutate: createTodo } = useMutation({
    mutationKey: ['createTodo'],
    mutationFn: async (e: any) => {
      e.preventDefault();
      try {
        setIsPending(true);
        const res = await fetch(`${BASE_URL}/todos`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ body: newTodo }),
        });
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.message || 'Something went wrong!');
        }
        return data;
      } catch (error) {
        console.log(error);
      } finally {
        setIsPending(false);
        setNewTodo('');
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
    },
  });
  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type='text'
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type='submit'
          _active={{
            transform: 'scale(.97)',
          }}
        >
          {isPending ? <Spinner size={'xs'} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
};
export default TodoForm;
