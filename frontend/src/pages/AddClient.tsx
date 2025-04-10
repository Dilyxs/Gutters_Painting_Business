
import { Input, Button, Field } from "@chakra-ui/react"

import {useRef} from "react"
const CustomerDetail = () => {

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

const returnAccept = () => {
    
}

return (
    <>
    <Field.Root>
        <Field.Label>Name String</Field.Label>
        <Input ref={nameRef}/>
    </Field.Root>

    <Field.Root>
        <Field.Label>Phone int64</Field.Label>
        <Input ref={phoneRef} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Address string</Field.Label>
        <Input ref={addressRef}  />
    </Field.Root>

    <Field.Root>
        <Field.Label>Message string</Field.Label>
        <Input ref={messageRef} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Estimation boolean</Field.Label>
        <Input ref={estimationRef} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Booking Time string</Field.Label>
        <Input ref={bookingTimeRef}/>
    </Field.Root>

    <Field.Root>
        <Field.Label>Signed Customer boolean</Field.Label>
        <Input ref={signedCustomerRef} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Work Done boolean</Field.Label>
        <Input ref={workDoneRef} />
    </Field.Root>

    <Field.Root>
        <Field.Label>Work Time string</Field.Label>
        <Input ref={workTimeRef}/>
    </Field.Root>

    <Button mt={4} onClick={handleSubmit}>
        Submit New Change
    </Button>
    </>
)}

export default CustomerDetail
