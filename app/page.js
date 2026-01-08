'use client'

import { useRouter } from 'next/navigation'

export default function HomePage() {
    const router = useRouter()

    return (
        <div style={{
            minHeight: '100vh',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            textAlign: 'center',
            padding: '20px'
        }}>
            <div style={{ maxWidth: '800px' }}>
                <h1 style={{
                    fontSize: '4rem',
                    fontWeight: 'bold',
                    marginBottom: '1.5rem',
                    background: 'linear-gradient(to right, #60a5fa, #a78bfa, #f472b6)',
                    WebkitBackgroundClip: 'text',
                    WebkitTextFillColor: 'transparent'
                }}>
                    WAR ROOM V1
                </h1>

                <p style={{
                    fontSize: '1.5rem',
                    color: '#d1d5db',
                    marginBottom: '2rem'
                }}>
                    ระบบบริหารจัดการเครือข่ายการเมือง พร้อม AI
                </p>

                <button
                    onClick={() => router.push('/login')}
                    style={{
                        padding: '1rem 2rem',
                        background: '#2563eb',
                        color: 'white',
                        border: 'none',
                        borderRadius: '0.5rem',
                        fontSize: '1.125rem',
                        fontWeight: 'bold',
                        cursor: 'pointer',
                        transition: 'all 0.2s'
                    }}
                    onMouseOver={(e) => e.target.style.background = '#1d4ed8'}
                    onMouseOut={(e) => e.target.style.background = '#2563eb'}
                >
                    เริ่มใช้งาน →
                </button>
            </div>
        </div>
    )
}
