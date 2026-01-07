import './globals.css'

export const metadata = {
    title: 'War Room V1',
    description: 'Political Campaign Management System',
}

export default function RootLayout({ children }) {
    return (
        <html lang="th">
            <body>{children}</body>
        </html>
    )
}
