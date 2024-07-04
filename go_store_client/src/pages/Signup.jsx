import { useState } from 'react'
import { useNavigate } from "react-router-dom";

const Signup = () => {
  const navigate = useNavigate()
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("")

  const signUp = async () => {
    const url = "http://localhost:8000/user/signup/";

    try {
        const res = await fetch(url, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({email, password, username})
        })

        const jsonResponse = await res.json();
        console.log(jsonResponse);

        navigate("/login");0

        if (!res.ok) {
            throw new Error(`Something went wrong. Code: ${res.status}`)
        }


    } catch (error) {
        console.error(error.message)
    }
  }



  return (
    <>
    <div>
        <div className="flex flex-col">
            <label htmlFor="username">Username*</label>
            <input type="text" name="username" onChange={(e) => setUsername(e.target.value)} />

            <label htmlFor="email">Email*</label>
            <input type="email" name="email" onChange={(e) => setEmail(e.target.value) } />

            <label htmlFor="email">Password*</label>
            <input type="password" name="password" onChange={(e) => setPassword(e.target.value)} />


            <button onClick={signUp}>Sign Up</button>
        </div>
    </div>
    </>

  )
}

export default Signup