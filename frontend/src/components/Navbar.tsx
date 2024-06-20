import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <div style={{ width: '200px', height: '100vh', backgroundColor: '#f0f0f0', padding: '10px' }}>
      <ul style={{ listStyleType: 'none', padding: 0 }}>
        <li>
          <Link to="/home">Home</Link>
        </li>
        <li>
          <Link to="/hello">Hello</Link>
        </li>
      </ul>
    </div>
  );
}

export default Navbar;