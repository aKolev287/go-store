import { useEffect, useState } from "react"

const HomePage = () => {

    const [user, setUser] = useState({})

    const fetchUser = async () => {
        const url = "http://localhost:8000/user/"
        try {

            const res = await fetch(url, {
                method: "GET",
                credentials: "include"
            })

            if (!res.ok) {
                throw new Error(`Response status: ${res.status}`);
            }

            const data = await res.json()
            setUser(data)


        } catch (error) {
            console.error(error.message)
        }
    }

    useEffect(() => {
        fetchUser()
    }, [])

    return (
        <div>
            <p>Hello {user.Username}</p>
            <p>{user.Email}</p>
        </div>
    )
}

export default HomePage