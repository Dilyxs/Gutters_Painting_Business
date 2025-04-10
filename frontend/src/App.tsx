import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import { Container } from "@chakra-ui/react"


import LoginVerification from "./pages/LoginVerification"
import DashBoard from "./pages/Dashboard"
import CustomerDetail from "./pages/CustomerDetail" // make sure it's created
import AddNewClient from "./pages/AddClient"


function App(){
  return (
    <Router>
      <Container maxW="container.lg" py={5}>
        <Routes>
        <Route path="/login" element={<LoginVerification />} />
        <Route path="/dashboard" element={<DashBoard />} />
        <Route path="/customer/:phone" element={<CustomerDetail />} />
        <Route path="/addclient" element={<AddNewClient />} />

        </Routes>
      </Container>
    </Router>
  )
}

export default App