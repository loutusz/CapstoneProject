import Link from "next/link";
import Layout from "../components/layout"

export default function Home() {
  <div>
    <head>
      <title>Home</title>
    </head>
  </div>
  return (
    <div>
      <Layout/>
    <h2>
      <Link href="/loginPage">Go to Login Page</Link>
    </h2>
    <h2>
      <Link href="/registerPage">Go to Register Page</Link>
    </h2>
    </div>
  )
}