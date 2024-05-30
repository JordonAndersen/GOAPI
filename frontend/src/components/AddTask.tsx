import React from 'react';
import  { useState } from 'react';
import { useForm } from '@mantine/form';
import { Button, Group, Modal, TextInput, Textarea  } from '@mantine/core';
import { ENDPOINT, Todo} from "../App";
import './button.css';
import './modal.css'
import {KeyedMutator } from "swr";

function AddTask({mutate}: {mutate: KeyedMutator<Todo[]>}){
    const [open, setOpen] = useState(false)
    const form = useForm({
        initialValues:{
            title: "",
            Description: "",
        },

    })

    async function createTask(values: { title: string; Description: string }) {
      // Assuming userID is available from somewhere in your frontend, replace '1' with the actual user ID
      const userID = 1; // Replace '1' with the actual user ID
      const taskData = { ...values, userID }; // Include user ID in task data
  
      const updated = await fetch(`${ENDPOINT}/tasks`, {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify(taskData),
      }).then((r) => r.json());
  
      mutate(updated);
      form.reset();
      setOpen(false);
  }
  

   return ( 
    <>
    <h1 className='center'>Get Organized, Get Things Done: Welcome to the GO Task Manager!</h1>
    <Modal className='box' opened={open} onClose={() => setOpen(false)} title="Create Task">
    <form onSubmit={form.onSubmit(createTask)}>
    
    <div className="modal">
      <div className="modal__title"><TextInput
            required
            mb={12}
            label="Task Name"      
            placeholder="What task would you like to create?"

            {...form.getInputProps("title")}
          /></div>
      <div className="modal__title"> <Textarea
      
            required
            mb={12}
            label="Description"
            placeholder="Tell me more..."
            {...form.getInputProps("body")}
          /></div>
       <Button className="modal__btn" type="submit">
            Create Task</Button>
    </div>

        

      
        
    
       

        
        </form>
  </Modal>

  <Group style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
    <Button className="modal__btn" fullWidth mb={12} onClick={() => setOpen(true)}>
        Add Task
    </Button>

  </Group>
  </>

    )
}

export default AddTask;

