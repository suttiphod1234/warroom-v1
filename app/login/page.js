'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'

export default function LoginPage() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const router = useRouter()

    const handleLogin = async (e) => {
        e.preventDefault()
        alert('Backend not deployed yet. Demo only.')
    }

    return (
        <div className="flex min-h-screen items-center justify-center bg-gray-900 text-white">
            <div className="w-full max-w-md p-8 space-y-6 bg-gray-800 rounded-xl shadow-2xl border border-gray-700">
                <h2 className="text-3xl font-bold text-center text-blue-500">WAR ROOM ACCESS</h2>
                <form onSubmit={handleLogin} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium">Username</label>
                        <input
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            className="w-full px-4 py-2 mt-1 bg-gray-700 border border-gray-600 rounded focus:ring-2 focus:ring-blue-500 outline-none text-white"
                        />
                    </div>
                    <div>
                        <label className="block text-sm font-medium">Password</label>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="w-full px-4 py-2 mt-1 bg-gray-700 border border-gray-600 rounded focus:ring-2 focus:ring-blue-500 outline-none text-white"
                        />
                    </div>
                    <button type="submit" className="w-full py-3 font-bold bg-blue-600 hover:bg-blue-700 rounded transition duration-200">
                        ENTER WAR ROOM
                    </button>
                </form>
            </div>
        </div>
    )
}
