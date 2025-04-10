import { useParams } from "react-router-dom"
import {useClients} from "../../src/components/helperFunc/Queries"
import { Input, Button, Field, Blockquote } from "@chakra-ui/react"

import {useRef} from "react"

const CustomerDetail = () => {
  const { phone } = useParams()
  const { data: clients, isLoading, error } = useClients()

  if (isLoading) return <div>Loading client...</div>
  if (error) return <div>Error loading client data.</div>

  const client = clients?.find((c) => c.phone.toString() === phone)

  if (!client) return <div>Client not found.</div>

const nameRef = useRef<HTMLInputElement>(null)
const phoneRef = useRef<HTMLInputElement>(null)
const addressRef = useRef<HTMLInputElement>(null)
const messageRef = useRef<HTMLInputElement>(null)
const estimationRef = useRef<HTMLInputElement>(null)
const bookingTimeRef = useRef<HTMLInputElement>(null)
const signedCustomerRef = useRef<HTMLInputElement>(null)
const workDoneRef = useRef<HTMLInputElement>(null)
const workTimeRef = useRef<HTMLInputElement>(null)

const handleSubmit = () => {
    const result = {
    name: nameRef.current?.value || nameRef.current?.placeholder,
    phone: phoneRef.current?.value || phoneRef.current?.placeholder,
    address: addressRef.current?.value || addressRef.current?.placeholder,
    message: messageRef.current?.value || messageRef.current?.placeholder,
    estimation: estimationRef.current?.value || estimationRef.current?.placeholder,
    booking_time: bookingTimeRef.current?.value || bookingTimeRef.current?.placeholder,
    signed_customer: signedCustomerRef.current?.value || signedCustomerRef.current?.placeholder,
    work_done: workDoneRef.current?.value || workDoneRef.current?.placeholder,
    work_time: workTimeRef.current?.value || workTimeRef.current?.placeholder,
    }

    console.log("Saved:", result)

    // Clear all fields
    const refs = [
    nameRef, phoneRef, addressRef, messageRef, estimationRef,
    bookingTimeRef, signedCustomerRef, workDoneRef, workTimeRef
    ]
    refs.forEach(ref => {
    if (ref.current) ref.current.value = ""
    })
}



return (
    <>

    <Blockquote.Root>
        <Blockquote.Content>
        Note that if you wish to change the number of the customer, 
        pls add a new client and delete the current client profile.
        </Blockquote.Content>
    </Blockquote.Root>
    <Field.Root>
        <Field.Label>Name</Field.Label>
        <Input ref={nameRef} placeholder={client.name} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Phone</Field.Label>
        <Input ref={phoneRef} placeholder={client.phone.toString()} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Address</Field.Label>
        <Input ref={addressRef} placeholder={client.address} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Message</Field.Label>
        <Input ref={messageRef} placeholder={client.message} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Estimation</Field.Label>
        <Input ref={estimationRef} placeholder={client.estimation.toString()} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Booking Time</Field.Label>
        <Input ref={bookingTimeRef} placeholder={client.booking_time} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Signed Customer</Field.Label>
        <Input ref={signedCustomerRef} placeholder={client.signed_customer.toString()} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Work Done</Field.Label>
        <Input ref={workDoneRef} placeholder={client.work_done.toString()} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Work Time</Field.Label>
        <Input ref={workTimeRef} placeholder={client.work_time} />
    </Field.Root>

    <Button mt={4} onClick={handleSubmit}>
        Submit New Change
    </Button>
    </>
)}

export default CustomerDetail
