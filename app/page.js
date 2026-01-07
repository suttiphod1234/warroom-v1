'use client'

import { useRouter } from 'next/navigation'

export default function HomePage() {
    const router = useRouter()

    return (
        <div className="min-h-screen bg-black text-white flex items-center justify-center">
            <div className="text-center max-w-4xl px-4">
                <h1 className="text-6xl md:text-7xl font-bold mb-6 bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">
                    WAR ROOM V1
                </h1>

                <p className="text-xl md:text-2xl text-gray-300 mb-8">
                    ระบบบริหารจัดการเครือข่ายการเมือง พร้อม AI วิเคราะห์ความเห็น
                </p>

                <div className="flex gap-4 justify-center">
                    <button
                        onClick={() => router.push('/login')}
                        className="px-8 py-4 bg-blue-600 hover:bg-blue-700 rounded-lg font-bold text-lg transition transform hover:scale-105"
                    >
                        เริ่มใช้งาน →
                    </button>
                </div>
            </div>
        </div>
    )
}
