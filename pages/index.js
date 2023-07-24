import Link from "next/link";
import Header from "../components/header";
import {FaDiscord} from 'react-icons/fa'
import { useRouter } from "next/router";
import Content from "@/components/content";
import Features from "@/components/feature";

export default function Home() {
  <div>
    <head>
      <title>Home</title>
    </head>
  </div>
  return (
    <div>
      <Header/>
      {/* <FaDiscord/> */}
      <Content/>
      <Features/>
    </div>
  )
}


