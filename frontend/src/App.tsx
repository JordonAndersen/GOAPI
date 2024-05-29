import './App.css'
import { Box, List, ThemeIcon, MantineProvider } from '@mantine/core';
import useSWR from "swr";
import AddTask from './components/AddTask';
import { CheckCircleFillIcon } from "@primer/octicons-react";

export interface Todo{
  id: number;
  Title: string;
  Description: string;
  Status: boolean;

}

export const ENDPOINT = 'http//:localhost:3000'

const fetcher = (url:string) => fetch(`${ENDPOINT}/${url}`).then((r) => r.json());


function App() {
  const { data, mutate } = useSWR<Todo[]>('tasks', fetcher)

  
async function markTodoAdDone(id: number) {
  const updated = await fetch(`${ENDPOINT}/tasks/${id}`, {
    method: "PATCH",
  }).then((r) => r.json());

  mutate(updated);
}


  return (
    
    <>
    <MantineProvider>
      <Box sx={(theme) => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}>
       <List spacing="xs" size="sm" mb={12} center>
        {data?.map((todo) => {
          return (
            <List.Item
              onClick={() => markTodoAdDone(todo.id)}
              key={`todo_list__${todo.id}`}
              icon={
                todo.Status ? (
                  <ThemeIcon color="teal" size={24} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
                ) : (
                  <ThemeIcon color="gray" size={24} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
                )
              }
            >
              {todo.Title}
            </List.Item>
          );
        })}
      </List>
        <AddTask mutate={mutate}/>
      </Box>
    </MantineProvider>
    

    </>
  )
}

export default App
