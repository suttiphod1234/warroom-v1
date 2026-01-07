"use client";

import { useRouter } from 'next/navigation';

export default function HomePage() {
    const router = useRouter();

    return (
        <div className="min-h-screen bg-black text-white flex items-center justify-center">
            <div className="text-center">
                <h1 className="text-6xl font-bold mb-6 bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
                    WAR ROOM V1
                </h1>
                <p className="text-xl text-gray-300 mb-8">
                    ระบบบริหารจัดการเครือข่ายการเมือง พร้อม AI
                </p>
                <button
                    onClick={() => router.push('/login')}
                    className="px-8 py-4 bg-blue-600 hover:bg-blue-700 rounded-lg font-bold text-lg"
                >
                    เข้าสู่ระบบ →
                </button>
            </div>
        </div>
    );
}
