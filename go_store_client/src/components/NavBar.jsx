import { Link } from "react-router-dom";

const NavBar = () => {
  return (
    <div>
        <ul>
            <li>
                <Link to={"/login"}>Login</Link>
            </li>
            <li>
                <Link to={"/"}>Home</Link>
            </li>
            <li>
              <Link to={"/signup"}>Sign Up</Link>
            </li>
        </ul>
    </div>
  )
}

export default NavBar