import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom"
import App from './pages/App/App'
import Hello from './pages/Hello/Hello'
import Home from './pages/Home/Home'

const router = createBrowserRouter([
    {
        path: '/',
        element: <App />,
        children: [
            {
                path: '/hello',
                element: <Hello />
            },
            {
                path: '/home',
                element: <Home/>
            }
        ]
    },
])

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <RouterProvider router = {router} />
    </React.StrictMode>
)
