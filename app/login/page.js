'use client'

import { useState } from 'react'

export default function LoginPage() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')

    const handleLogin = (e) => {
        e.preventDefault()
        alert('Backend not deployed yet. This is a demo.')
    }

    return (
        <div style={{
            minHeight: '100vh',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            background: '#111827'
        }}>
            <div style={{
                width: '100%',
                maxWidth: '28rem',
                padding: '2rem',
                background: '#1f2937',
                borderRadius: '0.75rem',
                boxShadow: '0 20px 25px -5px rgba(0, 0, 0, 0.1)',
                border: '1px solid #374151'
            }}>
                <h2 style={{
                    fontSize: '1.875rem',
                    fontWeight: 'bold',
                    textAlign: 'center',
                    color: '#3b82f6',
                    marginBottom: '1.5rem'
                }}>
                    WAR ROOM ACCESS
                </h2>

                <form onSubmit={handleLogin} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
                    <div>
                        <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', marginBottom: '0.25rem' }}>
                            Username
                        </label>
                        <input
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            style={{
                                width: '100%',
                                padding: '0.5rem 1rem',
                                background: '#374151',
                                border: '1px solid #4b5563',
                                borderRadius: '0.375rem',
                                color: 'white',
                                outline: 'none'
                            }}
                        />
                    </div>

                    <div>
                        <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', marginBottom: '0.25rem' }}>
                            Password
                        </label>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            style={{
                                width: '100%',
                                padding: '0.5rem 1rem',
                                background: '#374151',
                                border: '1px solid #4b5563',
                                borderRadius: '0.375rem',
                                color: 'white',
                                outline: 'none'
                            }}
                        />
                    </div>

                    <button
                        type="submit"
                        style={{
                            width: '100%',
                            padding: '0.75rem',
                            background: '#2563eb',
                            color: 'white',
                            border: 'none',
                            borderRadius: '0.375rem',
                            fontWeight: 'bold',
                            cursor: 'pointer',
                            transition: 'background 0.2s'
                        }}
                        onMouseOver={(e) => e.target.style.background = '#1d4ed8'}
                        onMouseOut={(e) => e.target.style.background = '#2563eb'}
                    >
                        ENTER WAR ROOM
                    </button>
                </form>
            </div>
        </div>
    )
}
