import React from "react";
import { Link } from "react-router-dom";
import {FaHome, FaInfo, FaEnvelope} from "react-icons/fa";

const Navbar = () => {
    return (
        <div id="navbar">
            <div id="navbar_list">
                <ul style={{ listStyleType: 'none', padding: 0 }}>
                    <li className="navbar_list_item">
                        <Link to="/home" className="navbar_list_item_link"><FaHome/>Home</Link>
                    </li>
                    <li className="navbar_list_item">

                        <Link to="/hello" className="navbar_list_item_link">Hello</Link>
                    </li>
                </ul>
            </div>
        </div>
    );
}

export default Navbar;