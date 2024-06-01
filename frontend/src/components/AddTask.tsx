import { useState } from 'react'; // Importing the useState hook from React.
import { useForm } from '@mantine/form'; // Importing the useForm hook from the Mantine library for form handling.
import { Button, Group, Modal, TextInput, Textarea } from '@mantine/core'; // Importing components from the Mantine library for UI elements.
import { ENDPOINT, Todo } from "../App"; // Importing constants and types from another module (Assuming ENDPOINT is the API endpoint URL and Todo is a type).
import '../styles/button.css'; // Importing custom CSS for button styling.
import '../styles/modal.css'; // Importing custom CSS for modal styling.
import { KeyedMutator } from "swr"; // Importing KeyedMutator type from SWR for mutation.

function AddTask({ mutate }: { mutate: KeyedMutator<Todo[]> }) { // Defining the AddTask component, which takes a mutate function as a prop. This function is used to revalidate the data in SWR.
    const [open, setOpen] = useState(false); // Using the useState hook to manage the open state of the modal (true if open, false if closed).
    const form = useForm({ // Using the useForm hook to manage form state and validation.
        initialValues: { // Initial form values.
            title: "", // Default value for the title field.
            description: "", // Default value for the description field.
        },
    });

    async function createTask(values: { title: string; description: string }) { // Defining the async function createTask to handle form submission.
        const updated = await fetch(`${ENDPOINT}/simple-tasks`, { // Making a POST request to the simple-tasks endpoint to create a new task.
            method: "POST", // Specifying the HTTP method as POST.
            headers: {
                "Content-Type": "application/json", // Setting the Content-Type header to indicate that the request body is in JSON format.
            },
            body: JSON.stringify(values), // Converting the form values to a JSON string to send in the request body.
        }).then((r) => r.json()); // Parsing the JSON response from the server.

        mutate(updated); // Calling the mutate function to update the local data cache with the new task.
        form.reset(); // Resetting the form fields to their initial values.
        setOpen(false); // Closing the modal.
    }

    return (
        <>
            <h1 className='center'>Get Organized, Get Things Done: Welcome to the GO Task Manager!</h1> {/* Title centered using CSS class 'center' */}
            <Modal className='box' opened={open} onClose={() => setOpen(false)} title="Create Task"> {/* Modal component from Mantine library */}
                <form onSubmit={form.onSubmit(createTask)}> {/* Form element with onSubmit handler that calls createTask */}
                    <div className="modal"> {/* Div container for modal content */}
                        <div className="modal__title"> {/* Div for modal title */}
                            <TextInput
                                required // Marking the input field as required.
                                mb={12} // Adding margin bottom of 12 units.
                                label="Task Name" // Label for the input field.
                                placeholder="What task would you like to create?" // Placeholder text for the input field.
                                {...form.getInputProps("title")} // Binding the input field to the form's title value using spread syntax.
                            />
                        </div>
                        <div className="modal__title"> {/* Div for modal description */}
                            <Textarea
                                required // Marking the textarea field as required.
                                mb={12} // Adding margin bottom of 12 units.
                                label="Description" // Label for the textarea field.
                                placeholder="Tell me more..." // Placeholder text for the textarea field.
                                {...form.getInputProps("description")} // Binding the textarea field to the form's description value using spread syntax.
                            />
                        </div>
                        <Button className="modal__btn" type="submit">Create Task</Button> {/* Submit button */}
                    </div>
                </form>
            </Modal>

            <Group style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}> {/* Group container for Add Task button, centered using inline styles */}
                <Button className="modal__btn" fullWidth mb={12} onClick={() => setOpen(true)}>Add Task</Button> {/* Button to open the modal */}
            </Group>
        </>
    );
}

export default AddTask; // Exporting the AddTask component as the default export of the module.
