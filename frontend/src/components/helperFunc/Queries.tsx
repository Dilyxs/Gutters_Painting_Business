// hooks/useClients.ts
import { useQuery } from "@tanstack/react-query"
import { useMutation } from "@tanstack/react-query"

export type Client = {
  name: string
  phone: number
  address: string
  message: string
  estimation: boolean
  booking_time: string
  signed_customer: boolean
  work_done: boolean
  work_time: string
}

export function useClients() {
  return useQuery<Client[]>({
    queryKey: ["clients"],
    queryFn: async () => {
      const res = await fetch("http://localhost:4321/")
      const data = await res.json()
      if (!res.ok) {
        throw new Error(data.error || "Something went wrong fetching the data.")
      }
      return data || []
    },
  })
}

export function updateClient(payload:object){
    return useMutation<Client[]>(
        {
            mutationKey:['updateClient'],
            mutationFn:async () => {
                const res = await fetch("http://localhost:4321/UpdateClient/", {
                    method: "PUT",
                    headers: {
                      "Content-Type": "application/json",
                    },
                    // optionally pass body:
                    body: JSON.stringify(payload),
                  })
                  
                const data = await res.json()
                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong fetching the data.")
                }
                return data || []

            }
        }
    )
}
export function LoginVerifier(payload:object){
    return useMutation<Client[]>(
        {
            mutationKey:['LoginVerifier'],
            mutationFn:async () => {
                const res = await fetch("http://localhost:4321/LoginVerification/", {
                    method: "POST",
                    headers: {
                      "Content-Type": "application/json",
                    },
                    // optionally pass body:
                    body: JSON.stringify(payload),
                  })
                  
                const data = await res.json()
                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong fetching the data.")
                }
                return data || []

            }
        }
    )}