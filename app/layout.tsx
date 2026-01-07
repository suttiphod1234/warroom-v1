import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
    title: "War Room V1 - Political Campaign Management",
    description: "ระบบบริหารจัดการเครือข่ายการเมือง พร้อม AI และ MLM",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="th">
            <body>{children}</body>
        </html>
    );
}
