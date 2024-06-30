import { useEffect, useState } from "react";

function App() {
  const [jwt, setJwt] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function getCookie() {
    const cookieUrl = "http://localhost:8000/getcookie/";

    try {
      const res = await fetch(cookieUrl, {
        method: "GET",
        credentials: "include", // Important
      });
      if (!res.ok) {
        throw new Error(`Response status: ${res.status}`);
      }

      const data = await res.json();
      setJwt(data.token);
    } catch (err) {
      console.error(err.message);
    }
  }

  async function getData() {
    const url = "http://localhost:8000/user/login/";

    try {
      const res = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include", // Important
        body: JSON.stringify({ email, password }),
      });
      if (!res.ok) {
        throw new Error(`Response status: ${res.status}`);
      }

      const jsonResponse = await res.json();
      console.log(jsonResponse);
    } catch (error) {
      console.error(error.message);
    }
  }

  useEffect(() => {
    getCookie()
  }, [])

  return (
    <>
      <p>Hello</p>

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

      {jwt ? <p>Token: {jwt}</p> : <p>No token</p>}
    </>
  );
}

export default App;
