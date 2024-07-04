import { useState } from 'react'
import { useNavigate } from "react-router-dom";

const Login = () => {

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate()

    const getData = async () => {
        const url = "http://localhost:8000/user/login/";

        try {
            const res = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({ email, password }),
            });
            if (!res.ok) {
                throw new Error(`Response status: ${res.status}`);
            }

            const jsonResponse = await res.json();
            console.log(jsonResponse);

            navigate("/");

        } catch (error) {
            console.error(error.message);
        }
    }

    return (
        <>
            <input
                type="email"
                onChange={(e) => setEmail(e.target.value)}
                value={email}
            />
            <input
                type="password"
                onChange={(e) => setPassword(e.target.value)}
                value={password}
            />
            <button onClick={getData}>Login</button>
        </>
    )
}

export default Login